package initiator

import (
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
	"github.com/BoruTamena/jobboard/internal/storage/application"
	"github.com/BoruTamena/jobboard/internal/storage/auth"
	"github.com/BoruTamena/jobboard/internal/storage/authz"
	"github.com/BoruTamena/jobboard/internal/storage/job"
	"github.com/BoruTamena/jobboard/internal/storage/user"
)

type Persistance struct {
	authStorage           storage.AuthStorage
	authzStorage          storage.AuthzStorage
	profileStorage        storage.UserProfie
	jobStorage            storage.JobStorage
	jobApplicationStorage storage.JobApplicationStorage
}

func InitPersistence(cfg dto.Config, db persistencedb.PersistenceDb) Persistance {

	return Persistance{

		authStorage:           auth.NewAuthStorage(db),
		authzStorage:          authz.InitAuthzStorage(cfg),
		profileStorage:        user.NewUserProfileStorage(db),
		jobStorage:            job.NewJobStorage(db),
		jobApplicationStorage: application.NewJobApplicationStorage(db),
	}

}
