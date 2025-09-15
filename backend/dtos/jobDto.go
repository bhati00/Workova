package dtos

import (
	"time"

	"github.com/bhati00/workova/backend/constant"
	"github.com/bhati00/workova/backend/internal/job/model"
)

type JobRequest struct {
	ExternalJobID   *string                   `json:"external_job_id,omitempty" example:"ext-12345"`
	Slug            *string                   `json:"slug,omitempty" example:"software-engineer-12345"`
	Title           string                    `json:"title" example:"Software Engineer"`
	Description     *string                   `json:"description" example:"Job description here"`
	CompanyName     string                    `json:"company" example:"Tech Corp"`
	CountryIso      string                    `json:"country" example:"USA"`
	City            *string                   `json:"city,omitempty" example:"San Francisco"`
	JobType         constant.JobType          `json:"job_type" example:"full-time"` // "full-time", "part-time", "contract", "remote"
	SalaryMin       *int                      `json:"salary_min,omitempty" example:"60000"`
	SalaryMax       *int                      `json:"salary_max,omitempty" example:"120000"`
	SalaryCurrency  constant.Currency         `json:"salary_currency,omitempty" example:"USD"`
	IsRemote        *bool                     `json:"is_remote,omitempty" example:"false"`
	Skills          []string                  `json:"skills,omitempty" example:"Go, Docker, Kubernetes"`
	Category        string                    `json:"categories,omitempty" example:"Engineering, IT"`
	PostedDate      *string                   `json:"posted_at,omitempty" example:"2023-10-01T12:00:00Z"`
	ApplicationURL  *string                   `json:"apply_url,omitempty" example:"https://techcorp.com/careers/12345"`
	Source          string                    `json:"source,omitempty" example:"LinkedIn"`
	Tags            []string                  `json:"tags,omitempty" example:"urgent,remote"`
	Industry        *string                   `json:"industry,omitempty" example:"Information Technology"`
	Department      *string                   `json:"department,omitempty" example:"Engineering"`
	VisaSponsorship *bool                     `json:"visa_sponsorship,omitempty" example:"false"`
	EducationLevel  *string                   `json:"education_level,omitempty" example:"Bachelor's"`
	ExperienceLevel *constant.ExperienceLevel `json:"experience_level,omitempty" example:"Mid-Level"`
	WorkMode        constant.WorkMode         `json:"work_mode,omitempty" example:"remote"` // "onsite", "remote", "hybrid"
}

// PaginatedJobsResponse represents paginated jobs response
type PaginatedJobsResponse struct {
	Jobs        []model.Job `json:"jobs"`
	TotalCount  int64       `json:"total_count"`
	CurrentPage int         `json:"current_page"`
	PageSize    int         `json:"page_size"`
	TotalPages  int         `json:"total_pages"`
}

// JobStatsResponse represents job statistics
type JobStatsResponse struct {
	TotalJobs      int64            `json:"total_jobs"`
	ActiveJobs     int64            `json:"active_jobs"`
	InactiveJobs   int64            `json:"inactive_jobs"`
	JobsByWorkMode map[string]int64 `json:"jobs_by_work_mode"`
	JobsByWorkType map[string]int64 `json:"jobs_by_work_type"`
	JobsBySource   map[string]int64 `json:"jobs_by_source"`
	RecentJobs     int64            `json:"recent_jobs"` // Last 7 days
}

type JobSearchParams struct {
	Query            string                     `json:"query"`
	Skills           []string                   `json:"skills"`
	WorkMode         []constant.WorkMode        `json:"work_mode"`
	JobType          []constant.JobType         `json:"job_type"`
	ExperienceLevel  []constant.ExperienceLevel `json:"experience_level"`
	MinSalary        *int                       `json:"min_salary"`
	MaxSalary        *int                       `json:"max_salary"`
	Currency         string                     `json:"currency"`
	IsRemote         *bool                      `json:"is_remote"`
	VisaSponsorship  *bool                      `json:"visa_sponsorship"`
	IsUrgent         *bool                      `json:"is_urgent"`
	CompanySize      []string                   `json:"company_size"`
	Industry         []string                   `json:"industry"`
	Department       []string                   `json:"department"`
	EducationLevel   []string                   `json:"education_level"`
	TravelRequired   []string                   `json:"travel_required"`
	Location         []string                   `json:"location"` // For filtering by job locations
	Source           []string                   `json:"source"`
	PostedAfter      *time.Time                 `json:"posted_after"`
	PostedBefore     *time.Time                 `json:"posted_before"`
	SalaryPeriod     []string                   `json:"salary_period"`
	ContractDuration *int                       `json:"contract_duration"`
	Offset           int                        `json:"offset"`
	Limit            int                        `json:"limit"`
	SortBy           string                     `json:"sort_by"`    // "created_at", "posted_date", "salary_max", etc.
	SortOrder        string                     `json:"sort_order"` // "asc", "desc"
}
