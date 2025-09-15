package mocks

import (
	"github.com/bhati00/workova/backend/internal/job/model"
	"github.com/stretchr/testify/mock"
)

type MockSkillRepository struct {
	mock.Mock
}

func (m *MockSkillRepository) GetByName(name string) (*model.Skill, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Skill), args.Error(1)
}

func (m *MockSkillRepository) Create(skill *model.Skill) (*model.Skill, error) {
	args := m.Called(skill)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Skill), args.Error(1)
}
func (m *MockSkillRepository) GetAll() ([]model.Skill, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Skill), args.Error(1)
}
