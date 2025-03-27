package application

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/module"
	"github.com/BoruTamena/jobboard/internal/storage"
)

type jobApplicationModule struct {
	jobApplicationStorage storage.JobApplicationStorage
}

func NewJobApplicationModule(japplicationStg storage.JobApplicationStorage) module.JobApplicationModule {
	return &jobApplicationModule{
		jobApplicationStorage: japplicationStg,
	}
}

func (j *jobApplicationModule) ApplyJob(cxt context.Context, jbappilicaiton dto.AppilicationDto) (error, dto.AppilicationDto) {

	// your implementation goes here
	if err := jbappilicaiton.Validate(); err != nil {

		err = errors.BadInput.Wrap(err, "invalid application input").
			WithProperty(errors.ErrorCode, 500)

		log.Println(err)

		return err, jbappilicaiton
	}

	// saving application to db

	err, jappilication := j.jobApplicationStorage.CreateJobApplication(cxt, jbappilicaiton)

	if err != nil {

		return err, jbappilicaiton

	}

	return nil, jappilication
}
