package initiator

import (
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
	"github.com/BoruTamena/jobboard/internal/storage/auth"
)

type Persistance struct {
	authStorage storage.AuthStorage
}

func InitPersistence(db persistencedb.PersistenceDb) Persistance {

	return Persistance{

		authStorage: auth.NewAuthStorage(db),
	}

}
