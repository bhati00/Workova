package jobaggregator

import (
	"time"

	"github.com/bhati00/workova/backend/internal/job/model"
)

type FetchOptions struct {
	Pages       int               `json:"pages"`
	MaxJobs     int               `json:"max_jobs"`
	Location    string            `json:"location,omitempty"`
	Keywords    []string          `json:"keywords,omitempty"`
	JobType     string            `json:"job_type,omitempty"` // "full-time", "part-time", "contract", "remote"
	DatePosted  *time.Time        `json:"date_posted,omitempty"`
	ExtraParams map[string]string `json:"extra_params,omitempty"` // Platform-specific params
}

type JobAggregator interface {
	FetchJobs(options FetchOptions) ([]model.Job, error)
	TransformJob(rawJob []any) (model.Job, error)
}
