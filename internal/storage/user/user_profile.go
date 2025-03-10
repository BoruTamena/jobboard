package user

import (
	"context"
	"database/sql"

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
		},
		Location: sql.NullString{
			String: userProfile.Location,
		},
	})

	if err != nil {
		err = errors.WriteErr.Wrap(err, "db write op error").
			WithProperty(errors.ErrorCode, 500)

		return db.UserProfile{}, err
	}
	return profile, nil
}
