package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserDto struct {
	UserName string `json:"user_name" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

// validation
func (u UserDto) Validate() error {

	return validation.ValidateStruct(&u,
		validation.Field(&u.UserName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 8)),
	)
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// login validation
func (l UserLogin) Validate() error {

	return validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(6, 8)),
	)

}

type UserProfie struct {
	Bio      string   `json:"bio"`
	Location string   `json:"location"`
	Skills   []string `json:"skills"`
}

// profile validation
func (up UserProfie) Validate() error {

	return validation.ValidateStruct(&up,
		validation.Field(&up.Bio, validation.Required),
		validation.Field(&up.Location, validation.Required),
	)

}
