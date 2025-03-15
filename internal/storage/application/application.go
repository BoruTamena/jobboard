package application

import (
	"context"

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
