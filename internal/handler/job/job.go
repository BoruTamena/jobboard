package job

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

type jobHandler struct {
	jobModule module.Job
}

func NewJobHandler(jmodule module.Job) handler.Job {
	return jobHandler{
		jobModule: jmodule,
	}
}

// @tags Job
// @Summary Create a job
// @Description Create a job
// @Accept json
// @Produce json
// @Param job body  dto.JobDto true "job request body"
// @Success 201 {object} response.Response
// @Router /job [post]
// @security BearerAuth
func (j jobHandler) CreateJob(ctx *gin.Context) {
	// create job
	var job dto.JobDto

	if err := ctx.Bind(&job); err != nil {

		err = errors.BadInput.Wrap(err, "bind error").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return

	}
	// calling job module
	err, job_res := j.jobModule.CreateJob(ctx, job)
	if err != nil {

		ctx.Error(err)
		return

	}
	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   job_res,
	})

}

// @tags Job
// @Summary Get a job
// @Description Get a job
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} response.Response
// @Router /job [get]
// @security BearerAuth
func (j jobHandler) GetJob(ctx *gin.Context) {
	// create job
	var pagination dto.Pagination

	if err := ctx.BindQuery(&pagination); err != nil {

		err = errors.BadInput.Wrap(err, "bind error").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return

	}

	// calling job module

	err, jobs := j.jobModule.GetJobs(ctx, pagination)

	if err != nil {

		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Status: string(response.Sucess),
		Data:   jobs,
	})

}

// @tags Job
// @Summary Create a job category
// @Description Create a job category
// @Accept json
// @Produce json
// @Param job body  dto.JobCategoryDto true "job category request body"
// @Success 201 {object} response.Response
// @Router /job/category [post]
// @security BearerAuth
func (j jobHandler) CreateJobCategory(ctx *gin.Context) {
	// create job
	var jobCategory dto.JobCategoryDto

	if err := ctx.Bind(&jobCategory); err != nil {

		err = errors.BadInput.Wrap(err, "bind error").
			WithProperty(errors.ErrorCode, 400)

		ctx.Error(err)

		return

	}
	// calling job module
	err, job_res := j.jobModule.CreateJobCategory(ctx, jobCategory)
	if err != nil {

		ctx.Error(err)
		return

	}
	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   job_res,
	})

}

// @tags Job
// @Summary Get a job category
// @Description Get a job category
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /job/category [get]
// @security BearerAuth
func (j jobHandler) GetJobCategory(ctx *gin.Context) {

	// calling job module
	err, job_categories := j.jobModule.GetJobCategories(ctx)

	if err != nil {

		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Status: string(response.Sucess),
		Data:   job_categories,
	})

}

// @tags Job
// @Summary update job status
// @Description update job status
// @Accept json
// @Produce json
// @Param id path string true "job id is required"
// @Param jobStatus body dto.JobStatusDto true "job status is required"
// @Success 201 {object} response.Response
// @Router /job/{id}/status [patch]
// @security BearerAuth
func (j jobHandler) UpdateJobStatus(ctx *gin.Context) {

	var jobStatus dto.JobStatusDto

	if err := ctx.Bind(&jobStatus); err != nil {

		err := errors.BadInput.Wrap(err, "couldn't bind user input").
			WithProperty(errors.ErrorCode, 500)

		log.Println(err)

		return

	}

	err, job := j.jobModule.UpdateJobStatus(ctx, jobStatus)

	if err != nil {

		ctx.Error(err)

		return
	}

	ctx.JSON(http.StatusCreated, response.Response{
		Status: string(response.Sucess),
		Data:   job,
	})

}
