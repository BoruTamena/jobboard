package routing

import "github.com/gin-gonic/gin"

type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
}

func RegisterRoute(rgroup *gin.RouterGroup, routes []Router) {

	var handler []gin.HandlerFunc

	for _, route := range routes {

		handler = append(handler, route.Middlewares...)
		handler = append(handler, route.Handler)

		rgroup.Handle(route.Method, route.Path, handler...)

	}

}
