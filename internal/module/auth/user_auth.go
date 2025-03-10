package auth

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/helper"
	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/storage"
)

type authModule struct {
	authStorage storage.AuthStorage
}

func NewAuthModule(astorage storage.AuthStorage) module.AuthModule {

	return &authModule{
		authStorage: astorage,
	}

}

func (atmd *authModule) RegisterUser(ctx context.Context, user dto.UserDto) (error, dto.UserDto) {

	if err := user.Validate(); err != nil {

		// validation error
		err = errors.BadInput.Wrap(err, "").
			WithProperty(errors.ErrorCode, 400)

		log.Println("BAD USER INTO :- %v", err.Error())
		return err, dto.UserDto{}

	}

	h_pass, _ := helper.HashPassword(user.Password)

	user.Password = h_pass

	err, db_res := atmd.authStorage.CreateUser(ctx, user)

	if err != nil {
		return err, dto.UserDto{}
	}

	return nil, dto.UserDto{
		UserName: db_res.UserName,
		Password: user.Password,
		Email:    user.Email,
	}
}

func (atmd *authModule) SignIn(ctx context.Context, userlg dto.UserLogin) (error, db.User) {

	if err := userlg.Validate(); err != nil {

		err = errors.BadInput.Wrap(err, "bad user input").
			WithProperty(errors.ErrorCode, 500)

		log.Println("bad user login input :-", err.Error())

		return err, db.User{}
	}

	err, user := atmd.authStorage.GetUserByEmail(ctx, userlg)

	//checking credential
	if !helper.VerifyPassword(userlg.Password, user.Password) {

		err = errors.AuthErr.New("Invalid username or password").
			WithProperty(errors.ErrorCode, 401)

		log.Println("invalid credential while login :-", err.Error())

		return err, db.User{}

	}
	return nil, user
}
