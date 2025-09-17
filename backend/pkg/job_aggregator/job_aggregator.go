package jobaggregator

import (
	"time"

	"github.com/bhati00/workova/backend/dtos"
)

type FetchOptions struct {
	Pages           int               `json:"pages"`
	MaxJobs         int               `json:"max_jobs"`
	Location        string            `json:"location,omitempty"`
	Keywords        []string          `json:"keywords,omitempty"`
	JobType         string            `json:"job_type,omitempty"` // "full-time", "part-time", "contract", "remote"
	DatePosted      *time.Time        `json:"date_posted,omitempty"`
	ExtraParams     map[string]string `json:"extra_params,omitempty"` // Platform-specific params
	VisaSponsorship bool              `json:"visa_sponsor_ship"`
}

type JobAggregator interface {
	FetchJobs(options FetchOptions) ([]dtos.JobRequest, error)
	RawJobtoDto(rawJob []any) (dtos.JobRequest, error)
}
