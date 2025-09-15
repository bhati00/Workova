package model

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
}

func (Category) TableName() string {
	return "categories"
}

type JobCategory struct {
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID      uint     `gorm:"index;not null" json:"job_id"`
	Job        Job      `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job"`
	CategoryID uint     `gorm:"index;not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"category"`
}

// TableName specifies the table name for the JobCategory model
func (JobCategory) TableName() string {
	return "job_categories"
}
