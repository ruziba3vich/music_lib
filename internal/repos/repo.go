package repos

import "github.com/ruziba3vich/music_lib/internal/models"

type (
	Repo interface {
		CreateSong(*models.Song) error
		DeleteSong(string) error
		GetSongByID(string) (*models.Song, error)
		GetSongLyricsPaginated(string, int, int) ([]string, error)
		GetSongs(map[string]any, int, int) ([]models.Song, error)
		UpdateSong(*models.Song) error
	}
)
