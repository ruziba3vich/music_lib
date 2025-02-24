package storage

import (
	"strings"

	"github.com/google/uuid"
	"github.com/ruziba3vich/music_lib/internal/models"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateSong(song *models.Song) error {
	return s.db.Create(song).Error
}

func (s *Storage) GetSongsWithFilters(filter map[string]any, limit, offset int) ([]models.Song, error) {
	var songs []models.Song
	query := s.db.Where(filter).Where("is_deleted = false").Limit(limit).Offset(offset)
	if err := query.Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (s *Storage) GetSongs(limit, offset int) ([]models.Song, error) {
	var songs []models.Song
	query := s.db.Where("is_deleted = false").Limit(limit).Offset(offset)
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

	if err := s.db.Where("id = ? AND is_deleted = false", songUUID).First(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (s *Storage) UpdateSong(song *models.Song) error {
	return s.db.Where("id = ? AND is_deleted = false", song.ID).Save(song).Error
}

func (s *Storage) DeleteSong(id string) error {
	songUUID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.db.Model(&models.Song{}).Where("id = ?", songUUID).Update("is_deleted", true).Error
}

func (s *Storage) GetSongLyricsPaginated(id string, limit, offset int) ([]string, error) {
	var song models.Song
	songUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := s.db.Where("id = ? AND is_deleted = false", songUUID).First(&song).Error; err != nil {
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

func (s *Storage) GetSongsByArtist(artist string, limit int, offset int) ([]models.Song, error) {
	var songs []models.Song
	query := s.db.Where("? = ANY(artists) AND is_deleted = false", artist).Limit(limit).Offset(offset).Find(&songs)
	if query.Error != nil {
		return nil, query.Error
	}
	return songs, nil
}
