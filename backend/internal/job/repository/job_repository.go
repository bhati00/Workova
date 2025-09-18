// repositories/job_repository.go
package repository

import (
	"log"

	"github.com/bhati00/workova/backend/dtos"
	"github.com/bhati00/workova/backend/internal/job/model"
	"gorm.io/gorm"
)

const MaxBatchSize = 100

// JobRepository interface defines all job-related database operations
type JobRepository interface {
	// Single operations
	Create(job *model.Job) (*model.Job, error)
	GetByID(id uint) (*model.Job, error)
	Update(job *model.Job) error
	Delete(id uint) error
	SoftDelete(id uint) error
	CreateJobLocation(location *model.JobLocation) (*model.JobLocation, error)
	CreateJobCategory(category *model.JobCategory) (*model.JobCategory, error)
	CreateJobSkill(skill *model.JobSkill) (*model.JobSkill, error)
	// IsDuplicateJob(externalJobID *string, slug *string) (bool, error)

	// Batch operations
	BatchDelete(ids []uint) (*dtos.BatchResult, error)

	// Query operations
	GetAll(offset, limit int) ([]model.Job, error)
	SearchJobs(searchParams *dtos.JobSearchParams) ([]model.Job, int64, error)
	CountActiveJobs() (int64, error)
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
func (r *jobRepository) Create(job *model.Job) (*model.Job, error) {
	if err := r.db.Create(job).Error; err != nil {
		return nil, err
	}
	return job, nil
}

// Create Job location record
func (r *jobRepository) CreateJobLocation(jobLocation *model.JobLocation) (*model.JobLocation, error) {
	if err := r.db.Create(jobLocation).Error; err != nil {
		return nil, err
	}
	return jobLocation, nil
}

// Create Job category record
func (r *jobRepository) CreateJobCategory(jobCategory *model.JobCategory) (*model.JobCategory, error) {
	if err := r.db.Create(jobCategory).Error; err != nil {
		return nil, err
	}
	return jobCategory, nil
}

// Create job skill record
func (r *jobRepository) CreateJobSkill(jobSkill *model.JobSkill) (*model.JobSkill, error) {
	if err := r.db.Create(jobSkill).Error; err != nil {
		return nil, err
	}
	return jobSkill, nil
}

// GetByID retrieves a job by its primary key ID
func (r *jobRepository) GetByID(id uint) (*model.Job, error) {
	var job model.Job
	err := r.db.Preload("job_skills").
		Preload("job_categories").
		Preload("job_locations").
		First(&job, id).Error
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// Update updates an existing job record
func (r *jobRepository) Update(job *model.Job) error {
	return r.db.Save(job).Error
}

// Delete hard deletes a job record
func (r *jobRepository) Delete(id uint) error {
	return r.db.Select("job_skills", "job_categories", "job_locations").Delete(&model.Job{}, id).Error
}

// SoftDelete soft deletes a job record
func (r *jobRepository) SoftDelete(id uint) error {
	return r.db.Delete(&model.Job{}, id).Error
}

// BatchDelete deletes multiple jobs by their primary key IDs
func (r *jobRepository) BatchDelete(ids []uint) (*dtos.BatchResult, error) {
	result := &dtos.BatchResult{
		TotalProcessed: len(ids),
		Successful:     0,
		Failed:         0,
		Errors:         make([]dtos.BatchError, 0),
	}

	// Process in batches
	for i := 0; i < len(ids); i += MaxBatchSize {
		end := i + MaxBatchSize
		if end > len(ids) {
			end = len(ids)
		}

		batch := ids[i:end]

		// Try batch delete first
		tx := r.db.Select("job_skills", "job_categories", "job_locations").Delete(&model.Job{}, batch)
		if tx.Error != nil {
			// If batch fails, try individual deletes
			for j, id := range batch {
				if err := r.db.Select("job_skills", "job_categories", "job_locations").Delete(&model.Job{}, id).Error; err != nil {
					result.Failed++
					result.Errors = append(result.Errors, dtos.BatchError{
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

// GetAll retrieves all jobs with pagination
func (r *jobRepository) GetAll(offset, limit int) ([]model.Job, error) {
	var jobs []model.Job
	err := r.db.Preload("job_skills").
		Preload("job_categories").
		Preload("job_locations").Offset(offset).Limit(limit).Find(&jobs).Error
	return jobs, err
}

// CountActiveJobs returns the count of active jobs
func (r *jobRepository) CountActiveJobs() (int64, error) {
	var count int64
	err := r.db.Model(&model.Job{}).Where("is_active = ?", true).Count(&count).Error
	return count, err
}

// SearchJobs searches jobs based on various parameters
func (r *jobRepository) SearchJobs(params *dtos.JobSearchParams) ([]model.Job, int64, error) {
	var jobs []model.Job
	var total int64

	query := r.db.Preload("job_skills").
		Preload("job_categories").
		Preload("job_locations")

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
func (r *jobRepository) applySearchFilters(query *gorm.DB, params *dtos.JobSearchParams) *gorm.DB {
	// Text search in multiple fields
	if params.Query != "" {
		searchTerm := "%" + params.Query + "%"
		query = query.Where(
			"title ILIKE ? OR description ILIKE ? OR company_name ILIKE ? OR summary ILIKE ? OR keywords ILIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm, searchTerm,
		)
	}

	// Work mode filter
	if len(params.WorkMode) > 0 {
		query = query.Where("work_mode IN ?", params.WorkMode)
	}

	// Job type filter
	if len(params.JobType) > 0 {
		query = query.Where("job_type IN ?", params.JobType)
	}

	// Experience level filter
	if len(params.ExperienceLevel) > 0 {
		query = query.Where("experience_level IN ?", params.ExperienceLevel)
	}

	// Salary filters
	if params.MinSalary != nil {
		query = query.Where("salary_min >= ? OR salary_max >= ?", *params.MinSalary, *params.MinSalary)
	}
	if params.MaxSalary != nil {
		query = query.Where("salary_max <= ? OR salary_min <= ?", *params.MaxSalary, *params.MaxSalary)
	}

	// Currency filter
	if params.Currency != "" {
		query = query.Where("salary_currency = ?", params.Currency)
	}

	// Salary period filter
	if len(params.SalaryPeriod) > 0 {
		query = query.Where("salary_period IN ?", params.SalaryPeriod)
	}

	// Boolean filters
	if params.IsRemote != nil {
		query = query.Where("is_remote = ?", *params.IsRemote)
	}
	if params.VisaSponsorship != nil {
		query = query.Where("visa_sponsorship = ?", *params.VisaSponsorship)
	}
	if params.IsUrgent != nil {
		query = query.Where("is_urgent = ?", *params.IsUrgent)
	}

	// Company and industry filters
	if len(params.CompanySize) > 0 {
		query = query.Where("company_size IN ?", params.CompanySize)
	}
	if len(params.Industry) > 0 {
		query = query.Where("industry IN ?", params.Industry)
	}
	if len(params.Department) > 0 {
		query = query.Where("department IN ?", params.Department)
	}

	// Education filters
	if len(params.EducationLevel) > 0 {
		query = query.Where("education_level IN ?", params.EducationLevel)
	}

	// Travel requirement filter
	if len(params.TravelRequired) > 0 {
		query = query.Where("travel_required IN ?", params.TravelRequired)
	}

	// Source filter
	if len(params.Source) > 0 {
		query = query.Where("source IN ?", params.Source)
	}

	// Date range filters
	if params.PostedAfter != nil {
		query = query.Where("posted_date >= ?", *params.PostedAfter)
	}
	if params.PostedBefore != nil {
		query = query.Where("posted_date <= ?", *params.PostedBefore)
	}

	// Contract duration filter
	if params.ContractDuration != nil {
		query = query.Where("contract_months_duration = ?", *params.ContractDuration)
	}

	// Skills filter using junction table
	if len(params.Skills) > 0 {
		query = query.Joins("JOIN job_skills ON jobs.id = job_skills.job_id").
			Joins("JOIN skills ON job_skills.skill_id = skills.id").
			Where("skills.name IN ? OR skills.slug IN ?", params.Skills, params.Skills).
			Distinct()
	}

	// Location filter using junction table
	if len(params.Location) > 0 {
		query = query.Joins("JOIN job_locations ON jobs.id = job_locations.job_id").
			Where("job_locations.city IN ? OR job_locations.state IN ? OR job_locations.country IN ?",
				params.Location, params.Location, params.Location).
			Distinct()
	}

	// Apply sorting
	if params.SortBy != "" {
		order := "DESC"
		if params.SortOrder == "asc" {
			order = "ASC"
		}

		switch params.SortBy {
		case "created_at", "posted_date", "updated_at":
			query = query.Order(params.SortBy + " " + order)
		case "salary_max", "salary_min":
			query = query.Order(params.SortBy + " " + order + " NULLS LAST")
		case "company_name", "title":
			query = query.Order(params.SortBy + " " + order)
		default:
			query = query.Order("created_at DESC") // Default sorting
		}
	} else {
		query = query.Order("created_at DESC") // Default sorting
	}

	return query
}

// func (r *jobRepository) IsDuplicateJob(externalJobID *string, slug *string) (bool, error) {
// 	var count int64
// 	query := r.db.Model(&model.Job{})
// 	if externalJobID != nil && *externalJobID != "" {
// 		query = query.Where("external_job_id = ?", *externalJobID)
// 	}
// 	if slug != nil && *slug != "" {
// 		query = query.Or("slug = ?", *slug)
// 	}
// 	err := query.Count(&count).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	return count > 0, nil
// }
