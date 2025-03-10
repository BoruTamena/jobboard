package initiator

import (
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/auth"
	"github.com/BoruTamena/jobboard/internal/handler/user"
)

type Handler struct {
	authHandler        handler.Auth
	userProfileHandler handler.UserProfile
}

func InitHandler(md Module) Handler {

	return Handler{
		authHandler:        auth.AuthHandler(md.authModule),
		userProfileHandler: user.InitUserProfileHandler(md.userProfileModule),
	}
}
