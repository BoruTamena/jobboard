package application

import (
	"context"
	"log"

	"github.com/BoruTamena/jobboard/internal/constant/errors"
	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/internal/constant/persistencedb"
	"github.com/BoruTamena/jobboard/internal/storage"
)

type jobApplicationStorage struct {
	db persistencedb.PersistenceDb
}

func NewJobApplicationStorage(db persistencedb.PersistenceDb) storage.JobApplicationStorage {
	return &jobApplicationStorage{
		db,
	}
}

func (j *jobApplicationStorage) CreateJobApplication(ctx context.Context, jbApplication dto.AppilicationDto) (error, dto.AppilicationDto) {
	// your implementation goes here
	japp := db.JobApplication{
		JobId:       jbApplication.JobID,
		UserId:      jbApplication.UserID,
		CoverLetter: jbApplication.CoverLetter,
		Resume:      jbApplication.ResumeURL,
	}

	res := j.db.WithContext(ctx).Create(&japp)

	if res.Error != nil {
		return res.Error, dto.AppilicationDto{}
	}

	return nil, dto.AppilicationDto{
		UserID:      japp.UserId,
		JobID:       japp.JobId,
		ResumeURL:   japp.Resume,
		CoverLetter: japp.CoverLetter,
		Status:      japp.Status}
}

func (j *jobApplicationStorage) UpdateJobApplicationStatus(ctx context.Context, jobStatus dto.AppilicationStatusDto) (error, dto.AppilicationDto) {

	application := db.JobApplication{}

	res := j.db.WithContext(ctx).Where("id=?", jobStatus.ID).First(&application)

	if err := res.Error; err != nil {

		err = errors.DbReadErr.Wrap(err, "invlaid application id").
			WithProperty(errors.ErrorCode, 400)

		log.Println(err)

		return err, dto.AppilicationDto{}
	}

	application.Status = jobStatus.Status
	res = j.db.WithContext(ctx).Save(&application)

	if err := res.Error; err != nil {

		err = errors.WriteErr.Wrap(err, "error :: updating application status ").
			WithProperty(errors.ErrorCode, 500)

		log.Println(err)

		return err, dto.AppilicationDto{}
	}

	return nil, dto.AppilicationDto{
		UserID:      application.UserId,
		JobID:       application.JobId,
		CoverLetter: application.CoverLetter,
		ResumeURL:   application.Resume,
		Status:      application.Status,
	}

}

func (j *jobApplicationStorage) GetJobApplications(ctx context.Context, jobId string) ([]dto.AppilicationDto, error) {
	var applications []db.JobApplication

	res := j.db.WithContext(ctx).Where("job_id=?", jobId).Find(&applications)
	if res.Error != nil {
		return nil, res.Error
	}

	var applicationDtos []dto.AppilicationDto
	for _, application := range applications {
		applicationDtos = append(applicationDtos, dto.AppilicationDto{
			UserID:      application.UserId,
			JobID:       application.JobId,
			ResumeURL:   application.Resume,
			CoverLetter: application.CoverLetter,
			Status:      application.Status,
		})
	}

	return applicationDtos, nil
}
