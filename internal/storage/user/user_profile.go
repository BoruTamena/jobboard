package user

import (
	"context"
	"database/sql"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"

	// "github.com/gofrs/uuid"
	// "github.com/gofrs/uuid"
	"github.com/google/uuid"
)

type userProfileStorage struct {
	db persistencedb.PersistenceDb
}

func NewUserProfileStorage(persistencedb persistencedb.PersistenceDb) storage.UserProfie {

	return &userProfileStorage{
		persistencedb,
	}

}

func (userPStorage *userProfileStorage) CreateUserProfile(ctx context.Context, user_id uuid.UUID, userProfile dto.UserProfie) (db.UserProfile, error) {

	profile, err := userPStorage.db.CreateProfile(ctx, db.CreateProfileParams{
		UserID: user_id,
		Bio: sql.NullString{
			String: userProfile.Bio,
			Valid:  true,
		},
		Location: sql.NullString{
			String: userProfile.Location,
			Valid:  true,
		},
	})

	if err != nil {
		err = errors.WriteErr.Wrap(err, "db write op error").
			WithProperty(errors.ErrorCode, 500)

		return db.UserProfile{}, err
	}
	return profile, nil
}

func (userPStorage *userProfileStorage) UpdateUserProfile(ctx context.Context, user_id uuid.UUID, userProfile dto.UserProfie) (db.UserProfile, error) {

	profile, err := userPStorage.db.UpdateProfile(ctx, db.UpdateProfileParams{
		UserID: user_id,
		Bio: sql.NullString{
			String: userProfile.Bio,
			Valid:  true,
		},
		Location: sql.NullString{
			String: userProfile.Location,
			Valid:  true,
		},
	})

	if err != nil {
		err = errors.WriteErr.Wrap(err, "db write op error").
			WithProperty(errors.ErrorCode, 500)

		return db.UserProfile{}, err
	}
	return profile, nil
}

func (userPStorage *userProfileStorage) GetUserProfile(ctx context.Context, user_id uuid.UUID) (db.GetUserProfileRow, error) {

	profile, err := userPStorage.db.GetUserProfile(ctx, user_id)

	if err != nil {
		err = errors.DbReadErr.Wrap(err, "fetching user profile").
			WithProperty(errors.ErrorCode, 500)

		log.Println("error occur while fetching user profile")
		return db.GetUserProfileRow{}, err
	}

	return profile, nil

}
