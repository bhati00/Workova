package model

type Skill struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"`
}

func (Skill) TableName() string {
	return "skills"
}

type JobSkill struct {
	ID      uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	JobID   uint  `gorm:"index;not null" json:"job_id"`
	Job     Job   `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE" json:"job"`
	Skill   Skill `gorm:"foreignKey:SkillID;constraint:OnDelete:CASCADE" json:"skill"`
	SkillID uint  `gorm:"index;not null" json:"skill_id"`
}

func (JobSkill) TableName() string {
	return "job_skills"
}
