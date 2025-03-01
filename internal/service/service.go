package service

import (
	"context"
	"log"

	"github.com/ruziba3vich/music_lib/internal/models"
	"github.com/ruziba3vich/music_lib/internal/storage"
)

type Service struct {
	storage *storage.Storage
	logger  *log.Logger
}

// NewService creates a new service instance with logging
func NewService(storage *storage.Storage, logger *log.Logger) *Service {

	return &Service{
		storage: storage,
		logger:  logger,
	}
}

// CreateSong logs and calls storage.CreateSong
func (s *Service) CreateSong(ctx context.Context, song *models.Song) error {
	s.logger.Printf("INFO: Creating song: %+v", song)
	err := s.storage.CreateSong(ctx, song)
	if err != nil {
		s.logger.Printf("ERROR: Failed to create song: %v", err)
	}
	return err
}

// DeleteSong logs and calls storage.DeleteSong
func (s *Service) DeleteSong(ctx context.Context, id string) error {
	s.logger.Printf("INFO: Deleting song ID: %s", id)
	err := s.storage.DeleteSong(ctx, id)
	if err != nil {
		s.logger.Printf("ERROR: Failed to delete song ID %s: %v", id, err)
	}
	return err
}

// GetSongByID logs and calls storage.GetSongByID
func (s *Service) GetSongByID(ctx context.Context, id string) (*models.Song, error) {
	s.logger.Printf("INFO: Fetching song ID: %s", id)
	song, err := s.storage.GetSongByID(ctx, id)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch song ID %s: %v", id, err)
	}
	return song, err
}

// GetSongLyricsPaginated logs and calls storage.GetSongLyricsPaginated
func (s *Service) GetSongLyricsPaginated(ctx context.Context, id string, limit, offset int) ([]string, error) {
	s.logger.Printf("INFO: Fetching lyrics for song ID %s (limit: %d, offset: %d)", id, limit, offset)
	verses, err := s.storage.GetSongLyricsPaginated(ctx, id, limit, offset)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch lyrics for song ID %s: %v", id, err)
	}
	return verses, err
}

// GetSongsWithFilters logs and calls storage.GetSongs
func (s *Service) GetSongsWithFilters(ctx context.Context, filter map[string]any, limit, offset int) ([]models.Song, error) {
	s.logger.Printf("INFO: Fetching songs with filter %+v (limit: %d, offset: %d)", filter, limit, offset)
	songs, err := s.storage.GetSongsWithFilters(ctx, filter, limit, offset)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch songs: %v", err)
	}
	return songs, err
}

// GetSongs logs and calls storage.GetSongs
func (s *Service) GetSongs(ctx context.Context, limit, offset int) ([]models.Song, error) {
	s.logger.Printf("INFO: Fetching songs with (limit: %d, offset: %d)", limit, offset)
	songs, err := s.storage.GetSongs(ctx, limit, offset)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch songs: %v", err)
	}
	return songs, err
}

// UpdateSong logs and calls storage.UpdateSong
func (s *Service) UpdateSong(ctx context.Context, song *models.Song) error {
	s.logger.Printf("INFO: Updating song ID %s", song.ID)
	err := s.storage.UpdateSong(ctx, song)
	if err != nil {
		s.logger.Printf("ERROR: Failed to update song ID %s: %v", song.ID, err)
	}
	return err
}

func (s *Service) GetSongsByArtist(ctx context.Context, artist string, limit, offset int) ([]models.Song, error) {
	s.logger.Printf("INFO: Searching for songs by artist: %s, limit: %d, offset: %d", artist, limit, offset)

	songs, err := s.storage.GetSongsByArtist(ctx, artist, limit, offset)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch songs for artist %s: %v", artist, err)
		return nil, err
	}

	return songs, nil
}
