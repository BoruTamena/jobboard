package user

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/storage"
	"github.com/google/uuid"
)

type userProfile struct {
	userStorageProfile storage.UserProfie
}

func InitUserProfileModule(userPrStorage storage.UserProfie) module.UserProfile {

	return &userProfile{
		userStorageProfile: userPrStorage,
	}

}

func (upm *userProfile) CreateUserProfile(ctx context.Context, user_id string, userProfileData dto.UserProfie) (db.UserProfile, error) {

	// TODO CONVERT USERID TO APROPRIATE TYPE

	Id, err := uuid.Parse(user_id)

	if err != nil {

		return db.UserProfile{}, err
	}

	if err := userProfileData.Validate(); err != nil {
		// validation
		err := errors.BadInput.Wrap(err, "bad user input ").
			WithProperty(errors.ErrorCode, 400)

		log.Println("bad user input for create profile", err.Error())

		return db.UserProfile{}, err

	}

	// calling storage

	profile, err := upm.userStorageProfile.CreateUserProfile(ctx, Id, userProfileData)

	if err != nil {
		return db.UserProfile{}, err
	}

	return profile, nil

}

func (upm *userProfile) UpdateUserProfile(ctx context.Context, user_id string, userProfileData dto.UserProfie) (db.UserProfile, error) {

	// TODO CONVERT USERID TO APROPRIATE TYPE

	Id, err := uuid.Parse(user_id)

	if err != nil {

		return db.UserProfile{}, err
	}

	if err := userProfileData.Validate(); err != nil {
		// validation
		err := errors.BadInput.Wrap(err, "bad user input ").
			WithProperty(errors.ErrorCode, 400)

		log.Println("bad user input for create profile", err.Error())

		return db.UserProfile{}, err

	}

	// calling storage

	profile, err := upm.userStorageProfile.UpdateUserProfile(ctx, Id, userProfileData)

	if err != nil {
		return db.UserProfile{}, err
	}

	return profile, nil

}

func (upm *userProfile) GetUserProfile(ctx context.Context, user_id string) (db.User, error) {

	Id, err := uuid.Parse(user_id)

	if err != nil {

		err = errors.BadInput.Wrap(err, "invalid user_id").
			WithProperty(errors.ErrorCode, 500)

		log.Println("Invalid user id")
		return db.User{}, err
	}

	// calling module
	profile, err := upm.userStorageProfile.GetUserProfile(ctx, Id)

	if err != nil {

		return db.User{}, err
	}

	return profile, nil

}
