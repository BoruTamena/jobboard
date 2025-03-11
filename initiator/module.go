package initiator

import (
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/module/auth"
	"github.com/BoruTamena/jobboard/internal/module/user"
	"github.com/BoruTamena/jobboard/platform"
)

type Module struct {
	authModule        module.AuthModule
	userProfileModule module.UserProfile

	cache platform.RedisCache
}

func InitModule(persistence Persistance, platform Platform) Module {

	return Module{
		authModule:        auth.NewAuthModule(persistence.authStorage, platform.redis),
		userProfileModule: user.InitUserProfileModule(persistence.profileStorage),
	}

}
