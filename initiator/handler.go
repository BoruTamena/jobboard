package initiator

import (
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/auth"
)

type Handler struct {
	authHandler handler.Auth
}

func InitHandler(md Module) Handler {

	return Handler{
		authHandler: auth.AuthHandler(md.authModule),
	}
}
