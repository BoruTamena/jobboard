package routing

import "github.com/gin-gonic/gin"

type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func RegisterRoute(rgroup *gin.RouterGroup, routes []Router) {

	for _, route := range routes {

		rgroup.Handle(route.Method, route.Path, route.Handler)

	}

}
