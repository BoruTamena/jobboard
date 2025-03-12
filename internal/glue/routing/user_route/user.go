package userroute

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/glue/routing"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/middleware"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/gin-gonic/gin"
)

func InitUserProfileRouter(rg *gin.RouterGroup, handler handler.UserProfile, md module.AuthzModule) {

	routes := []routing.Router{
		{
			Path:   "/users/:id/profile",
			Method: http.MethodPost,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
				middleware.Authorize(md),
			},
			Handler: handler.CreateUserProfile,
		},

		{
			Path:   "/users/:id/profile",
			Method: http.MethodPut,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
				middleware.Authorize(md),
			},
			Handler: handler.UpdateUserProfile,
		},

		{
			Path:   "/users/:id/profile",
			Method: http.MethodGet,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
				middleware.Authorize(md),
			},
			Handler: handler.GetUserProfile,
		},
	}

	routing.RegisterRoute(rg, routes)

}
