package module

import (
	"context"

	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
)

type AuthzModule interface {
	CheckUserPermission(user, resource, act string) bool
	AddRoleForUser(ctx context.Context, user, role string) error
	AddPolicy(ctx context.Context, role, obj, act string) error
}

type AuthModule interface {
	RegisterUser(ctx context.Context, user dto.UserDto) (error, dto.UserDto)
	SignIn(ctx context.Context, userlg dto.UserLogin) (error, map[string]interface{})
}

type UserProfile interface {
	CreateUserProfile(ctx context.Context, user_id string, userProfileData dto.UserProfie) (db.UserProfile, error)
	UpdateUserProfile(ctx context.Context, user_id string, userProfileData dto.UserProfie) (db.UserProfile, error)
	GetUserProfile(ctx context.Context, user_id string) (db.User, error)
}

type Job interface {
	CreateJob(ctx context.Context, job dto.JobDto) (error, dto.JobDto)
	GetJobs(cxt context.Context, pagination dto.Pagination) (error, []dto.JobDto)
	CreateJobCategory(cxt context.Context, jobCategory dto.JobCategoryDto) (error, dto.JobCategoryDto)
	GetJobCategories(cxt context.Context) (error, []dto.JobCategoryDto)
	UpdateJobStatus(ctx context.Context, jobStatus dto.JobStatusDto) (error, dto.JobDto)
}

type JobApplicationModule interface {
	ApplyJob(cxt context.Context, jbappilicaiton dto.AppilicationDto) (error, dto.AppilicationDto)
	UpdateApplicationStatus(ctx context.Context, applicationStatus dto.AppilicationStatusDto) (error, dto.AppilicationDto)
	GetJobApplicationByJobID(ctx context.Context, jobID string) (error, []dto.AppilicationDto)
}
