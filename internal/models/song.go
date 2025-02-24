package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Group       string `gorm:"not null"`
	Song        string `gorm:"not null"`
	ReleaseDate string `gorm:"not null"`
	Text        string `gorm:"type:text;not null"`
	Link        string `gorm:"not null"`
}
