package mocks

import (
	"github.com/bhati00/workova/backend/internal/job/model"
	"github.com/stretchr/testify/mock"
)

type MockLocationRepository struct {
	mock.Mock
}

func (m *MockLocationRepository) GetCountryByISO(iso string) (*model.Country, error) {
	args := m.Called(iso)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Country), args.Error(1)
}

func (m *MockLocationRepository) CreateCountry(country *model.Country) (*model.Country, error) {
	args := m.Called(country)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Country), args.Error(1)
}

func (m *MockLocationRepository) CreateJobLocation(jobLocation *model.JobLocation) (*model.JobLocation, error) {
	args := m.Called(jobLocation)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.JobLocation), args.Error(1)
}
