package initiator

import (
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/module/auth"
	"github.com/BoruTamena/jobboard/internal/module/user"
)

type Module struct {
	authModule module.AuthModule

	userProfileModule module.UserProfile
}

func InitModule(persistence Persistance) Module {

	return Module{
		authModule:        auth.NewAuthModule(persistence.authStorage),
		userProfileModule: user.InitUserProfileModule(persistence.profileStorage),
	}

}
