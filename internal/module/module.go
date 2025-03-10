package module

import (
	"context"

	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
)

// define module interface here

type AuthModule interface {
	RegisterUser(ctx context.Context, user dto.UserDto) (error, dto.UserDto)
	SignIn(ctx context.Context, userlg dto.UserLogin) (error, db.User)
}

type UserProfile interface {
	CreateUserProfile(ctx context.Context, user_id string, userProfileData dto.UserProfie) (db.UserProfile, error)
}
