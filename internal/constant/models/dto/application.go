package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
)

type AppilicationDto struct {
	// define your application struct here
	UserID      uuid.UUID
	JobID       uuid.UUID
	ResumeURL   string `json:"resume_url"`
	CoverLetter string `json:"cover_letter"`
	Status      string `json:"status"` // pending  accepted rejected
}

func (j AppilicationDto) Validate() error {
	return validation.ValidateStruct(&j,
		validation.Field(&j.UserID, validation.Required, is.UUIDv4),
		validation.Field(&j.JobID, validation.Required, is.UUIDv4),
		validation.Field(&j.ResumeURL, validation.Required),
		validation.Field(&j.CoverLetter, validation.Required),
		validation.Field(&j.Status, validation.Required, validation.In("pending", "accepted", "rejected")),
	)
}
