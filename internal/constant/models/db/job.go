package db

import "github.com/google/uuid"

type Job struct {
	BaseModel
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	JobType        string           `json:"job_type"` // full-time, part-time, contract, internship
	Location       string           `json:"location"`
	Status         string           ` json:"status"`
	CategoryId     uuid.UUID        `json:"category_id"`
	JobCategory    JobCategory      `gorm:"foreignKey:CategoryId"`
	JobApplication []JobApplication `gorm:"foreignKey:JobId"`
}

// job category
type JobCategory struct {
	BaseModel
	Name        string `json:"name"` // IT & Software, Marketing, Sales, etc
	Description string `json:"description"`
	Job         []Job  `gorm:"foreignKey:CategoryId"`
}

// job application
type JobApplication struct {
	BaseModel
	UserId      uuid.UUID
	JobId       uuid.UUID
	Resume      string `json:"resume_url"`
	CoverLetter string `json:"cover_letter"`
	Status      string `json:"status"` // pending  accepted rejected
	User        User   `gorm:"foreignKey:UserId"`
	Job         Job    `gorm:"foreignKey:JobId"`
}
