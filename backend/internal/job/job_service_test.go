// internal/job/service/job_service_test.go
package job

import (
	"errors"
	"testing"

	"github.com/bhati00/workova/backend/constant"
	"github.com/bhati00/workova/backend/dtos"
	"github.com/bhati00/workova/backend/internal/job/mocks"
	"github.com/bhati00/workova/backend/internal/job/model"
	"github.com/bhati00/workova/backend/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Test helper functions
func createValidJobRequest() dtos.JobRequest {
	return dtos.JobRequest{
		ExternalJobID: utils.String("ext-123"),
		Slug:          utils.String("software-engineer-job"),
		Title:         "Software Engineer",
		Description:   utils.String("Great job opportunity"),
		CompanyName:   "Tech Corp",
		JobType:       constant.JobTypeContract, // assuming 1 is valid
		WorkMode:      constant.WorkModeOnsite,  // assuming 1 is valid
		Skills:        []string{"Go", "Python"},
		CountryIso:    "US",
		City:          utils.String("New York"),
		Category:      "Engineering",
	}
}

func createInvalidJobRequest() dtos.JobRequest {
	return dtos.JobRequest{
		// Missing required fields
		Title:       "",
		CompanyName: "",
	}
}

// Test CreateJob function
func TestJobService_CreateJob(t *testing.T) {
	tests := []struct {
		name          string
		jobRequest    dtos.JobRequest
		setupMocks    func(*mocks.MockJobRepository, *mocks.MockSkillRepository, *mocks.MockCategoryRepository, *mocks.MockLocationRepository)
		expectedError string
	}{
		{
			name:       "successful_job_creation",
			jobRequest: createValidJobRequest(),
			setupMocks: func(jobRepo *mocks.MockJobRepository, skillRepo *mocks.MockSkillRepository, categoryRepo *mocks.MockCategoryRepository, locationRepo *mocks.MockLocationRepository) {
				// No duplicate found
				jobRepo.On("IsDuplicateJob", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(false, nil)

				// Successful job creation
				createdJob := &model.Job{ID: 1, Title: "Software Engineer"}
				jobRepo.On("Create", mock.AnythingOfType("*model.Job")).Return(createdJob, nil)

				// Skills handling
				goSkill := &model.Skill{ID: 1, Name: "Go"}
				pythonSkill := &model.Skill{ID: 2, Name: "Python"}
				skillRepo.On("GetByName", "Go").Return(goSkill, nil)
				skillRepo.On("GetByName", "Python").Return(pythonSkill, nil)
				jobRepo.On("CreateJobSkill", mock.AnythingOfType("*model.JobSkill")).Return(&model.JobSkill{}, nil).Twice()

				// Location handling
				country := &model.Country{ID: 1, Name: "US", ISO: "US"}
				locationRepo.On("GetCountryByISO", "US").Return(country, nil)
				locationRepo.On("CreateJobLocation", mock.AnythingOfType("*model.JobLocation")).Return(&model.JobLocation{}, nil)

				// Category handling
				category := &model.Category{ID: 1, Name: "Engineering"}
				categoryRepo.On("GetCategoryByName", "Engineering").Return(category, nil)
				categoryRepo.On("CreateJobCategory", mock.AnythingOfType("*model.JobCategory")).Return(&model.JobCategory{}, nil)
			},
			expectedError: "",
		},
		{
			name:       "duplicate_job_error",
			jobRequest: createValidJobRequest(),
			setupMocks: func(jobRepo *mocks.MockJobRepository, skillRepo *mocks.MockSkillRepository, categoryRepo *mocks.MockCategoryRepository, locationRepo *mocks.MockLocationRepository) {
				// Duplicate found
				jobRepo.On("IsDuplicateJob", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(true, nil)
			},
			expectedError: "duplicate job entry",
		},
		{
			name:       "invalid_job_data_error",
			jobRequest: createInvalidJobRequest(),
			setupMocks: func(jobRepo *mocks.MockJobRepository, skillRepo *mocks.MockSkillRepository, categoryRepo *mocks.MockCategoryRepository, locationRepo *mocks.MockLocationRepository) {
				// No mocks needed - validation happens before repo calls
			},
			expectedError: "invalid job data",
		},
		{
			name:       "job_creation_fails",
			jobRequest: createValidJobRequest(),
			setupMocks: func(jobRepo *mocks.MockJobRepository, skillRepo *mocks.MockSkillRepository, categoryRepo *mocks.MockCategoryRepository, locationRepo *mocks.MockLocationRepository) {
				jobRepo.On("IsDuplicateJob", mock.Anything, mock.Anything).Return(false, nil)
				jobRepo.On("Create", mock.AnythingOfType("*model.Job")).Return(nil, errors.New("database error"))
			},
			expectedError: "failed to create job",
		},
		{
			name:       "skill_not_found_creates_new_skill",
			jobRequest: createValidJobRequest(),
			setupMocks: func(jobRepo *mocks.MockJobRepository, skillRepo *mocks.MockSkillRepository, categoryRepo *mocks.MockCategoryRepository, locationRepo *mocks.MockLocationRepository) {
				jobRepo.On("IsDuplicateJob", mock.Anything, mock.Anything).Return(false, nil)
				createdJob := &model.Job{ID: 1}
				jobRepo.On("Create", mock.AnythingOfType("*model.Job")).Return(createdJob, nil)

				// First skill not found, second skill found
				skillRepo.On("GetByName", "Go").Return(nil, gorm.ErrRecordNotFound)
				skillRepo.On("Create", mock.AnythingOfType("*model.Skill")).Return(&model.Skill{ID: 1, Name: "Go"}, nil)
				pythonSkill := &model.Skill{ID: 2, Name: "Python"}
				skillRepo.On("GetByName", "Python").Return(pythonSkill, nil)
				jobRepo.On("CreateJobSkill", mock.AnythingOfType("*model.JobSkill")).Return(nil, errors.New("skill cretion error"))

				// Skip location and category for simplicity
				locationRepo.On("GetCountryByISO", mock.Anything).Return(&model.Country{ID: 1}, nil)
				locationRepo.On("CreateJobLocation", mock.Anything).Return(nil, errors.New("sdfd"))
				categoryRepo.On("GetCategoryByName", mock.Anything).Return(&model.Category{ID: 1}, nil)
				categoryRepo.On("CreateJobCategory", mock.Anything).Return(nil, errors.New("asdf"))
			},
			expectedError: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mocks
			mockJobRepo := &mocks.MockJobRepository{}
			mockSkillRepo := &mocks.MockSkillRepository{}
			mockCategoryRepo := &mocks.MockCategoryRepository{}
			mockLocationRepo := &mocks.MockLocationRepository{}

			tt.setupMocks(mockJobRepo, mockSkillRepo, mockCategoryRepo, mockLocationRepo)

			// Create service
			service := NewJobService(mockJobRepo, mockSkillRepo, mockCategoryRepo, mockLocationRepo)

			// Execute
			_, err := service.CreateJob(tt.jobRequest)

			// Assert
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}

			// Verify all expected calls were made
			mockJobRepo.AssertExpectations(t)
			mockSkillRepo.AssertExpectations(t)
			mockCategoryRepo.AssertExpectations(t)
			mockLocationRepo.AssertExpectations(t)
		})
	}
}

// Test GetJobByID function
func TestJobService_GetJobByID(t *testing.T) {
	tests := []struct {
		name          string
		jobID         uint
		mockReturn    *model.Job
		mockError     error
		expectedError bool
	}{
		{
			name:          "job_found_successfully",
			jobID:         1,
			mockReturn:    &model.Job{ID: 1, Title: "Software Engineer"},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "job_not_found",
			jobID:         999,
			mockReturn:    nil,
			mockError:     gorm.ErrRecordNotFound,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockJobRepo := &mocks.MockJobRepository{}
			mockJobRepo.On("GetByID", tt.jobID).Return(tt.mockReturn, tt.mockError)

			// Create service
			service := NewJobService(mockJobRepo, &mocks.MockSkillRepository{}, &mocks.MockCategoryRepository{}, &mocks.MockLocationRepository{})

			// Execute
			job, err := service.GetJobByID(tt.jobID)

			// Assert
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, job)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, job)
				assert.Equal(t, tt.jobID, job.ID)
			}

			mockJobRepo.AssertExpectations(t)
		})
	}
}

// Test DeleteJob function
func TestJobService_DeleteJob(t *testing.T) {
	tests := []struct {
		name          string
		jobID         uint
		setupMocks    func(*mocks.MockJobRepository)
		expectedError string
	}{
		{
			name:  "successful_deletion",
			jobID: 1,
			setupMocks: func(jobRepo *mocks.MockJobRepository) {
				jobRepo.On("GetByID", uint(1)).Return(&model.Job{ID: 1, Title: "Test Job"}, nil)
				jobRepo.On("Delete", uint(1)).Return(nil)
			},
			expectedError: "",
		},
		{
			name:  "job_not_found",
			jobID: 999,
			setupMocks: func(jobRepo *mocks.MockJobRepository) {
				jobRepo.On("GetByID", uint(999)).Return(nil, gorm.ErrRecordNotFound)
			},
			expectedError: "job not found",
		},
		{
			name:  "delete_operation_fails",
			jobID: 1,
			setupMocks: func(jobRepo *mocks.MockJobRepository) {
				jobRepo.On("GetByID", uint(1)).Return(&model.Job{ID: 1}, nil)
				jobRepo.On("Delete", uint(1)).Return(errors.New("database error"))
			},
			expectedError: "failed to delete job",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup mock
			mockJobRepo := &mocks.MockJobRepository{}
			tt.setupMocks(mockJobRepo)

			// Create service
			service := NewJobService(mockJobRepo, &mocks.MockSkillRepository{}, &mocks.MockCategoryRepository{}, &mocks.MockLocationRepository{})

			// Execute
			err := service.DeleteJob(tt.jobID)

			// Assert
			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}

			mockJobRepo.AssertExpectations(t)
		})
	}
}

// Test ConvertJobRequest function (validation logic)
func TestConvertJobRequest(t *testing.T) {
	tests := []struct {
		name        string
		jobRequest  dtos.JobRequest
		expectedErr string
	}{
		{
			name:        "valid_request",
			jobRequest:  createValidJobRequest(),
			expectedErr: "",
		},
		{
			name: "missing_title",
			jobRequest: dtos.JobRequest{
				CompanyName: "Tech Corp",
				JobType:     1,
				WorkMode:    1,
			},
			expectedErr: "title is required",
		},
		{
			name: "missing_company_name",
			jobRequest: dtos.JobRequest{
				Title:    "Software Engineer",
				JobType:  1,
				WorkMode: 1,
			},
			expectedErr: "company name is required",
		},
		{
			name: "missing_job_type",
			jobRequest: dtos.JobRequest{
				Title:       "Software Engineer",
				CompanyName: "Tech Corp",
				WorkMode:    1,
			},
			expectedErr: "job type is required",
		},
		{
			name: "missing_work_mode",
			jobRequest: dtos.JobRequest{
				Title:       "Software Engineer",
				CompanyName: "Tech Corp",
				JobType:     1,
			},
			expectedErr: "work mode is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute
			job, err := ConvertJobRequest(tt.jobRequest)

			// Assert
			if tt.expectedErr == "" {
				assert.NoError(t, err)
				assert.NotNil(t, job)
				assert.Equal(t, tt.jobRequest.Title, job.Title)
				assert.Equal(t, tt.jobRequest.CompanyName, job.CompanyName)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
				assert.Nil(t, job)
			}
		})
	}
}
