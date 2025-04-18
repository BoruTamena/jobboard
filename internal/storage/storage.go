package storage

import (
	"context"

	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/google/uuid"
)

type AuthzStorage interface {
	AddRole(userId string, role string) error
	AddPolicy(role, obj, act string) error
	CheckPermision(sub, resource, act string) (bool, error)
	CreateRole(userId string, role string) error
}

type AuthStorage interface {
	CreateUser(ctx context.Context, user dto.UserDto) (error, db.User)
	GetUserByEmail(ctx context.Context, userlg dto.UserLogin) (error, db.User)
}

type UserProfie interface {
	CreateUserProfile(ctx context.Context, user_id uuid.UUID, userProfile dto.UserProfie) (db.UserProfile, error)
	UpdateUserProfile(ctx context.Context, user_id uuid.UUID, userProfile dto.UserProfie) (db.UserProfile, error)
	GetUserProfile(ctx context.Context, user_id uuid.UUID) (db.User, error)
}

type JobStorage interface {
	CreateJob(ctx context.Context, job dto.JobDto) (error, dto.JobDto)
	GetJobs(ctx context.Context, limit, offset int) (error, []dto.JobDto)
	CreateJobCategory(ctx context.Context, jobCategory dto.JobCategoryDto) (error, dto.JobCategoryDto)
	GetJobCategories(ctx context.Context) (error, []dto.JobCategoryDto)
	UpdateJobStatus(ctx context.Context, jobStatus dto.JobStatusDto) (error, dto.JobDto)
}

type JobApplicationStorage interface {
	CreateJobApplication(ctx context.Context, jbApplication dto.AppilicationDto) (error, dto.AppilicationDto)
	UpdateJobApplicationStatus(ctx context.Context, jobStatus dto.AppilicationStatusDto) (error, dto.AppilicationDto)
	GetJobApplications(ctx context.Context, jobId string) ([]dto.AppilicationDto, error)
}
