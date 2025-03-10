package userroute

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/glue/routing"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserProfileRouter(rg *gin.RouterGroup, handler handler.UserProfile) {

	routes := []routing.Router{
		{
			Path:   "/users/:id/profile",
			Method: http.MethodPost,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.CreateUserProfile,
		},

		{
			Path:   "/users/:id/profile",
			Method: http.MethodPut,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.UpdateUserProfile,
		},

		{
			Path:   "/users/:id/profile",
			Method: http.MethodGet,
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.GetUserProfile,
		},
	}

	routing.RegisterRoute(rg, routes)

}
