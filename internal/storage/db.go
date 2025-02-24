package storage

import (
	"fmt"
	"log"

	"github.com/ruziba3vich/music_lib/internal/models"
	"github.com/ruziba3vich/music_lib/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
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

	return db, nil
}
