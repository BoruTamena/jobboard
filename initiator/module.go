package initiator

import (
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/module/auth"
)

type Module struct {

	/* all your modules goes here

	user  module.User

	*/

	authModule module.AuthModule
}

func InitModule(persistence Persistance) Module {

	return Module{
		authModule: auth.NewAuthModule(persistence.authStorage),
	}

}
