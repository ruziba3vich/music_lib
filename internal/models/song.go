package models

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Group     string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Lyrics    string    `gorm:"type:text"`
	IsDeleted bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
