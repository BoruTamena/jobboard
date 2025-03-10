package initiator

import (
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
	"github.com/BoruTamena/jobboard/internal/storage/auth"
	"github.com/BoruTamena/jobboard/internal/storage/user"
)

type Persistance struct {
	authStorage    storage.AuthStorage
	profileStorage storage.UserProfie
}

func InitPersistence(db persistencedb.PersistenceDb) Persistance {

	return Persistance{

		authStorage:    auth.NewAuthStorage(db),
		profileStorage: user.NewUserProfileStorage(db),
	}

}
