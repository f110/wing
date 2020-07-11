// Generated by protoc-ddl.
// protoc-gen-entity: v0.1
package database

import (
	"bytes"
	"sync"
	"time"

	"go.f110.dev/protoc-ddl"
)

var _ = time.Time{}
var _ = bytes.Buffer{}

type Column struct {
	Name  string
	Value interface{}
}

type SourceRepository struct {
	Id        int32
	Url       string
	CloneUrl  string
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time

	mu   sync.Mutex
	mark *SourceRepository
}

func (e *SourceRepository) ResetMark() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.mark = e.Copy()
}

func (e *SourceRepository) IsChanged() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.Url != e.mark.Url ||
		e.CloneUrl != e.mark.CloneUrl ||
		e.Name != e.mark.Name ||
		!e.CreatedAt.Equal(e.mark.CreatedAt) ||
		((e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil))
}

func (e *SourceRepository) ChangedColumn() []ddl.Column {
	e.mu.Lock()
	defer e.mu.Unlock()

	res := make([]ddl.Column, 0)
	if e.Url != e.mark.Url {
		res = append(res, ddl.Column{Name: "url", Value: e.Url})
	}
	if e.CloneUrl != e.mark.CloneUrl {
		res = append(res, ddl.Column{Name: "clone_url", Value: e.CloneUrl})
	}
	if e.Name != e.mark.Name {
		res = append(res, ddl.Column{Name: "name", Value: e.Name})
	}
	if !e.CreatedAt.Equal(e.mark.CreatedAt) {
		res = append(res, ddl.Column{Name: "created_at", Value: e.CreatedAt})
	}
	if (e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil) {
		if e.UpdatedAt != nil {
			res = append(res, ddl.Column{Name: "updated_at", Value: *e.UpdatedAt})
		} else {
			res = append(res, ddl.Column{Name: "updated_at", Value: nil})
		}
	}

	return res
}

func (e *SourceRepository) Copy() *SourceRepository {
	n := &SourceRepository{
		Id:        e.Id,
		Url:       e.Url,
		CloneUrl:  e.CloneUrl,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
	}
	if e.UpdatedAt != nil {
		v := *e.UpdatedAt
		n.UpdatedAt = &v
	}

	return n
}

type Job struct {
	Id           int32
	RepositoryId int32
	Command      string
	Target       string
	Active       bool
	AllRevision  bool
	GithubStatus bool
	CreatedAt    time.Time
	UpdatedAt    *time.Time

	Repository *SourceRepository

	mu   sync.Mutex
	mark *Job
}

func (e *Job) ResetMark() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.mark = e.Copy()
}

func (e *Job) IsChanged() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.RepositoryId != e.mark.RepositoryId ||
		e.Command != e.mark.Command ||
		e.Target != e.mark.Target ||
		e.Active != e.mark.Active ||
		e.AllRevision != e.mark.AllRevision ||
		e.GithubStatus != e.mark.GithubStatus ||
		!e.CreatedAt.Equal(e.mark.CreatedAt) ||
		((e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil))
}

func (e *Job) ChangedColumn() []ddl.Column {
	e.mu.Lock()
	defer e.mu.Unlock()

	res := make([]ddl.Column, 0)
	if e.RepositoryId != e.mark.RepositoryId {
		res = append(res, ddl.Column{Name: "repository_id", Value: e.RepositoryId})
	}
	if e.Command != e.mark.Command {
		res = append(res, ddl.Column{Name: "command", Value: e.Command})
	}
	if e.Target != e.mark.Target {
		res = append(res, ddl.Column{Name: "target", Value: e.Target})
	}
	if e.Active != e.mark.Active {
		res = append(res, ddl.Column{Name: "active", Value: e.Active})
	}
	if e.AllRevision != e.mark.AllRevision {
		res = append(res, ddl.Column{Name: "all_revision", Value: e.AllRevision})
	}
	if e.GithubStatus != e.mark.GithubStatus {
		res = append(res, ddl.Column{Name: "github_status", Value: e.GithubStatus})
	}
	if !e.CreatedAt.Equal(e.mark.CreatedAt) {
		res = append(res, ddl.Column{Name: "created_at", Value: e.CreatedAt})
	}
	if (e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil) {
		if e.UpdatedAt != nil {
			res = append(res, ddl.Column{Name: "updated_at", Value: *e.UpdatedAt})
		} else {
			res = append(res, ddl.Column{Name: "updated_at", Value: nil})
		}
	}

	return res
}

func (e *Job) Copy() *Job {
	n := &Job{
		Id:           e.Id,
		RepositoryId: e.RepositoryId,
		Command:      e.Command,
		Target:       e.Target,
		Active:       e.Active,
		AllRevision:  e.AllRevision,
		GithubStatus: e.GithubStatus,
		CreatedAt:    e.CreatedAt,
	}
	if e.UpdatedAt != nil {
		v := *e.UpdatedAt
		n.UpdatedAt = &v
	}

	return n
}

type Task struct {
	Id         int32
	JobId      int32
	Revision   string
	Success    bool
	LogFile    string
	Via        string
	FinishedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  *time.Time

	Job *Job

	mu   sync.Mutex
	mark *Task
}

func (e *Task) ResetMark() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.mark = e.Copy()
}

func (e *Task) IsChanged() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.JobId != e.mark.JobId ||
		e.Revision != e.mark.Revision ||
		e.Success != e.mark.Success ||
		e.LogFile != e.mark.LogFile ||
		e.Via != e.mark.Via ||
		((e.FinishedAt != nil && (e.mark.FinishedAt == nil || !e.FinishedAt.Equal(*e.mark.FinishedAt))) || (e.FinishedAt == nil && e.mark.FinishedAt != nil)) ||
		!e.CreatedAt.Equal(e.mark.CreatedAt) ||
		((e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil))
}

func (e *Task) ChangedColumn() []ddl.Column {
	e.mu.Lock()
	defer e.mu.Unlock()

	res := make([]ddl.Column, 0)
	if e.JobId != e.mark.JobId {
		res = append(res, ddl.Column{Name: "job_id", Value: e.JobId})
	}
	if e.Revision != e.mark.Revision {
		res = append(res, ddl.Column{Name: "revision", Value: e.Revision})
	}
	if e.Success != e.mark.Success {
		res = append(res, ddl.Column{Name: "success", Value: e.Success})
	}
	if e.LogFile != e.mark.LogFile {
		res = append(res, ddl.Column{Name: "log_file", Value: e.LogFile})
	}
	if e.Via != e.mark.Via {
		res = append(res, ddl.Column{Name: "via", Value: e.Via})
	}
	if (e.FinishedAt != nil && (e.mark.FinishedAt == nil || !e.FinishedAt.Equal(*e.mark.FinishedAt))) || (e.FinishedAt == nil && e.mark.FinishedAt != nil) {
		if e.FinishedAt != nil {
			res = append(res, ddl.Column{Name: "finished_at", Value: *e.FinishedAt})
		} else {
			res = append(res, ddl.Column{Name: "finished_at", Value: nil})
		}
	}
	if !e.CreatedAt.Equal(e.mark.CreatedAt) {
		res = append(res, ddl.Column{Name: "created_at", Value: e.CreatedAt})
	}
	if (e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil) {
		if e.UpdatedAt != nil {
			res = append(res, ddl.Column{Name: "updated_at", Value: *e.UpdatedAt})
		} else {
			res = append(res, ddl.Column{Name: "updated_at", Value: nil})
		}
	}

	return res
}

func (e *Task) Copy() *Task {
	n := &Task{
		Id:        e.Id,
		JobId:     e.JobId,
		Revision:  e.Revision,
		Success:   e.Success,
		LogFile:   e.LogFile,
		Via:       e.Via,
		CreatedAt: e.CreatedAt,
	}
	if e.FinishedAt != nil {
		v := *e.FinishedAt
		n.FinishedAt = &v
	}
	if e.UpdatedAt != nil {
		v := *e.UpdatedAt
		n.UpdatedAt = &v
	}

	return n
}

type TrustedUser struct {
	Id        int32
	GithubId  int64
	Username  string
	CreatedAt time.Time
	UpdatedAt *time.Time

	mu   sync.Mutex
	mark *TrustedUser
}

func (e *TrustedUser) ResetMark() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.mark = e.Copy()
}

func (e *TrustedUser) IsChanged() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.GithubId != e.mark.GithubId ||
		e.Username != e.mark.Username ||
		!e.CreatedAt.Equal(e.mark.CreatedAt) ||
		((e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil))
}

func (e *TrustedUser) ChangedColumn() []ddl.Column {
	e.mu.Lock()
	defer e.mu.Unlock()

	res := make([]ddl.Column, 0)
	if e.GithubId != e.mark.GithubId {
		res = append(res, ddl.Column{Name: "github_id", Value: e.GithubId})
	}
	if e.Username != e.mark.Username {
		res = append(res, ddl.Column{Name: "username", Value: e.Username})
	}
	if !e.CreatedAt.Equal(e.mark.CreatedAt) {
		res = append(res, ddl.Column{Name: "created_at", Value: e.CreatedAt})
	}
	if (e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil) {
		if e.UpdatedAt != nil {
			res = append(res, ddl.Column{Name: "updated_at", Value: *e.UpdatedAt})
		} else {
			res = append(res, ddl.Column{Name: "updated_at", Value: nil})
		}
	}

	return res
}

func (e *TrustedUser) Copy() *TrustedUser {
	n := &TrustedUser{
		Id:        e.Id,
		GithubId:  e.GithubId,
		Username:  e.Username,
		CreatedAt: e.CreatedAt,
	}
	if e.UpdatedAt != nil {
		v := *e.UpdatedAt
		n.UpdatedAt = &v
	}

	return n
}

type PermitPullRequest struct {
	Id         int32
	Repository string
	Number     int32
	CreatedAt  time.Time
	UpdatedAt  *time.Time

	mu   sync.Mutex
	mark *PermitPullRequest
}

func (e *PermitPullRequest) ResetMark() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.mark = e.Copy()
}

func (e *PermitPullRequest) IsChanged() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	return e.Repository != e.mark.Repository ||
		e.Number != e.mark.Number ||
		!e.CreatedAt.Equal(e.mark.CreatedAt) ||
		((e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil))
}

func (e *PermitPullRequest) ChangedColumn() []ddl.Column {
	e.mu.Lock()
	defer e.mu.Unlock()

	res := make([]ddl.Column, 0)
	if e.Repository != e.mark.Repository {
		res = append(res, ddl.Column{Name: "repository", Value: e.Repository})
	}
	if e.Number != e.mark.Number {
		res = append(res, ddl.Column{Name: "number", Value: e.Number})
	}
	if !e.CreatedAt.Equal(e.mark.CreatedAt) {
		res = append(res, ddl.Column{Name: "created_at", Value: e.CreatedAt})
	}
	if (e.UpdatedAt != nil && (e.mark.UpdatedAt == nil || !e.UpdatedAt.Equal(*e.mark.UpdatedAt))) || (e.UpdatedAt == nil && e.mark.UpdatedAt != nil) {
		if e.UpdatedAt != nil {
			res = append(res, ddl.Column{Name: "updated_at", Value: *e.UpdatedAt})
		} else {
			res = append(res, ddl.Column{Name: "updated_at", Value: nil})
		}
	}

	return res
}

func (e *PermitPullRequest) Copy() *PermitPullRequest {
	n := &PermitPullRequest{
		Id:         e.Id,
		Repository: e.Repository,
		Number:     e.Number,
		CreatedAt:  e.CreatedAt,
	}
	if e.UpdatedAt != nil {
		v := *e.UpdatedAt
		n.UpdatedAt = &v
	}

	return n
}
