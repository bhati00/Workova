// services/job_service.go
package job

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bhati00/workova/backend/dtos"
	"github.com/bhati00/workova/backend/internal/job/model"
	"github.com/bhati00/workova/backend/internal/job/repository"
	"gorm.io/gorm"
)

// JobService interface defines business logic operations for jobs
type JobService interface {
	// Single job operations
	CreateJob(dtos.JobRequest) (*dtos.JobResponse, error)
	GetJobByID(id uint) (*model.Job, error)
	DeleteJob(id uint) error
	DeactivateJob(id uint) error

	// Batch operations
	DeleteJobsBatch(ids []uint) (*dtos.BatchResult, error)

	// Query operations
	GetAllJobs(page, pageSize int) (*dtos.PaginatedJobsResponse, error)
	SearchJobs(params *dtos.JobSearchParams) (*dtos.PaginatedJobsResponse, error)
	GetJobStats() (*dtos.JobStatsResponse, error)
}

// jobService implements JobService interface
type jobService struct {
	jobRepo      repository.JobRepository
	skillRepo    repository.SkillRepository
	categoryRepo repository.CategoryRepository
	locationRepo repository.LocationRepository
}

// NewJobService creates a new job service instance
func NewJobService(jobRepo repository.JobRepository, skillRepo repository.SkillRepository, categoryRepo repository.CategoryRepository, locationRep repository.LocationRepository) JobService {
	return &jobService{
		jobRepo:      jobRepo,
		skillRepo:    skillRepo,
		categoryRepo: categoryRepo,
		locationRepo: locationRep,
	}
}

// CreateJob creates a new job with validation
func (s *jobService) CreateJob(jobRequest dtos.JobRequest) (*dtos.JobResponse, error) {
	job, err := ConvertJobRequest(jobRequest)
	if err != nil {
		return nil, fmt.Errorf("invalid job data: %w", err)
	}

	job, err = s.jobRepo.Create(job)
	if err != nil {
		log.Printf("Failed to create job (JobTitle: %s): %v", jobRequest.Title, err)
		return nil, fmt.Errorf("failed to create job: %w", err)
	}

	// adding job skills
	skills := jobRequest.Skills
	if len(skills) > 0 {
		for _, skillName := range skills {
			skillObj, err := s.skillRepo.GetByName(skillName)
			if err != nil {
				skillObj, err = s.skillRepo.Create(&model.Skill{Name: skillName})
				if err != nil {
					continue
				}
			}
			jobSkill := model.JobSkill{
				JobID:   job.ID,
				SkillID: skillObj.ID,
			}
			s.jobRepo.CreateJobSkill(&jobSkill)
		}
	}
	// adding job locations
	CountryIso := jobRequest.CountryIso
	if CountryIso != "" {
		countryObj, err := s.locationRepo.GetCountryByISO(CountryIso)
		if err != nil {
			// i need to map the country ISO to the country name
			countryObj, _ = s.locationRepo.CreateCountry(&model.Country{Name: CountryIso, ISO: CountryIso})
		}
		jobLocation := model.JobLocation{
			JobID:     job.ID,
			CountryID: countryObj.ID,
			City:      jobRequest.City,
		}
		s.locationRepo.CreateJobLocation(&jobLocation)
	}
	// add categories
	category := jobRequest.Category
	if category != "" {
		categoryObj, err := s.categoryRepo.GetCategoryByName(category)
		if err != nil {
			categoryObj, _ = s.categoryRepo.Create(&model.Category{Name: category})
		}
		JobCategory := model.JobCategory{
			JobID:      job.ID,
			CategoryID: categoryObj.ID,
		}
		s.categoryRepo.CreateJobCategory(&JobCategory)
	}
	response := &dtos.JobResponse{
		ID:            job.ID,
		ExternalJobID: *job.ExternalJobID,
		Title:         job.Title,
		Slug:          *job.Slug,
		CreatedAt:     job.CreatedAt,
		Message:       "Job created successfully",
	}

	return response, nil
}

// GetJobByID retrieves a job by its ID
func (s *jobService) GetJobByID(id uint) (*model.Job, error) {
	return s.jobRepo.GetByID(id)
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

	job.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
	if err := s.jobRepo.Update(job); err != nil {
		log.Printf("Failed to deactivate job (ID: %d): %v", id, err)
		return fmt.Errorf("failed to deactivate job: %w", err)
	}

	log.Printf("Successfully deactivated job (ID: %d, JobTitle: %s)", job.ID, job.Title)
	return nil
}

// DeleteJobsBatch deletes multiple jobs by IDs
func (s *jobService) DeleteJobsBatch(ids []uint) (*dtos.BatchResult, error) {
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

// GetAllJobs retrieves all jobs with pagination
func (s *jobService) GetAllJobs(page, pageSize int) (*dtos.PaginatedJobsResponse, error) {
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

	return &dtos.PaginatedJobsResponse{
		Jobs:        jobs,
		TotalCount:  totalCount,
		CurrentPage: page,
		PageSize:    pageSize,
		TotalPages:  int((totalCount + int64(pageSize) - 1) / int64(pageSize)),
	}, nil
}

// SearchJobs searches jobs based on parameters
func (s *jobService) SearchJobs(params *dtos.JobSearchParams) (*dtos.PaginatedJobsResponse, error) {
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

	return &dtos.PaginatedJobsResponse{
		Jobs:        jobs,
		TotalCount:  totalCount,
		CurrentPage: currentPage,
		PageSize:    params.Limit,
		TotalPages:  totalPages,
	}, nil
}

// GetJobStats returns job statistics
func (s *jobService) GetJobStats() (*dtos.JobStatsResponse, error) {
	// This is a placeholder implementation
	// You would implement actual stats gathering from the repository
	activeCount, err := s.jobRepo.CountActiveJobs()
	if err != nil {
		return nil, fmt.Errorf("failed to get job stats: %w", err)
	}

	return &dtos.JobStatsResponse{
		TotalJobs:      activeCount, // Placeholder
		ActiveJobs:     activeCount,
		InactiveJobs:   0, // Placeholder
		JobsByWorkMode: make(map[string]int64),
		JobsByWorkType: make(map[string]int64),
		JobsBySource:   make(map[string]int64),
		RecentJobs:     0, // Placeholder
	}, nil
}

func ConvertJobRequest(jobDto dtos.JobRequest) (*model.Job, error) {
	// 1. Basic validations
	if jobDto.Title == "" {
		return nil, errors.New("title is required")
	}
	if jobDto.CompanyName == "" {
		return nil, errors.New("company name is required")
	}
	if jobDto.JobType == 0 {
		return nil, errors.New("job type is required")
	}
	if jobDto.WorkMode == 0 {
		return nil, errors.New("work mode is required")
	}

	// 2. Parse dates
	var postedDate *time.Time
	if jobDto.PostedDate != nil {
		t, err := time.Parse(time.RFC3339, *jobDto.PostedDate)
		if err != nil {
			return nil, errors.New("invalid posted_date format (use RFC3339)")
		}
		postedDate = &t
	}

	// 3. Helper function to safely dereference pointers
	var externalJobID, slug, description, applicationURL, industry, department, educationLevel *string

	var salaryMin, salaryMax *int
	var isRemote, visaSponsorship *bool

	if jobDto.ExternalJobID != nil {
		externalJobID = jobDto.ExternalJobID
	}
	if jobDto.Slug != nil {
		slug = jobDto.Slug
	}
	if jobDto.Description != nil {
		description = jobDto.Description
	}

	if jobDto.ApplicationURL != nil {
		applicationURL = jobDto.ApplicationURL
	}
	if jobDto.Industry != nil {
		industry = jobDto.Industry
	}
	if jobDto.Department != nil {
		department = jobDto.Department
	}
	if jobDto.EducationLevel != nil {
		educationLevel = jobDto.EducationLevel
	}

	salaryMin = jobDto.SalaryMin
	salaryMax = jobDto.SalaryMax
	isRemote = jobDto.IsRemote
	visaSponsorship = jobDto.VisaSponsorship

	// 4. Convert into Job model
	job := model.Job{
		ExternalJobID: externalJobID,
		Slug:          slug,
		Title:         jobDto.Title,
		Description:   description,
		CompanyName:   jobDto.CompanyName,

		JobType:         jobDto.JobType,
		WorkMode:        jobDto.WorkMode,
		ExperienceLevel: jobDto.ExperienceLevel,

		SalaryMin:      salaryMin,
		SalaryMax:      salaryMax,
		SalaryCurrency: jobDto.SalaryCurrency,
		IsRemote:       isRemote,

		ApplicationURL:  applicationURL,
		Source:          jobDto.Source,
		Industry:        industry,
		Department:      department,
		VisaSponsorship: visaSponsorship,
		EducationLevel:  educationLevel,
		PostedDate:      postedDate,
	}
	return &job, nil
}
