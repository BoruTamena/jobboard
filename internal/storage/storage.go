package storage

import (
	"context"

	"github.com/BoruTamena/jobboard/internal/constant/models/db"
	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
)

// define U storage interface here

type AuthStorage interface {
	CreateUser(ctx context.Context, user dto.UserDto) (error, db.User)
	GetUserByEmail(ctx context.Context, userlg dto.UserLogin) (error, db.User)
}
