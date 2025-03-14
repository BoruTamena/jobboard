package user

import (
	"context"

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

	profile := db.UserProfile{
		UserID:   user_id,
		Bio:      userProfile.Bio,
		Location: userProfile.Location,
	}

	res := userPStorage.db.WithContext(ctx).Create(&profile)

	if err := res.Error; err != nil {

		err = errors.WriteErr.Wrap(err, "error while createing profile").
			WithProperty(errors.ErrorCode, 500)

		return db.UserProfile{}, err

	}

	return profile, nil

}

func (userPStorage *userProfileStorage) UpdateUserProfile(ctx context.Context, user_id uuid.UUID, userProfile dto.UserProfie) (db.UserProfile, error) {

	var profile db.UserProfile

	res := userPStorage.db.WithContext(ctx).Model(&profile).
		Updates(map[string]interface{}{
			"bio":      userProfile.Bio,
			"location": userProfile.Location,
		})

	if err := res.Error; err != nil {

		err := errors.WriteErr.Wrap(err, "error occur while updating user profile").
			WithProperty(errors.ErrorCode, 500)

		return db.UserProfile{}, err

	}

	return profile, nil

}

func (userPStorage *userProfileStorage) GetUserProfile(ctx context.Context, user_id uuid.UUID) (db.User, error) {

	var user db.User
	res := userPStorage.db.WithContext(ctx).Preload("Profile").
		Where(map[string]interface{}{"id": user_id}).
		Find(&user)

	if err := res.Error; err != nil {

		err := errors.DbReadErr.Wrap(err, "error while reading user").
			WithProperty(errors.ErrorCode, 500)

		return db.User{}, err

	}

	return user, nil

}
