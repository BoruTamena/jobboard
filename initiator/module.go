package initiator

import (

	// "github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/module/auth"
	"github.com/BoruTamena/jobboard/internal/module/authz"
	"github.com/BoruTamena/jobboard/internal/module/job"
	"github.com/BoruTamena/jobboard/internal/module/user"
	"github.com/BoruTamena/jobboard/platform"
)

type Module struct {
	authModule        module.AuthModule
	authzModule       module.AuthzModule
	userProfileModule module.UserProfile
	jobModule         module.Job
	cache             platform.RedisCache
}

func InitModule(persistence Persistance, platform Platform) Module {

	return Module{
		authModule:        auth.NewAuthModule(persistence.authStorage, platform.redis),
		authzModule:       authz.InitAuthzModule(persistence.authzStorage),
		userProfileModule: user.InitUserProfileModule(persistence.profileStorage),
		jobModule:         job.NewJobModule(persistence.jobStorage),
		cache:             platform.redis,
	}

}
