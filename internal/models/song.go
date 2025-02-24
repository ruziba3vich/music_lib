package models

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Artists     []string  `gorm:"type:array" json:"artists"`
	Group       string    `gorm:"not null" json:"group"`
	Name        string    `gorm:"not null" json:"name"`
	Lyrics      string    `gorm:"type:text" json:"lyrics"`
	IsDeleted   bool      `gorm:"default:false" json:"-"`
	ReleaseDate time.Time `json:"release_date"`
	CreatedAt   time.Time
}
