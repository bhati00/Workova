package model

import (
	"time"

	"github.com/bhati00/workova/backend/constant"
	"gorm.io/gorm"
)

type Job struct {
	// Primary identifiers
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Slug          *string `gorm:"size:150;index" json:"slug"`            // URL-friendly identifier
	ExternalJobID *string `gorm:"size:255;index" json:"external_job_id"` // ID from external source

	// Core job information
	Title       string  `gorm:"size:255;not null;index:idx_title" json:"title"`
	Description *string `gorm:"type:longtext" json:"description"` // Changed from 'text' to 'longtext'
	Summary     *string `gorm:"type:text" json:"summary"`         // Short description/excerpt

	// Company information. NOTE : 	i need to separate company as a different table
	CompanyName     string  `gorm:"size:255;not null;index:idx_company" json:"company_name"`
	CompanySize     *string `gorm:"size:50" json:"company_size"` // "1-10", "11-50", "51-200", etc.
	CompanyLogo     *string `gorm:"type:text" json:"company_logo_url"`
	CompanyWebsite  *string `gorm:"type:text" json:"company_website"`
	CompanyIndustry *string `gorm:"size:100" json:"company_industry"`

	// Employment details
	JobType         constant.JobType          `gorm:"type:int;not null;default:1;index:idx_job_type" json:"job_type"`   // Default: FullTime
	WorkMode        constant.WorkMode         `gorm:"type:int;not null;default:1;index:idx_work_mode" json:"work_mode"` // Default: Remote
	ExperienceLevel *constant.ExperienceLevel `gorm:"type:int;index:idx_experience" json:"experience_level"`            // Nullable

	// Location information (normalized)
	IsRemote *bool `gorm:"index:idx_remote;default:false" json:"is_remote"`

	// Compensation details
	SalaryMin        *int              `json:"salary_min"`
	SalaryMax        *int              `json:"salary_max"`
	SalaryCurrency   constant.Currency `gorm:"size:10;default:'USD'" json:"salary_currency"` // Default: USD, but nullable
	SalaryPeriod     *string           `gorm:"size:20" json:"salary_period"`                 // "hourly", "monthly", "yearly"
	SalaryIsEstimate *bool             `gorm:"default:false" json:"salary_is_estimate"`
	EquityOffered    *bool             `json:"equity_offered"`

	// Experience requirements
	YearsExperienceMin *int `json:"years_experience_min"`
	YearsExperienceMax *int `json:"years_experience_max"`

	// Benefits and perks
	Benefits         *string `gorm:"type:text" json:"benefits"` // JSON array or comma-separated
	HealthInsurance  *bool   `json:"health_insurance"`
	PaidTimeOff      *bool   `json:"paid_time_off"`
	FlexibleSchedule *bool   `json:"flexible_schedule"`

	// Application and contact details
	ApplicationURL   *string `gorm:"type:text" json:"application_url"` // Made nullable since not all jobs might have this
	ApplicationEmail *string `gorm:"size:255" json:"application_email"`
	ContactPerson    *string `gorm:"size:255" json:"contact_person"`
	HowToApply       *string `gorm:"type:text" json:"how_to_apply"` // Application instructions

	// Requirements
	EducationLevel    *string `gorm:"size:100" json:"education_level"`    // "high_school", "bachelor", "master", "phd"
	EducationRequired *string `gorm:"size:255" json:"education_required"` // Specific education requirements

	// Work arrangement details
	ContractMonthsDuration    *int       `gorm:"type:int" json:"contract_duration"` // "6 months", "1 year", "permanent"
	StartDate                 *time.Time `json:"start_date"`
	IsUrgent                  *bool      `gorm:"default:false;index:idx_urgent" json:"is_urgent"`
	TravelRequired            *string    `gorm:"size:50" json:"travel_required"`              // "none", "occasional", "frequent"
	RemoteLocationRestriction *string    `gorm:"size:255" json:"remote_location_restriction"` // Geographic restrictions for remote work

	// Interview and hiring process
	InterviewProcess   *string    `gorm:"size:50" json:"interview_process"` // "phone", "video", "in_person", "mixed"
	HiringProcessSteps *int       `json:"hiring_process_steps"`             // Number of interview rounds
	ExpectedHireDate   *time.Time `json:"expected_hire_date"`

	// Visa and legal
	VisaSponsorship   *bool   `json:"visa_sponsorship"`
	SecurityClearance *string `gorm:"size:50" json:"security_clearance"` // "none", "confidential", "secret", "top_secret"
	BackgroundCheck   *bool   `json:"background_check_required"`

	// Job posting metadata
	Source              string     `gorm:"size:100;not null;uniqueIndex:idx_job_source" json:"source"` // "indeed", "linkedin", "arbeitnow"
	SourceURL           *string    `gorm:"type:text" json:"source_url"`                                // Original job posting URL
	PostedDate          *time.Time `gorm:"index:idx_posted" json:"posted_date"`                        // When originally posted by employer
	ExpiryDate          *time.Time `gorm:"index:idx_expiry" json:"expiry_date"`                        // When job posting expires
	ApplicationDeadline *time.Time `json:"application_deadline"`

	// SEO and categorization
	Keywords   *string `gorm:"type:text" json:"keywords"`                       // SEO keywords, comma-separated
	Industry   *string `gorm:"size:100;index:idx_industry" json:"industry"`     // "Technology", "Healthcare", etc.
	Department *string `gorm:"size:100;index:idx_department" json:"department"` // "Engineering", "Marketing", etc.
	Tags       *string `gorm:"type:text" json:"tags"`                           // JSON array of tags

	CreatedAt time.Time      `gorm:"index:idx_created" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Foreign key relationships
	JobSkills     []JobSkill    `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job_skills,omitempty"`
	JobCategories []JobCategory `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job_categories,omitempty"`
	JobLocations  []JobLocation `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job_locations,omitempty"`
}

// TableName specifies the table name for the Job model
func (Job) TableName() string {
	return "jobs"
}

// BeforeCreate hook to set default values
func (j *Job) BeforeCreate(tx *gorm.DB) error {
	// Set default job type if not provided
	if j.JobType == 0 {
		j.JobType = constant.JobTypeFullTime
	}

	// Set default work mode if not provided
	if j.WorkMode == 0 {
		j.WorkMode = constant.WorkModeRemote
	}

	// Set default values for boolean fields if not explicitly set
	if j.IsRemote == nil {
		defaultRemote := false
		j.IsRemote = &defaultRemote
	}

	if j.SalaryIsEstimate == nil {
		defaultEstimate := false
		j.SalaryIsEstimate = &defaultEstimate
	}

	if j.IsUrgent == nil {
		defaultUrgent := false
		j.IsUrgent = &defaultUrgent
	}
	return nil
}
