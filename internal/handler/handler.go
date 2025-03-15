package handler

import "github.com/gin-gonic/gin"

// define your handlers interface here

type Auth interface {
	RegisterUser(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type UserProfile interface {
	CreateUserProfile(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
	UpdateUserProfile(ctx *gin.Context)
}

type Job interface {
	CreateJob(ctx *gin.Context)
	GetJob(ctx *gin.Context)
	GetJobCategory(ctx *gin.Context)
}
