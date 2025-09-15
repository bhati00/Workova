// internal/job/mocks/job_repository_mock.go
package mocks

import (
	"github.com/bhati00/workova/backend/dtos"
	"github.com/bhati00/workova/backend/internal/job/model"
	"github.com/stretchr/testify/mock"
)

type MockJobRepository struct {
	mock.Mock
}

func (m *MockJobRepository) IsDuplicateJob(externalID *string, slug *string) (bool, error) {
	args := m.Called(externalID, slug)
	return args.Bool(0), args.Error(1)
}

func (m *MockJobRepository) Create(job *model.Job) (*model.Job, error) {
	args := m.Called(job)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJobRepository) GetByID(id uint) (*model.Job, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Job), args.Error(1)
}

func (m *MockJobRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockJobRepository) Update(job *model.Job) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockJobRepository) BatchDelete(ids []uint) (*dtos.BatchResult, error) {
	args := m.Called(ids)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.BatchResult), args.Error(1)
}

func (m *MockJobRepository) GetAll(offset, limit int) ([]model.Job, error) {
	args := m.Called(offset, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Job), args.Error(1)
}

func (m *MockJobRepository) SearchJobs(params *dtos.JobSearchParams) ([]model.Job, int64, error) {
	args := m.Called(params)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]model.Job), args.Get(1).(int64), args.Error(2)
}

func (m *MockJobRepository) CountActiveJobs() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockJobRepository) CreateJobSkill(jobSkill *model.JobSkill) (*model.JobSkill, error) {
	args := m.Called(jobSkill)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.JobSkill), args.Error(1)
}
func (m *MockJobRepository) CreateJobCategory(jobCategory *model.JobCategory) (*model.JobCategory, error) {
	args := m.Called(jobCategory)
	return args.Get(0).(*model.JobCategory), args.Error(1)
}
func (m *MockJobRepository) CreateJobLocation(jobLocation *model.JobLocation) (*model.JobLocation, error) {
	args := m.Called(jobLocation)
	return args.Get(0).(*model.JobLocation), args.Error(0)
}
func (m *MockJobRepository) SoftDelete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
