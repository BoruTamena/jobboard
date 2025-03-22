package applicationroute

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/glue/routing"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func InitApplicationRoute(rg *gin.RouterGroup, handler handler.JobApplication) {

	routes := []routing.Router{

		{
			Method: http.MethodPost,
			Path:   "job/apply",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.ApplyJob,
		},

		{
			Method: http.MethodPatch,
			Path:   "job/:id/status",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.ErrorMiddleWare(),
			},

			Handler: handler.UpdateApplicationStatus,
		},

		{
			Method: http.MethodGet,
			Path:   "job/:id?/appilications",
			Middlewares: []gin.HandlerFunc{
				middleware.AuthMiddleware(),
				middleware.ErrorMiddleWare(),
			},
			Handler: handler.GetJobApplicationByID,
		},
	}

	routing.RegisterRoute(rg, routes)

}
