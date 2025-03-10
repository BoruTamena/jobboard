package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BoruTamena/jobboard/helper"
	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/models/response"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authModule module.AuthModule
}

func AuthHandler(authmodule module.AuthModule) handler.Auth {
	return &authHandler{
		authModule: authmodule,
	}
}

// @Summary create new user
// @Tags auth
// @Description create new user on boardapi
// @Accept json
// @Produce json
// @Param body body dto.UserDto true "user request body"
// @Success 201 {object} response.Response
// @Router /user [post]
func (ath *authHandler) RegisterUser(ctx *gin.Context) {

	var User dto.UserDto

	if err := ctx.Bind(&User); err != nil {

		// bind error
		err = errors.BadInput.Wrap(err, "bind error").
			WithProperty(errors.ErrorCode, 400)

		log.Println("user bind error:-", err.Error())

		return

	}

	err, user := ath.authModule.RegisterUser(ctx.Request.Context(), User)

	if err != nil {

		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   user,
	})

}

// @Summary login to the system
// @Tags auth
// @Description user login
// @Accept json
// @Produce json
// @Param body body dto.UserLogin true "user login request body"
// @Success 201 {object} response.Response
// @Router /signin [post]
func (ath *authHandler) SignIn(ctx *gin.Context) {

	var UserLg dto.UserLogin

	if err := ctx.Bind(&UserLg); err != nil {

		err = errors.BadInput.Wrap(err, "user bind error").
			WithProperty(errors.ErrorCode, 400)

		log.Println("user bind error:-", err.Error())

		ctx.Error(err)

		return

	}

	fmt.Println("calling auth module signin")
	err, userLg := ath.authModule.SignIn(ctx, UserLg)

	if err != nil {

		ctx.Error(err)

		return
	}

	access_token, refresh_token, err := helper.CreateToken(userLg)

	if err != nil {

		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data: map[string]string{
			"access_token":  access_token,
			"refresh_token": refresh_token,
		},
	})

}
