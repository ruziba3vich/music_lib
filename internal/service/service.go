package service

import (
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
func (s *Service) CreateSong(song *models.Song) error {
	s.logger.Printf("INFO: Creating song: %+v", song)
	err := s.storage.CreateSong(song)
	if err != nil {
		s.logger.Printf("ERROR: Failed to create song: %v", err)
	}
	return err
}

// DeleteSong logs and calls storage.DeleteSong
func (s *Service) DeleteSong(id string) error {
	s.logger.Printf("INFO: Deleting song ID: %s", id)
	err := s.storage.DeleteSong(id)
	if err != nil {
		s.logger.Printf("ERROR: Failed to delete song ID %s: %v", id, err)
	}
	return err
}

// GetSongByID logs and calls storage.GetSongByID
func (s *Service) GetSongByID(id string) (*models.Song, error) {
	s.logger.Printf("INFO: Fetching song ID: %s", id)
	song, err := s.storage.GetSongByID(id)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch song ID %s: %v", id, err)
	}
	return song, err
}

// GetSongLyricsPaginated logs and calls storage.GetSongLyricsPaginated
func (s *Service) GetSongLyricsPaginated(id string, limit, offset int) ([]string, error) {
	s.logger.Printf("INFO: Fetching lyrics for song ID %s (limit: %d, offset: %d)", id, limit, offset)
	verses, err := s.storage.GetSongLyricsPaginated(id, limit, offset)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch lyrics for song ID %s: %v", id, err)
	}
	return verses, err
}

// GetSongs logs and calls storage.GetSongs
func (s *Service) GetSongs(filter map[string]any, limit, offset int) ([]models.Song, error) {
	s.logger.Printf("INFO: Fetching songs with filter %+v (limit: %d, offset: %d)", filter, limit, offset)
	songs, err := s.storage.GetSongs(filter, limit, offset)
	if err != nil {
		s.logger.Printf("ERROR: Failed to fetch songs: %v", err)
	}
	return songs, err
}

// UpdateSong logs and calls storage.UpdateSong
func (s *Service) UpdateSong(song *models.Song) error {
	s.logger.Printf("INFO: Updating song ID %s", song.ID)
	err := s.storage.UpdateSong(song)
	if err != nil {
		s.logger.Printf("ERROR: Failed to update song ID %s: %v", song.ID, err)
	}
	return err
}

/*
	// Create log file
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Multi-writer for both file and console logging
	multiWriter := log.MultiWriter(os.Stdout, logFile)
	logger := log.New(multiWriter, "[MusicLib] ", log.Ldate|log.Ltime|log.Lshortfile)

*/
