package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type JobDto struct {
	JobTitle       string    `json:"job_title"`
	JobDescription string    `json:"job_description"`
	JobType        string    `json:"job_type"`
	Location       string    `json:"location"`
	Status         string    `json:"status"`
	CategoryId     uuid.UUID `json:"category_id"`
}

func (j JobDto) Validate() error {
	return validation.ValidateStruct(&j,
		validation.Field(&j.JobTitle, validation.Required),
		validation.Field(&j.JobDescription, validation.Required),
		validation.Field(&j.JobType, validation.Required),
		validation.Field(&j.Location, validation.Required),
		validation.Field(&j.Status, validation.Required),
		validation.Field(&j.CategoryId, validation.Required, is.UUIDv4),
	)
}

// jobcategory
type JobCategoryDto struct {
	ID           uuid.UUID `json:"id"`
	CategoryName string    `json:"name"`
	Description  string    `json:"description"`
}
