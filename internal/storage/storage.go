package storage

import (
	"context"

	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/google/uuid"
)

type AuthzStorage interface {
	AddRole(user string, role string) error
	AddPolicy(role, obj, act string) error
	CheckPermision(sub, resource, act string) (bool, error)
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
	GetJobCategories(ctx context.Context) (error, []dto.JobCategoryDto)
}
