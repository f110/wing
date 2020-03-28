package main

import (
	"context"
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog"

	"github.com/f110/tools/controllers/harbor-project-operator/pkg/controller"
	"github.com/f110/tools/lib/signals"
)

func main() {
	id := ""
	enableLeaderElection := false
	leaseLockName := ""
	leaseLockNamespace := ""
	harborNamespace := ""
	harborServiceName := ""
	adminSecretName := ""
	dev := false
	fs := flag.NewFlagSet("harbor-project-operator", flag.ExitOnError)
	fs.StringVar(&id, "id", uuid.New().String(), "the holder identity name")
	fs.BoolVar(&enableLeaderElection, "enable-leader-election", enableLeaderElection,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	fs.StringVar(&leaseLockName, "lease-lock-name", "", "the lease lock resource name")
	fs.StringVar(&leaseLockNamespace, "lease-lock-namespace", "", "the lease lock resource namespace")
	fs.StringVar(&harborNamespace, "harbor-namespace", "", "the namespace name to which harbor service belongs")
	fs.StringVar(&harborServiceName, "harbor-service-name", "", "the service name of harbor")
	fs.StringVar(&adminSecretName, "admin-secret-name", "", "the secret name that including admin password")
	fs.BoolVar(&dev, "dev", dev, "development mode")
	klog.InitFlags(fs)
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	signals.SetupSignalHandler(cancelFunc)

	kubeconfigPath := ""
	if dev {
		h, err := os.UserHomeDir()
		if err != nil {
			klog.Error(err)
			os.Exit(1)
		}
		kubeconfigPath = filepath.Join(h, ".kube", "config")
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		os.Exit(1)
	}

	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      leaseLockName,
			Namespace: leaseLockNamespace,
		},
		Client: kubeClient.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: id,
		},
	}
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   30 * time.Second,
		RenewDeadline:   15 * time.Second,
		RetryPeriod:     5 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				c, err := controller.NewHarborProjectController(ctx, kubeClient, cfg, harborNamespace, harborServiceName, adminSecretName, dev)
				if err != nil {
					klog.Error(err)
					os.Exit(1)
				}

				c.Run(ctx, 1)
				klog.Info("Shutdown")
			},
			OnStoppedLeading: func() {
				klog.Infof("leader lost: %s", id)
			},
			OnNewLeader: func(identity string) {
				if identity == id {
					return
				}
				klog.Infof("new leader elected: %s", identity)
			},
		},
	})
}
