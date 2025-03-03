package handler

import "github.com/gin-gonic/gin"

// define your handlers interface here

type Auth interface {
	RegisterUser(ctx *gin.Context)
}
