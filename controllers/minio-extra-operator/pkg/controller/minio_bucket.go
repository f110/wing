package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/minio/minio-go/v6"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"

	miniov1alpha1 "github.com/f110/tools/controllers/minio-extra-operator/pkg/api/minio/v1alpha1"
	clientset "github.com/f110/tools/controllers/minio-extra-operator/pkg/client/versioned"
	informers "github.com/f110/tools/controllers/minio-extra-operator/pkg/informers/externalversions"
	mbLister "github.com/f110/tools/controllers/minio-extra-operator/pkg/listers/minio/v1alpha1"
)

type MinIOBucketController struct {
	config            *rest.Config
	coreClient        *kubernetes.Clientset
	mClient           *clientset.Clientset
	mbLister          mbLister.MinIOBucketLister
	mbListerHasSynced cache.InformerSynced

	queue workqueue.RateLimitingInterface

	runOutsideCluster bool
}

func NewMinioBucketController(ctx context.Context, client *kubernetes.Clientset, cfg *rest.Config, runOutsideCluster bool) (*MinIOBucketController, error) {
	mClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	_, apiList, err := client.ServerGroupsAndResources()
	if err != nil {
		return nil, err
	}
	found := false
	for _, v := range apiList {
		if v.GroupVersion == "miniocontroller.min.io/v1beta1" {
			for _, v := range v.APIResources {
				if v.Kind == "MinIOInstance" {
					found = true
					break
				}
			}
		}
	}
	if !found {
		return nil, errors.New("minio-operator is not installed")
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(mClient, 30*time.Second)
	mbInformer := sharedInformerFactory.Minio().V1alpha1().MinIOBuckets()

	c := &MinIOBucketController{
		config:            cfg,
		coreClient:        client,
		mClient:           mClient,
		mbLister:          mbInformer.Lister(),
		mbListerHasSynced: mbInformer.Informer().HasSynced,
		queue:             workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "MinioBucket"),
		runOutsideCluster: runOutsideCluster,
	}

	mbInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addMinioBucket,
		UpdateFunc: c.updateMinioBucket,
		DeleteFunc: c.deleteMinioBucket,
	})

	sharedInformerFactory.Start(ctx.Done())

	return c, nil
}

func (c *MinIOBucketController) syncMinioBucket(key string) error {
	klog.V(4).Info("syncMinioBucket")
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	minioBucket, err := c.mbLister.MinIOBuckets(namespace).Get(name)
	if err != nil && apierrors.IsNotFound(err) {
		klog.V(4).Infof("%s/%s is not found", namespace, name)
		return nil
	} else if err != nil {
		return err
	}
	currentBucket := minioBucket.DeepCopy()

	instances, err := c.mClient.MinV1beta1().MinIOInstances(namespace).List(metav1.ListOptions{LabelSelector: metav1.FormatLabelSelector(&minioBucket.Spec.Selector)})
	if err != nil {
		return err
	}
	if len(instances.Items) == 0 {
		klog.V(4).Infof("%s not found", metav1.FormatLabelSelector(&minioBucket.Spec.Selector))
		return nil
	}
	if len(instances.Items) > 1 {
		return errors.New("found some instances")
	}
	instance := instances.Items[0]

	creds, err := c.coreClient.CoreV1().Secrets(instance.Namespace).Get(instance.Spec.CredsSecret.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	svc, err := c.coreClient.CoreV1().Services(instance.Namespace).Get(fmt.Sprintf("%s-hl-svc", instance.Name), metav1.GetOptions{})
	if err != nil {
		return err
	}

	instanceEndpoint := fmt.Sprintf("%s-hl-svc.%s.svc:%d", instance.Name, instance.Namespace, svc.Spec.Ports[0].Port)
	if c.runOutsideCluster {
		forwarder, err := c.portForward(svc, int(svc.Spec.Ports[0].Port))
		if err != nil {
			return err
		}
		defer forwarder.Close()

		ports, err := forwarder.GetPorts()
		if err != nil {
			return err
		}
		instanceEndpoint = fmt.Sprintf("0:%d", ports[0].Local)
	}

	mc, err := minio.New(
		instanceEndpoint,
		string(creds.Data["accesskey"]),
		string(creds.Data["secretkey"]),
		false,
	)
	if err != nil {
		return err
	}
	if exists, err := mc.BucketExists(minioBucket.Name); err != nil {
		return err
	} else if exists {
		klog.V(4).Infof("%s already exists", minioBucket.Name)
		return nil
	}

	if err := mc.MakeBucket(minioBucket.Name, ""); err != nil {
		return err
	}

	minioBucket.Status.Ready = true

	if !reflect.DeepEqual(minioBucket, currentBucket) {
		_, err = c.mClient.MinioV1alpha1().MinIOBuckets(minioBucket.Namespace).Update(minioBucket)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *MinIOBucketController) portForward(svc *corev1.Service, port int) (*portforward.PortForwarder, error) {
	selector := labels.SelectorFromSet(svc.Spec.Selector)
	podList, err := c.coreClient.CoreV1().Pods(svc.Namespace).List(metav1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		return nil, err
	}
	var pod *corev1.Pod
	for i, v := range podList.Items {
		if v.Status.Phase == corev1.PodRunning {
			pod = &podList.Items[i]
			break
		}
	}
	if pod == nil {
		return nil, errors.New("all pods are not running yet")
	}

	req := c.coreClient.CoreV1().RESTClient().Post().Resource("pods").Namespace(svc.Namespace).Name(pod.Name).SubResource("portforward")
	transport, upgrader, err := spdy.RoundTripperFor(c.config)
	if err != nil {
		return nil, err
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, http.MethodPost, req.URL())

	readyCh := make(chan struct{})
	pf, err := portforward.New(dialer, []string{fmt.Sprintf(":%d", port)}, context.Background().Done(), readyCh, nil, nil)
	if err != nil {
		return nil, err
	}
	go func() {
		err := pf.ForwardPorts()
		if err != nil {
			switch v := err.(type) {
			case *apierrors.StatusError:
				klog.Info(v)
			}
			klog.Error(err)
		}
	}()

	select {
	case <-readyCh:
	case <-time.After(5 * time.Second):
		return nil, errors.New("timed out")
	}

	return pf, nil
}

func (c *MinIOBucketController) Run(ctx context.Context, workers int) {
	defer c.queue.ShutDown()

	klog.Info("Wait for informer caches to sync")
	if !cache.WaitForCacheSync(ctx.Done(), c.mbListerHasSynced) {
		klog.Error("Failed to sync informer caches")
		return
	}

	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, ctx.Done())
	}

	klog.V(2).Info("Start workers of MinIOBucketController")
	<-ctx.Done()
	klog.V(2).Info("Shutdown workers")
}

func (c *MinIOBucketController) worker() {
	defer klog.V(4).Info("Finish worker")

	for c.processNextItem() {
	}
}

func (c *MinIOBucketController) processNextItem() bool {
	defer klog.V(4).Info("Finish processNextItem")

	obj, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	klog.V(4).Infof("Get next queue: %s", obj)

	err := func(obj interface{}) error {
		defer c.queue.Done(obj)

		err := c.syncMinioBucket(obj.(string))
		if err != nil {
			c.queue.AddRateLimited(obj)
			return err
		}

		c.queue.Forget(obj)
		return nil
	}(obj)
	if err != nil {
		klog.Info(err)
		return true
	}

	return true
}

func (c *MinIOBucketController) enqueue(bucket *miniov1alpha1.MinIOBucket) {
	if key, err := cache.MetaNamespaceKeyFunc(bucket); err != nil {
		return
	} else {
		c.queue.Add(key)
	}
}

func (c *MinIOBucketController) addMinioBucket(obj interface{}) {
	b := obj.(*miniov1alpha1.MinIOBucket)

	c.enqueue(b)
}

func (c *MinIOBucketController) updateMinioBucket(old, cur interface{}) {
	oldBucket := old.(*miniov1alpha1.MinIOBucket)
	curBucket := cur.(*miniov1alpha1.MinIOBucket)

	if oldBucket.UID != curBucket.UID {
		if key, err := cache.MetaNamespaceKeyFunc(oldBucket); err != nil {
			klog.Info(err)
			return
		} else {
			c.deleteMinioBucket(cache.DeletedFinalStateUnknown{Key: key, Obj: oldBucket})
		}
	}

	c.enqueue(curBucket)
}

func (c *MinIOBucketController) deleteMinioBucket(obj interface{}) {
	b, ok := obj.(*miniov1alpha1.MinIOBucket)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			return
		}
		b, ok = tombstone.Obj.(*miniov1alpha1.MinIOBucket)
		if !ok {
			return
		}
	}

	c.enqueue(b)
}
