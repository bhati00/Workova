// services/job_service.go
package job

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bhati00/workova/backend/constants"
)

// JobService interface defines business logic operations for jobs
type JobService interface {
	// Single job operations
	CreateJob(job *Job) error
	GetJobByID(id uint) (*Job, error)
	GetJobByJobID(jobID string) (*Job, error)
	UpdateJob(job *Job) error
	DeleteJob(id uint) error
	DeactivateJob(id uint) error

	// Batch operations
	CreateJobsBatch(jobs []Job) (*BatchResult, error)
	DeleteJobsBatch(ids []uint) (*BatchResult, error)
	DeleteJobsByJobIDsBatch(jobIDs []string) (*BatchResult, error)

	// Query operations
	GetAllJobs(page, pageSize int) (*PaginatedJobsResponse, error)
	GetActiveJobs(page, pageSize int) (*PaginatedJobsResponse, error)
	SearchJobs(params *JobSearchParams) (*PaginatedJobsResponse, error)
	GetJobStats() (*JobStatsResponse, error)
}

// PaginatedJobsResponse represents paginated jobs response
type PaginatedJobsResponse struct {
	Jobs        []Job `json:"jobs"`
	TotalCount  int64 `json:"total_count"`
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
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

// jobService implements JobService interface
type jobService struct {
	jobRepo JobRepository
}

// NewJobService creates a new job service instance
func NewJobService(jobRepo JobRepository) JobService {
	return &jobService{
		jobRepo: jobRepo,
	}
}

// CreateJob creates a new job with validation
func (s *jobService) CreateJob(job *Job) error {
	// Validate job data
	if err := s.validateJob(job); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Set default values
	s.setJobDefaults(job)

	// Validate and clean skills
	if err := s.validateAndCleanSkills(job); err != nil {
		return fmt.Errorf("skills validation failed: %w", err)
	}

	// Create the job
	if err := s.jobRepo.Create(job); err != nil {
		log.Printf("Failed to create job (JobID: %s): %v", job.JobID, err)
		return fmt.Errorf("failed to create job: %w", err)
	}

	log.Printf("Successfully created job (ID: %d, JobID: %s)", job.ID, job.JobID)
	return nil
}

// GetJobByID retrieves a job by its ID
func (s *jobService) GetJobByID(id uint) (*Job, error) {
	return s.jobRepo.GetByID(id)
}

// GetJobByJobID retrieves a job by its external job ID
func (s *jobService) GetJobByJobID(jobID string) (*Job, error) {
	if strings.TrimSpace(jobID) == "" {
		return nil, errors.New("job ID cannot be empty")
	}
	return s.jobRepo.GetByJobID(jobID)
}

// UpdateJob updates an existing job with validation
func (s *jobService) UpdateJob(job *Job) error {
	// Check if job exists
	_, err := s.jobRepo.GetByID(job.ID)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	// Validate updated job data
	if err := s.validateJob(job); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Validate and clean skills
	if err := s.validateAndCleanSkills(job); err != nil {
		return fmt.Errorf("skills validation failed: %w", err)
	}

	// Update timestamp
	job.LastUpdated = time.Now()

	if err := s.jobRepo.Update(job); err != nil {
		log.Printf("Failed to update job (ID: %d, JobID: %s): %v", job.ID, job.JobID, err)
		return fmt.Errorf("failed to update job: %w", err)
	}

	log.Printf("Successfully updated job (ID: %d, JobID: %s)", job.ID, job.JobID)
	return nil
}

// DeleteJob deletes a job
func (s *jobService) DeleteJob(id uint) error {
	// Check if job exists
	if _, err := s.jobRepo.GetByID(id); err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	if err := s.jobRepo.Delete(id); err != nil {
		log.Printf("Failed to delete job (ID: %d): %v", id, err)
		return fmt.Errorf("failed to delete job: %w", err)
	}

	log.Printf("Successfully deleted job (ID: %d)", id)
	return nil
}

// DeactivateJob deactivates a job instead of deleting it
func (s *jobService) DeactivateJob(id uint) error {
	job, err := s.jobRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("job not found: %w", err)
	}

	job.IsActive = false
	job.LastUpdated = time.Now()

	if err := s.jobRepo.Update(job); err != nil {
		log.Printf("Failed to deactivate job (ID: %d): %v", id, err)
		return fmt.Errorf("failed to deactivate job: %w", err)
	}

	log.Printf("Successfully deactivated job (ID: %d, JobID: %s)", job.ID, job.JobID)
	return nil
}

// CreateJobsBatch creates multiple jobs in batch with validation
func (s *jobService) CreateJobsBatch(jobs []Job) (*BatchResult, error) {
	if len(jobs) == 0 {
		return nil, errors.New("no jobs provided")
	}

	// Validate all jobs first
	validJobs := make([]Job, 0, len(jobs))
	invalidCount := 0

	for i, job := range jobs {
		// Validate each job
		if err := s.validateJob(&job); err != nil {
			log.Printf("Skipping invalid job at index %d (JobID: %s): %v", i, job.JobID, err)
			invalidCount++
			continue
		}

		// Set default values
		s.setJobDefaults(&job)

		// Validate and clean skills
		if err := s.validateAndCleanSkills(&job); err != nil {
			log.Printf("Skipping job with invalid skills at index %d (JobID: %s): %v", i, job.JobID, err)
			invalidCount++
			continue
		}

		validJobs = append(validJobs, job)
	}

	if len(validJobs) == 0 {
		return &BatchResult{
			TotalProcessed: len(jobs),
			Successful:     0,
			Failed:         len(jobs),
		}, errors.New("no valid jobs to process")
	}

	log.Printf("Processing batch of %d jobs (%d valid, %d invalid)", len(jobs), len(validJobs), invalidCount)

	// Process batch insert
	result, err := s.jobRepo.BatchInsert(validJobs)
	if err != nil {
		return nil, fmt.Errorf("batch insert failed: %w", err)
	}

	// Adjust counts to include validation failures
	result.Failed += invalidCount
	log.Printf("Batch insert completed: %d successful, %d failed out of %d total", result.Successful, result.Failed, result.TotalProcessed)

	return result, nil
}

// DeleteJobsBatch deletes multiple jobs by IDs
func (s *jobService) DeleteJobsBatch(ids []uint) (*BatchResult, error) {
	if len(ids) == 0 {
		return nil, errors.New("no job IDs provided")
	}

	log.Printf("Processing batch delete of %d jobs", len(ids))
	result, err := s.jobRepo.BatchDelete(ids)
	if err != nil {
		return nil, fmt.Errorf("batch delete failed: %w", err)
	}

	log.Printf("Batch delete completed: %d successful, %d failed", result.Successful, result.Failed)
	return result, nil
}

// DeleteJobsByJobIDsBatch deletes multiple jobs by external job IDs
func (s *jobService) DeleteJobsByJobIDsBatch(jobIDs []string) (*BatchResult, error) {
	if len(jobIDs) == 0 {
		return nil, errors.New("no job IDs provided")
	}

	// Filter out empty job IDs
	validJobIDs := make([]string, 0, len(jobIDs))
	for _, jobID := range jobIDs {
		if strings.TrimSpace(jobID) != "" {
			validJobIDs = append(validJobIDs, strings.TrimSpace(jobID))
		}
	}

	if len(validJobIDs) == 0 {
		return nil, errors.New("no valid job IDs provided")
	}

	log.Printf("Processing batch delete of %d jobs by JobID", len(validJobIDs))
	result, err := s.jobRepo.BatchDeleteByJobIDs(validJobIDs)
	if err != nil {
		return nil, fmt.Errorf("batch delete by job IDs failed: %w", err)
	}

	log.Printf("Batch delete by JobID completed: %d successful, %d failed", result.Successful, result.Failed)
	return result, nil
}

// GetAllJobs retrieves all jobs with pagination
func (s *jobService) GetAllJobs(page, pageSize int) (*PaginatedJobsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	jobs, err := s.jobRepo.GetAll(offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get jobs: %w", err)
	}

	// Get total count (you might want to add this method to repository)
	totalCount := int64(len(jobs)) // Placeholder - implement proper counting

	return &PaginatedJobsResponse{
		Jobs:        jobs,
		TotalCount:  totalCount,
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  int((totalCount + int64(pageSize) - 1) / int64(pageSize)),
	}, nil
}

// GetActiveJobs retrieves active jobs with pagination
func (s *jobService) GetActiveJobs(page, pageSize int) (*PaginatedJobsResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	jobs, err := s.jobRepo.GetActiveJobs(offset, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get active jobs: %w", err)
	}

	totalCount, err := s.jobRepo.CountActiveJobs()
	if err != nil {
		return nil, fmt.Errorf("failed to count active jobs: %w", err)
	}

	return &PaginatedJobsResponse{
		Jobs:        jobs,
		TotalCount:  totalCount,
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  int((totalCount + int64(pageSize) - 1) / int64(pageSize)),
	}, nil
}

// SearchJobs searches jobs based on parameters
func (s *jobService) SearchJobs(params *JobSearchParams) (*PaginatedJobsResponse, error) {
	if params.Limit < 1 || params.Limit > 100 {
		params.Limit = 20
	}
	if params.Offset < 0 {
		params.Offset = 0
	}

	jobs, totalCount, err := s.jobRepo.SearchJobs(params)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	currentPage := (params.Offset / params.Limit) + 1
	totalPages := int((totalCount + int64(params.Limit) - 1) / int64(params.Limit))

	return &PaginatedJobsResponse{
		Jobs:        jobs,
		TotalCount:  totalCount,
		CurrentPage: currentPage,
		PageSize:    params.Limit,
		TotalPages:  totalPages,
	}, nil
}

// GetJobStats returns job statistics
func (s *jobService) GetJobStats() (*JobStatsResponse, error) {
	// This is a placeholder implementation
	// You would implement actual stats gathering from the repository
	activeCount, err := s.jobRepo.CountActiveJobs()
	if err != nil {
		return nil, fmt.Errorf("failed to get job stats: %w", err)
	}

	return &JobStatsResponse{
		TotalJobs:      activeCount, // Placeholder
		ActiveJobs:     activeCount,
		InactiveJobs:   0, // Placeholder
		JobsByWorkMode: make(map[string]int64),
		JobsByWorkType: make(map[string]int64),
		JobsBySource:   make(map[string]int64),
		RecentJobs:     0, // Placeholder
	}, nil
}

// validateJob validates job data
func (s *jobService) validateJob(job *Job) error {
	if strings.TrimSpace(job.JobTitle) == "" {
		return errors.New("job title is required")
	}

	if strings.TrimSpace(job.JobID) == "" {
		return errors.New("job ID is required")
	}

	// Validate work mode
	if job.WorkMode != "" && !constants.IsValidWorkMode(job.WorkMode) {
		return fmt.Errorf("invalid work mode: %s", job.WorkMode)
	}

	// Validate work type
	if job.WorkType != "" && !constants.IsValidWorkType(job.WorkType) {
		return fmt.Errorf("invalid work type: %s", job.WorkType)
	}

	// Validate currency
	if job.Currency != "" && !constants.IsValidCurrency(job.Currency) {
		return fmt.Errorf("invalid currency: %s", job.Currency)
	}

	// Validate interview mode
	if job.InterviewMode != "" && !constants.IsValidInterviewMode(job.InterviewMode) {
		return fmt.Errorf("invalid interview mode: %s", job.InterviewMode)
	}

	// Validate pay schedule
	if job.PaySchedule != "" && !constants.IsValidPaySchedule(job.PaySchedule) {
		return fmt.Errorf("invalid pay schedule: %s", job.PaySchedule)
	}

	// Validate source
	if job.Source != "" && !constants.IsValidSource(job.Source) {
		return fmt.Errorf("invalid source: %s", job.Source)
	}

	// Validate salary range
	if job.CompensationMin != nil && job.CompensationMax != nil {
		if *job.CompensationMin > *job.CompensationMax {
			return errors.New("minimum compensation cannot be greater than maximum compensation")
		}
	}

	// Validate experience range
	if job.ExperienceMin != nil && job.ExperienceMax != nil {
		if *job.ExperienceMin > *job.ExperienceMax {
			return errors.New("minimum experience cannot be greater than maximum experience")
		}
	}

	return nil
}

// setJobDefaults sets default values for job fields
func (s *jobService) setJobDefaults(job *Job) {
	now := time.Now()

	job.LastUpdated = now

	// Set default values if not provided
	if job.Currency == "" {
		job.Currency = constants.CurrencyINR
	}

	if job.WorkMode == "" {
		job.WorkMode = constants.WorkModeOnsite
	}

	if job.WorkType == "" {
		job.WorkType = constants.WorkTypeFullTime
	}

	if job.PaySchedule == "" {
		job.PaySchedule = constants.PayScheduleMonthly
	}

	// Set default active status
	// Note: IsActive has a default value in the model, but we can ensure it here too
}

// validateAndCleanSkills validates and cleans job skills
func (s *jobService) validateAndCleanSkills(job *Job) error {
	if len(job.JobSkills) == 0 {
		return nil // Skills are optional
	}

	validSkills := make([]JobSkill, 0, len(job.JobSkills))

	for i, skill := range job.JobSkills {
		// Clean skill name
		skill.Skill = strings.TrimSpace(skill.Skill)
		if skill.Skill == "" {
			log.Printf("Skipping empty skill at index %d for job %s", i, job.JobID)
			continue
		}

		// Validate skill type
		if skill.Type != "" && !constants.IsValidSkillType(skill.Type) {
			log.Printf("Invalid skill type '%s' for skill '%s', setting to required", skill.Type, skill.Skill)
			skill.Type = constants.SkillTypeRequired
		}

		// Set default skill type if not provided
		if skill.Type == "" {
			skill.Type = constants.SkillTypeRequired
		}

		validSkills = append(validSkills, skill)
	}

	job.JobSkills = validSkills
	return nil
}
