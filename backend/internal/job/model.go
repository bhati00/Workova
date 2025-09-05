package job

import (
	"errors"
	"time"

	"github.com/bhati00/workova/backend/constants"
	"gorm.io/gorm"
)

// Job represents a job posting
// @Description Job posting information
type Job struct {
	ID                        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID                     string    `gorm:"size:100;index" json:"job_id"`
	JobTitle                  string    `gorm:"size:255;not null" json:"job_title"`
	Description               string    `gorm:"type:text" json:"description"`
	ExperienceMin             *int      `json:"experience_min"`
	ExperienceMax             *int      `json:"experience_max"`
	CompensationMin           *int      `json:"compensation_min"`
	CompensationMax           *int      `json:"compensation_max"`
	Currency                  string    `gorm:"size:10" json:"currency"`
	VisaRequired              *bool     `json:"visa_required"`
	WorkMode                  string    `gorm:"size:50" json:"work_mode"`
	WorkType                  string    `gorm:"size:50" json:"work_type"`
	Preference                string    `gorm:"size:255" json:"preference"`
	NoOfRounds                *int      `json:"no_of_rounds"`
	InterviewMode             string    `gorm:"size:50" json:"interview_mode"`
	BondPeriod                string    `gorm:"size:100" json:"bond_period"`
	ShiftTimings              string    `gorm:"size:100" json:"shift_timings"`
	OvertimeApplicable        *bool     `json:"overtime_applicable"`
	Bonuses                   string    `gorm:"type:text" json:"bonuses"`
	PaySchedule               string    `gorm:"size:50" json:"pay_schedule"`
	LastUpdated               time.Time `gorm:"autoUpdateTime" json:"last_updated"`
	IsActive                  bool      `gorm:"default:true" json:"is_active"`
	ApplicationURL            string    `gorm:"type:text" json:"application_url"`
	EducationRequired         string    `gorm:"size:255" json:"education_required"`
	LanguagesRequired         string    `gorm:"type:text" json:"languages_required"`
	ContractDuration          string    `gorm:"size:100" json:"contract_duration"`
	RemoteLocationRestriction string    `gorm:"size:255" json:"remote_location_restriction"`
	Source                    string    `gorm:"size:100" json:"source"`

	// CompanyID uint    `json:"company_id"`
	// Company   Company `gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE" json:"company"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// @swagger:ignore
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	JobSkills []JobSkill `gorm:"foreignKey:JobID" json:"job_skills"`
}

// JobSkill represents a skill required for a job
// @Description Job skill information
type JobSkill struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID     uint           `gorm:"index;not null" json:"job_id"`
	Skill     string         `gorm:"size:100;not null" json:"skill"`
	Type      string         `gorm:"size:20" json:"type"` // Required, Good-to-have
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Job Job `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"-"`
}

func (j *Job) Validate() error {
	if !constants.IsValidWorkMode(j.WorkMode) {
		return errors.New("invalid work mode")
	}

	if !constants.IsValidWorkType(j.WorkType) {
		return errors.New("invalid work type")
	}

	if !constants.IsValidCurrency(j.Currency) {
		return errors.New("invalid currency")
	}

	return nil
}
