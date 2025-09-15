package repository

import (
	"github.com/bhati00/workova/backend/internal/job/model"
	"gorm.io/gorm"
)

type SkillRepository interface {
	Create(skill *model.Skill) (*model.Skill, error)
	GetAll() ([]model.Skill, error)
	GetByName(name string) (*model.Skill, error)
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) skillRepository {
	return skillRepository{db: db}
}

func (r skillRepository) Create(skill *model.Skill) (*model.Skill, error) {
	if err := r.db.Create(&skill).Error; err != nil {
		if gorm.ErrDuplicatedKey.Error() == err.Error() {
			// If duplicate key error, fetch existing category
			var existing *model.Skill
			if err := r.db.Where("name = ?", skill.Name).First(&existing).Error; err != nil {
				return &model.Skill{}, err
			}
			return existing, nil
		}
		return &model.Skill{}, err
	}
	return skill, nil
}

func (r skillRepository) GetAll() ([]model.Skill, error) {
	var skills []model.Skill
	if err := r.db.Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}
func (r skillRepository) GetByName(name string) (*model.Skill, error) {
	var skill model.Skill
	if err := r.db.Where("name = ?", name).First(&skill).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}
