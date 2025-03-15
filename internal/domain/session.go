package domain

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID              `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID              `gorm:"type:uuid;not null"`
	User      User                   `gorm:"foreignKey:UserID"`
	Data      map[string]interface{} `gorm:"type:jsonb;not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
