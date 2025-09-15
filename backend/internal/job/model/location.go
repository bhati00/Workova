package model

type Country struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
	ISO  string `gorm:"type:varchar(10);not null;unique" json:"iso"`
}

func (Country) TableName() string {
	return "countries"
}

type JobLocation struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID     uint    `gorm:"index;not null" json:"job_id"`
	Job       Job     `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job"`
	CountryID uint    `gorm:"index;not null" json:"country_id"`
	Country   Country `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE" json:"country"`
	City      *string `gorm:"type:varchar(100)" json:"city,omitempty"`
}

func (JobLocation) TableName() string {
	return "job_locations"
}
