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
// @Tags user
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
// @Tags user
// @Description create profile
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} response.Response
// @Router /user/{id}/profile [get]
func (up *userProfile) GetUserProfile(ctx *gin.Context) {

	// Id := ctx.Param("id")

}
