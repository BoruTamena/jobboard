package application

import (
	"net/http"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/models/response"
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/gin-gonic/gin"
)

type jobApplicationHandler struct {
	jApplicationMd module.JobApplicationModule
}

func NewJobApplication(jAmd module.JobApplicationModule) handler.JobApplicationHandler {
	return &jobApplicationHandler{
		jApplicationMd: jAmd,
	}
}

// @tags application
// @Description job application
// @Summary apply for job
// @Accept json
// @Produce json
// @Param jobapplication body dto.AppilicationDto true "job application body is required"
// @Success 201 {object} response.Response
// @security BearerAuth
// @Router /job/apply [post]
func (j jobApplicationHandler) ApplyJob(ctx *gin.Context) {

	var jApplication dto.AppilicationDto

	if err := ctx.Bind(&jApplication); err != nil {

		err = errors.BadInput.Wrap(err, "can't bind job application ").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return
	}

	err, japplication := j.jApplicationMd.ApplyJob(ctx, jApplication)

	if err != nil {

		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   japplication,
	})
}
