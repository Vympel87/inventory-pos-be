package domain

import (
	"time"
)

type Category struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:255;not null"`
	Description string
	Products    []Product `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
