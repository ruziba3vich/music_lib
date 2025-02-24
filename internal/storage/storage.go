package storage

import "gorm.io/gorm"

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) (*Storage, error) {

	return &Storage{db: db}, nil
}
