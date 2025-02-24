package storage

import (
	"strings"

	"github.com/google/uuid"
	"github.com/ruziba3vich/music_lib/internal/models"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) CreateSong(song *models.Song) error {
	return s.DB.Create(song).Error
}

func (s *Storage) GetSongs(filter map[string]any, limit, offset int) ([]models.Song, error) {
	var songs []models.Song
	query := s.DB.Where(filter).Where("is_deleted = false").Limit(limit).Offset(offset)
	if err := query.Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (s *Storage) GetSongByID(id string) (*models.Song, error) {
	var song models.Song
	songUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := s.DB.Where("id = ? AND is_deleted = false", songUUID).First(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (s *Storage) UpdateSong(song *models.Song) error {
	return s.DB.Where("id = ? AND is_deleted = false", song.ID).Save(song).Error
}

func (s *Storage) DeleteSong(id string) error {
	songUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.DB.Model(&models.Song{}).Where("id = ?", songUUID).Update("is_deleted", true).Error
}

func (s *Storage) GetSongLyricsPaginated(id string, limit, offset int) ([]string, error) {
	var song models.Song
	songUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := s.DB.Where("id = ? AND is_deleted = false", songUUID).First(&song).Error; err != nil {
		return nil, err
	}

	verses := strings.Split(song.Lyrics, "\n\n")

	start := offset
	end := offset + limit
	if start >= len(verses) {
		return []string{}, nil
	}
	if end > len(verses) {
		end = len(verses)
	}

	return verses[start:end], nil
}
