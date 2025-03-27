package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

func (p Pagination) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Page, validation.Required, is.Int),
		validation.Field(&p.PerPage, validation.Required, is.Int),
	)
}
