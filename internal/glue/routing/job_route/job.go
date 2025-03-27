package jobroute

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/glue/routing"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

func InitJobRoute(rg *gin.RouterGroup, handler handler.Job) {
	routes := []routing.Router{
		{
			Method: http.MethodPost,
			Path:   "/job",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.CreateJob,
		},
		{
			Method: http.MethodGet,
			Path:   "/job",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.GetJob,
		},
		{
			Method: http.MethodPost,
			Path:   "/job/category",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.CreateJobCategory,
		},
		{
			Method: http.MethodGet,
			Path:   "/job/category",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.GetJobCategory,
		},

		{
			Method: http.MethodPatch,
			Path:   "/job/:id/status",
			Middlewares: []gin.HandlerFunc{
				middleware.ErrorMiddleWare(),
				middleware.AuthMiddleware(),
			},
			Handler: handler.UpdateJobStatus,
		},
	}

	routing.RegisterRoute(rg, routes)

}
