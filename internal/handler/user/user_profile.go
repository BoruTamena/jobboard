package user

import (
	"log"
	"net/http"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/models/response"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/gin-gonic/gin"
)

type userProfile struct {
	userModule module.UserProfile
}

func InitUserProfileHandler(umd module.UserProfile) handler.UserProfile {

	return &userProfile{
		userModule: umd,
	}

}

// @Summary create new user
// @Tags profile
// @Description create profile
// @Accept json
// @Produce json
// @Param id path string true "user id "
// @Param profile body  dto.UserProfie true "user profile request body"
// @Success 201 {object} response.Response
// @Router /user/{id}/profile [post]
func (up *userProfile) CreateUserProfile(ctx *gin.Context) {

	var userProfile dto.UserProfie

	Id := ctx.Param("id")

	if err := ctx.Bind(&userProfile); err != nil {
		// bind error
		err = errors.BadInput.Wrap(err, "bind error").
			WithProperty(errors.ErrorCode, 400)

		log.Println("user bind error:-", err.Error())

		return
	}

	// sending to module
	profile, err := up.userModule.CreateUserProfile(ctx, Id, userProfile)
	if err != nil {
		ctx.Error(err)

		return
	}

	// returning json response
	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   profile,
	})

}

// @Summary get users
// @Tags profile
// @Description create profile
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Param  body body dto.UserProfie true "user_id"
// @Success 201 {object} response.Response
// @Router /user/{id}/profile [put]
func (up *userProfile) UpdateUserProfile(ctx *gin.Context) {
	var userProfile dto.UserProfie

	Id := ctx.Param("id")

	if Id == "" {

		log.Println("id=====", Id)
		err := errors.BadInput.New("Id is required").
			WithProperty(errors.ErrorCode, 500)
		log.Println("id is required")
		ctx.Error(err)

		return

	}

	if err := ctx.Bind(&userProfile); err != nil {
		// bind error
		err = errors.BadInput.Wrap(err, "bind error").
			WithProperty(errors.ErrorCode, 400)

		log.Println("user bind error:-", err.Error())

		ctx.Error(err)

		return
	}

	// sending to module
	profile, err := up.userModule.UpdateUserProfile(ctx, Id, userProfile)

	if err != nil {
		ctx.Error(err)
		return
	}
	// returning json response
	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   profile,
	})

}

// @Summary get users
// @Tags profile
// @Description create profile
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} response.Response
// @Router /user/{id}/profile [get]
func (up *userProfile) GetUserProfile(ctx *gin.Context) {

	Id := ctx.Param("id")

	if Id == "" {

		err := errors.BadInput.New("Id is required").
			WithProperty(errors.ErrorCode, 500)
		log.Println("id is required")
		ctx.Error(err)

		return

	}

	profile, err := up.userModule.GetUserProfile(ctx, Id)

	if err != nil {

		ctx.Error(err)
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   profile,
	})

}
