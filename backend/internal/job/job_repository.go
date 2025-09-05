// repositories/job_repository.go
package job

import (
	"log"

	"gorm.io/gorm"
)

const MaxBatchSize = 100

// JobRepository interface defines all job-related database operations
type JobRepository interface {
	// Single operations
	Create(job *Job) error
	GetByID(id uint) (*Job, error)
	GetByJobID(jobID string) (*Job, error)
	Update(job *Job) error
	Delete(id uint) error
	SoftDelete(id uint) error

	// Batch operations
	BatchInsert(jobs []Job) (*BatchResult, error)
	BatchDelete(ids []uint) (*BatchResult, error)
	BatchDeleteByJobIDs(jobIDs []string) (*BatchResult, error)

	// Query operations
	GetAll(offset, limit int) ([]Job, error)
	GetActiveJobs(offset, limit int) ([]Job, error)
	SearchJobs(searchParams *JobSearchParams) ([]Job, int64, error)
	CountActiveJobs() (int64, error)
}

// BatchResult represents the result of batch operations
type BatchResult struct {
	TotalProcessed int          `json:"total_processed"`
	Successful     int          `json:"successful"`
	Failed         int          `json:"failed"`
	Errors         []BatchError `json:"errors,omitempty"`
}

// BatchError represents individual batch operation errors
type BatchError struct {
	Index int    `json:"index"`
	JobID string `json:"job_id,omitempty"`
	ID    uint   `json:"id,omitempty"`
	Error string `json:"error"`
}

// JobSearchParams defines search parameters for jobs
type JobSearchParams struct {
	Query     string   `json:"query"`
	Skills    []string `json:"skills"`
	WorkMode  []string `json:"work_mode"`
	WorkType  []string `json:"work_type"`
	MinSalary *int     `json:"min_salary"`
	MaxSalary *int     `json:"max_salary"`
	Currency  string   `json:"currency"`
	IsActive  *bool    `json:"is_active"`
	Source    []string `json:"source"`
	Offset    int      `json:"offset"`
	Limit     int      `json:"limit"`
}

// jobRepository implements JobRepository interface
type jobRepository struct {
	db *gorm.DB
}

// NewJobRepository creates a new job repository instance
func NewJobRepository(db *gorm.DB) JobRepository {
	return &jobRepository{db: db}
}

// Create creates a new job record
func (r *jobRepository) Create(job *Job) error {
	return r.db.Create(job).Error
}

// GetByID retrieves a job by its primary key ID
func (r *jobRepository) GetByID(id uint) (*Job, error) {
	var job Job
	err := r.db.Preload("JobSkills").First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// GetByJobID retrieves a job by its external job ID
func (r *jobRepository) GetByJobID(jobID string) (*Job, error) {
	var job Job
	err := r.db.Preload("JobSkills").Where("job_id = ?", jobID).First(&job).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// Update updates an existing job record
func (r *jobRepository) Update(job *Job) error {
	return r.db.Save(job).Error
}

// Delete hard deletes a job record
func (r *jobRepository) Delete(id uint) error {
	return r.db.Select("JobSkills").Delete(&Job{}, id).Error
}

// SoftDelete soft deletes a job record
func (r *jobRepository) SoftDelete(id uint) error {
	return r.db.Delete(&Job{}, id).Error
}

// BatchInsert inserts multiple jobs in batches
func (r *jobRepository) BatchInsert(jobs []Job) (*BatchResult, error) {
	result := &BatchResult{
		TotalProcessed: len(jobs),
		Successful:     0,
		Failed:         0,
		Errors:         make([]BatchError, 0),
	}

	// Process jobs in batches of MaxBatchSize
	for i := 0; i < len(jobs); i += MaxBatchSize {
		end := i + MaxBatchSize
		if end > len(jobs) {
			end = len(jobs)
		}

		batch := jobs[i:end]
		batchResult := r.processBatchInsert(batch, i)

		result.Successful += batchResult.Successful
		result.Failed += batchResult.Failed
		result.Errors = append(result.Errors, batchResult.Errors...)
	}

	return result, nil
}

// processBatchInsert processes a single batch of job insertions
func (r *jobRepository) processBatchInsert(batch []Job, startIndex int) *BatchResult {
	result := &BatchResult{
		TotalProcessed: len(batch),
		Successful:     0,
		Failed:         0,
		Errors:         make([]BatchError, 0),
	}

	// Try to insert the entire batch first
	if err := r.db.CreateInBatches(&batch, MaxBatchSize).Error; err != nil {
		// If batch insert fails, try individual inserts
		log.Printf("Batch insert failed, falling back to individual inserts: %v", err)

		for i, job := range batch {
			if err := r.db.Create(&job).Error; err != nil {
				result.Failed++
				result.Errors = append(result.Errors, BatchError{
					Index: startIndex + i,
					JobID: job.JobID,
					Error: err.Error(),
				})
				log.Printf("Failed to insert job at index %d (JobID: %s): %v", startIndex+i, job.JobID, err)
			} else {
				result.Successful++
			}
		}
	} else {
		result.Successful = len(batch)
	}

	return result
}

// BatchDelete deletes multiple jobs by their primary key IDs
func (r *jobRepository) BatchDelete(ids []uint) (*BatchResult, error) {
	result := &BatchResult{
		TotalProcessed: len(ids),
		Successful:     0,
		Failed:         0,
		Errors:         make([]BatchError, 0),
	}

	// Process in batches
	for i := 0; i < len(ids); i += MaxBatchSize {
		end := i + MaxBatchSize
		if end > len(ids) {
			end = len(ids)
		}

		batch := ids[i:end]

		// Try batch delete first
		tx := r.db.Select("JobSkills").Delete(&Job{}, batch)
		if tx.Error != nil {
			// If batch fails, try individual deletes
			for j, id := range batch {
				if err := r.db.Select("JobSkills").Delete(&Job{}, id).Error; err != nil {
					result.Failed++
					result.Errors = append(result.Errors, BatchError{
						Index: i + j,
						ID:    id,
						Error: err.Error(),
					})
					log.Printf("Failed to delete job with ID %d: %v", id, err)
				} else {
					result.Successful++
				}
			}
		} else {
			result.Successful += int(tx.RowsAffected)
		}
	}

	return result, nil
}

// BatchDeleteByJobIDs deletes multiple jobs by their external job IDs
func (r *jobRepository) BatchDeleteByJobIDs(jobIDs []string) (*BatchResult, error) {
	result := &BatchResult{
		TotalProcessed: len(jobIDs),
		Successful:     0,
		Failed:         0,
		Errors:         make([]BatchError, 0),
	}

	// Process in batches
	for i := 0; i < len(jobIDs); i += MaxBatchSize {
		end := i + MaxBatchSize
		if end > len(jobIDs) {
			end = len(jobIDs)
		}

		batch := jobIDs[i:end]

		// Try batch delete first
		tx := r.db.Select("JobSkills").Where("job_id IN ?", batch).Delete(&Job{})
		if tx.Error != nil {
			// If batch fails, try individual deletes
			for j, jobID := range batch {
				if err := r.db.Select("JobSkills").Where("job_id = ?", jobID).Delete(&Job{}).Error; err != nil {
					result.Failed++
					result.Errors = append(result.Errors, BatchError{
						Index: i + j,
						JobID: jobID,
						Error: err.Error(),
					})
					log.Printf("Failed to delete job with JobID %s: %v", jobID, err)
				} else {
					result.Successful++
				}
			}
		} else {
			result.Successful += int(tx.RowsAffected)
		}
	}

	return result, nil
}

// GetAll retrieves all jobs with pagination
func (r *jobRepository) GetAll(offset, limit int) ([]Job, error) {
	var jobs []Job
	err := r.db.Preload("JobSkills").Offset(offset).Limit(limit).Find(&jobs).Error
	return jobs, err
}

// GetActiveJobs retrieves all active jobs with pagination
func (r *jobRepository) GetActiveJobs(offset, limit int) ([]Job, error) {
	var jobs []Job
	err := r.db.Preload("JobSkills").Where("is_active = ?", true).Offset(offset).Limit(limit).Find(&jobs).Error
	return jobs, err
}

// CountActiveJobs returns the count of active jobs
func (r *jobRepository) CountActiveJobs() (int64, error) {
	var count int64
	err := r.db.Model(&Job{}).Where("is_active = ?", true).Count(&count).Error
	return count, err
}

// SearchJobs searches jobs based on various parameters
func (r *jobRepository) SearchJobs(params *JobSearchParams) ([]Job, int64, error) {
	var jobs []Job
	var total int64

	query := r.db.Model(&Job{}).Preload("JobSkills")

	// Apply filters
	query = r.applySearchFilters(query, params)

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and get results
	if err := query.Offset(params.Offset).Limit(params.Limit).Find(&jobs).Error; err != nil {
		return nil, 0, err
	}

	return jobs, total, nil
}

// applySearchFilters applies search filters to the query
func (r *jobRepository) applySearchFilters(query *gorm.DB, params *JobSearchParams) *gorm.DB {
	// Text search in job title and description
	if params.Query != "" {
		searchTerm := "%" + params.Query + "%"
		query = query.Where("job_title ILIKE ? OR description ILIKE ?", searchTerm, searchTerm)
	}

	// Work mode filter
	if len(params.WorkMode) > 0 {
		query = query.Where("work_mode IN ?", params.WorkMode)
	}

	// Work type filter
	if len(params.WorkType) > 0 {
		query = query.Where("work_type IN ?", params.WorkType)
	}

	// Salary range filter
	if params.MinSalary != nil {
		query = query.Where("compensation_min >= ? OR compensation_max >= ?", *params.MinSalary, *params.MinSalary)
	}
	if params.MaxSalary != nil {
		query = query.Where("compensation_max <= ? OR compensation_min <= ?", *params.MaxSalary, *params.MaxSalary)
	}

	// Currency filter
	if params.Currency != "" {
		query = query.Where("currency = ?", params.Currency)
	}

	// Active status filter
	if params.IsActive != nil {
		query = query.Where("is_active = ?", *params.IsActive)
	}

	// Source filter
	if len(params.Source) > 0 {
		query = query.Where("source IN ?", params.Source)
	}

	// Skills filter (if jobs have specific skills)
	if len(params.Skills) > 0 {
		query = query.Joins("JOIN job_skills ON jobs.id = job_skills.job_id").
			Where("job_skills.skill IN ?", params.Skills).
			Distinct()
	}

	return query
}
