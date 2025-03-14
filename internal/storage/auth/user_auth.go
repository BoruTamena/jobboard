package auth

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
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

	db_user := db.User{UserName: user.UserName, Email: user.Email, Password: user.Password}
	res := ats.db.WithContext(ctx).Create(&db_user)
	if err := res.Error; err != nil {

		err = errors.WriteErr.Wrap(err, "error while creating user").
			WithProperty(errors.ErrorCode, 500)

		log.Println("ERROR occur while creating user", err.Error())

		return err, db.User{}

	}

	return nil, db_user

}

func (ats *authStorage) GetUserByEmail(ctx context.Context, userlg dto.UserLogin) (error, db.User) {

	var user db.User
	res := ats.db.WithContext(ctx).Where(map[string]interface{}{
		"email": userlg.Email,
	}).Find(&user)

	if err := res.Error; err != nil {

		err = errors.DbReadErr.Wrap(err, "error while fetching user by email").
			WithProperty(errors.ErrorCode, 500)

		return err, db.User{}

	}

	return nil, user

}
