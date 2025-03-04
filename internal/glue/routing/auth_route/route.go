package authroute

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/glue/routing"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(rg *gin.RouterGroup, handler handler.Auth) {

	routes := []routing.Router{
		{

			Method: http.MethodPost,
			Path:   "user",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
			},
			Handler: handler.RegisterUser,
		},

		{

			Method: http.MethodPost,
			Path:   "signin",
			Middlewares: []gin.HandlerFunc{
				middleware.AuthMiddleware(),
				middleware.ErrorMiddleWare(),
			},
			Handler: handler.SignIn,
		},
	}

	routing.RegisterRoute(rg, routes)

}
