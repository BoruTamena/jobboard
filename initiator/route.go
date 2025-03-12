package initiator

import (
	"github.com/BoruTamena/jobboard/docs"
	authroute "github.com/BoruTamena/jobboard/internal/glue/routing/auth_route"
	userroute "github.com/BoruTamena/jobboard/internal/glue/routing/user_route"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouting(
	rg *gin.RouterGroup,
	handler Handler,
	module Module,
) {

	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Host = "localhost"

	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler))
	// intializing auth route
	authroute.InitAuthRoute(rg, handler.authHandler)
	// initializign userprofile route
	userroute.InitUserProfileRouter(rg, handler.userProfileHandler, module.authzModule)

}
