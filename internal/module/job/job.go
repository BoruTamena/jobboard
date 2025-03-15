package job

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/storage"
)

type jobModule struct {

	// job storage
	jobStorage storage.JobStorage
}

// NewJobModule creates a new job module
func NewJobModule(jobStorage storage.JobStorage) module.Job {
	return jobModule{
		jobStorage: jobStorage,
	}
}

func (j jobModule) CreateJob(ctx context.Context, job dto.JobDto) (error, dto.JobDto) {
	// validation
	if err := job.Validate(); err != nil {
		err := errors.BadInput.Wrap(err, "validation error").
			WithProperty(errors.ErrorCode, 400)

		log.Println(err)
		return err, job

	}

	// save job to storage
	err, jb := j.jobStorage.CreateJob(ctx, job)

	if err != nil {
		return err, job
	}

	return nil, jb

}

func (j jobModule) GetJobs(cxt context.Context, pagination dto.Pagination) (error, []dto.JobDto) {

	if err := pagination.Validate(); err != nil {
		err := errors.BadInput.Wrap(err, "validation error").
			WithProperty(errors.ErrorCode, 400)

		log.Println(err)
		return err, nil
	}
	limits := pagination.PerPage
	offset := (pagination.Page - 1) * limits
	// get job from storage

	err, jobs := j.jobStorage.GetJobs(cxt, limits, offset)

	if err != nil {

		return err, nil

	}

	return nil, jobs

}

func (j jobModule) GetJobCategories(cxt context.Context) (error, []dto.JobCategoryDto) {
	// get job category from storage

	err, jobCategories := j.jobStorage.GetJobCategories(cxt)

	if err != nil {
		return err, nil
	}

	return nil, jobCategories
}
