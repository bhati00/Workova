package repository

import (
	"github.com/bhati00/workova/backend/internal/job/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	CreateJobCategory(jobCategory *model.JobCategory) (*model.JobCategory, error)
	GetCategoryByName(name string) (*model.Category, error)
	GetAll() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) categoryRepository {
	return categoryRepository{db: db}
}

func (r categoryRepository) Create(category *model.Category) (*model.Category, error) {

	if err := r.db.Create(&category).Error; err != nil {
		if gorm.ErrDuplicatedKey.Error() == err.Error() {
			// If duplicate key error, fetch existing category
			var existing model.Category
			if err := r.db.Where("name = ?", category.Name).First(&existing).Error; err != nil {
				return &model.Category{}, err
			}
			return &existing, nil
		}
		return nil, err
	}
	return category, nil
}
func (r categoryRepository) CreateJobCategory(jobCategory *model.JobCategory) (*model.JobCategory, error) {
	if err := r.db.Create(&jobCategory).Error; err != nil {
		return &model.JobCategory{}, err
	}
	return jobCategory, nil
}
func (r categoryRepository) GetCategoryByName(name string) (*model.Category, error) {
	var category model.Category
	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (r categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
