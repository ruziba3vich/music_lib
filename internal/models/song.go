package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Song struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Artists     pq.StringArray `gorm:"type:text[]" json:"artists"`
	Group       string         `gorm:"not null" json:"group"`
	Name        string         `gorm:"not null" json:"name"`
	Lyrics      string         `gorm:"type:text" json:"lyrics"`
	IsDeleted   bool           `gorm:"default:false" json:"-"`
	ReleaseDate time.Time      `json:"release_date"`
	CreatedAt   time.Time
}

type RequestByArtistName struct {
	Artist string `json:"artist"`
}
