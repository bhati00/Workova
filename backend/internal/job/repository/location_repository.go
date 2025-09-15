package repository

import (
	"github.com/bhati00/workova/backend/internal/job/model"
	"gorm.io/gorm"
)

type LocationRepository interface {
	// CreateJobLocation creates a new job location record
	CreateCountry(country *model.Country) (*model.Country, error)
	CreateJobLocation(jobLocation *model.JobLocation) (*model.JobLocation, error)
	GetCountryByISO(iso string) (*model.Country, error)
}

type locationRepostiory struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *locationRepostiory {
	repo := locationRepostiory{db: db}
	return &repo
}

func (r locationRepostiory) CreateCountry(country *model.Country) (*model.Country, error) {
	if err := r.db.Create(&country).Error; err != nil {
		if gorm.ErrDuplicatedKey.Error() == err.Error() {
			// If duplicate key error, fetch existing country
			var existing model.Country
			if err := r.db.Where("iso = ?", country.ISO).First(&existing).Error; err != nil {
				return &model.Country{}, err
			}
			return &existing, nil
		}
		return &model.Country{}, err
	}
	return country, nil
}
func (r locationRepostiory) CreateJobLocation(jobLocation *model.JobLocation) (*model.JobLocation, error) {
	if err := r.db.Create(&jobLocation).Error; err != nil {
		return &model.JobLocation{}, err
	}
	return jobLocation, nil
}
func (r locationRepostiory) GetCountryByISO(iso string) (*model.Country, error) {
	var country model.Country
	if err := r.db.Where("iso = ?", iso).First(&country).Error; err != nil {
		return nil, err
	}
	return &country, nil
}
