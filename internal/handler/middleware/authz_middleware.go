package middleware

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/constant/models/response"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/gin-gonic/gin"
)

func Authorize(authz module.AuthzModule) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// authorizing user
		user := ctx.GetString("user")
		resource := ctx.GetString("resource")
		action := ctx.Request.Method

		// checking if user is allowed to take action

		permission := authz.CheckUserPermission(user, resource, action)

		if !permission {

			ctx.AbortWithStatusJSON(http.StatusForbidden, response.Response{
				Status: string(response.Fail),
				Data:   "you are not allowed to take this is action",
			})

			return

		}

		// allow if user has permission
		ctx.Next()

	}

}
