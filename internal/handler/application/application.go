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

func NewJobApplication(jAmd module.JobApplicationModule) handler.JobApplication {
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

// @tags application
// @Description accept/reject job application
// @Summary updating a job application status
// @Accept json
// @Produce json
// @Param id path string true "job application id is requried"
// @Param jobapplication body dto.AppilicationStatusDto true "job status  body is required"
// @Success 201 {object} response.Response
// @security BearerAuth
// @Router /job/:id/status [patch]
func (j jobApplicationHandler) UpdateApplicationStatus(ctx *gin.Context) {

	var ApplicationStatus dto.AppilicationStatusDto

	Id := ctx.Param("id")

	if Id == "" {
		err := errors.BadInput.New("application id ").
			WithProperty(errors.ErrorCode, 500)

		ctx.Error(err)

		return
	}

	// binding body
	if err := ctx.Bind(&ApplicationStatus); err != nil {

		err = errors.BadInput.Wrap(err, "err:: bind application status ").
			WithProperty(errors.ErrorCode, 500)

		ctx.Error(err)

		return

	}

	ApplicationStatus.ID = Id

	err, application := j.jApplicationMd.UpdateApplicationStatus(ctx, ApplicationStatus)

	if err != nil {

		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   application,
	})
}

// @tags application
// @Description get job application by id
// @Summary get job application
// @Produce json
// @Param id path string true "job application id is required"
// @Success 200 {object} response.Response
// @security BearerAuth
// @Router /job/:id/applications [get]
func (j jobApplicationHandler) GetJobApplicationByID(ctx *gin.Context) {

	Id := ctx.Param("id")

	if Id == "" {
		err := errors.BadInput.New("application id is required").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return
	}

	err, application := j.jApplicationMd.GetJobApplicationByJobID(ctx, Id)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Status: string(response.Sucess),
		Data:   application,
	})
}
