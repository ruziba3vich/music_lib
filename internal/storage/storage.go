package storage

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) (*Storage, error) {

	return &Storage{db: db}, nil
}

/*
dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", err.Error())
	}

	err = db.AutoMigrate(&models.Song{})
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %s", err.Error())
	}

	log.Println("Database connected and migrated successfully!")

*/
