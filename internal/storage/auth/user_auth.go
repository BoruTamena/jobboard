package auth

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
	"github.com/google/uuid"
)

type authStorage struct {
	db persistencedb.PersistenceDb
}

func NewAuthStorage(db persistencedb.PersistenceDb) storage.AuthStorage {

	return &authStorage{
		db: db,
	}

}

func (ats *authStorage) CreateUser(ctx context.Context, user dto.UserDto) (error, db.User) {

	user_db, err := ats.db.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		Role:     "job_seeker",
	})

	if err != nil {

		err = errors.WriteErr.Wrap(err, "db write error").
			WithProperty(errors.ErrorCode, 500)

		log.Println("db write error while creating user", err.Error())

		return err, db.User{}

	}

	return nil, user_db

}

func (ats *authStorage) GetUserByEmail(ctx context.Context, userlg dto.UserLogin) (error, db.User) {

	user, err := ats.db.GetUserByEmail(ctx, userlg.Email)

	if err != nil {

		err = errors.DbReadErr.Wrap(err, "error while getting user by email").
			WithProperty(errors.ErrorCode, 500)

		log.Println("ERROR occur while fetching user by email", err.Error())

		return err, db.User{}
	}

	return nil, user
}
