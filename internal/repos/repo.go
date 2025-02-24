package repos

import (
	"context"

	"github.com/ruziba3vich/music_lib/internal/models"
)

type (
	Repo interface {
		CreateSong(context.Context, *models.Song) error
		DeleteSong(context.Context, string) error
		GetSongByID(context.Context, string) (*models.Song, error)
		GetSongLyricsPaginated(context.Context, string, int, int) ([]string, error)
		GetSongsWithFilters(context.Context, map[string]any, int, int) ([]models.Song, error)
		GetSongs(context.Context, int, int) ([]models.Song, error)
		UpdateSong(context.Context, *models.Song) error
		GetSongsByArtist(context.Context, string, int, int) ([]models.Song, error)
	}
)
