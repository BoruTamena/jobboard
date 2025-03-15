package job

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
)

type jobStorage struct {
	db persistencedb.PersistenceDb
}

// NewJobStorage creates a new job storage

func NewJobStorage(db persistencedb.PersistenceDb) storage.JobStorage {

	return &jobStorage{
		db: db,
	}
}

func (js *jobStorage) CreateJob(ctx context.Context, job dto.JobDto) (error, dto.JobDto) {
	// create job
	job_db := db.Job{
		Title:       job.JobTitle,
		Description: job.JobDescription,
		Location:    job.Location,
		JobType:     job.JobType,
		CategoryId:  job.CategoryId,
	}

	res := js.db.WithContext(ctx).Create(&job_db)

	if err := res.Error; err != nil {

		err := errors.WriteErr.Wrap(err, "error creating job").
			WithProperty(errors.ErrorCode, 500)

		log.Println(err)
		return err, job
	}

	return nil, job

}

func (js *jobStorage) GetJobs(ctx context.Context, limit, offset int) (error, []dto.JobDto) {
	// get job from storage
	var jobs []db.Job
	res := js.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&jobs)

	if err := res.Error; err != nil {
		err := errors.DbReadErr.Wrap(err, "error reading job").
			WithProperty(errors.ErrorCode, 500)

		log.Println(err)
		return err, nil
	}

	job_dtos := []dto.JobDto{}

	for _, job := range jobs {
		job_dtos = append(job_dtos, dto.JobDto{
			JobTitle:       job.Title,
			JobDescription: job.Description,
			Location:       job.Location,
			JobType:        job.JobType,
		})
	}

	return nil, job_dtos
}

func (js *jobStorage) GetJobCategories(ctx context.Context) (error, []dto.JobCategoryDto) {

	var jobCategories []db.JobCategory

	res := js.db.WithContext(ctx).Find(&jobCategories)

	if err := res.Error; err != nil {
		err := errors.DbReadErr.Wrap(err, "error reading job category").
			WithProperty(errors.ErrorCode, 500)

		log.Println(err)
		return err, nil
	}

	jobCategoryDtos := []dto.JobCategoryDto{}

	for _, jobCategory := range jobCategories {
		jobCategoryDtos = append(jobCategoryDtos, dto.JobCategoryDto{
			ID:           jobCategory.ID,
			CategoryName: jobCategory.Name,
			Description:  jobCategory.Description,
		})
	}

	return nil, jobCategoryDtos

}
