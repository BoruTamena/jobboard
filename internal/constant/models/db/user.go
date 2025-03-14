package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// hook for base model
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

type User struct {
	BaseModel
	UserName string
	Email    string
	Password string
	Profiles []UserProfile `gorm:"foreignKey:UserID"`
}

type UserProfile struct {
	BaseModel
	UserID   uuid.UUID `gorm:"not null"`
	Bio      string
	Location string
}
