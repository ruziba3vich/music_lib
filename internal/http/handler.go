package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ruziba3vich/music_lib/internal/models"
	"github.com/ruziba3vich/music_lib/internal/repos"
)

type Handler struct {
	repo   repos.Repo
	logger *log.Logger
}

func NewHandler(repo repos.Repo, logger *log.Logger) *Handler {

	return &Handler{
		repo:   repo,
		logger: logger,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/songs", h.CreateSongHandler)
		api.GET("/songs/filtered", h.GetSongsWithFiltersHandler)
		api.GET("/songs", h.GetSongsHandler)
		api.GET("/songs/:id", h.GetSongByIDHandler)
		api.GET("/songs/:id/lyrics", h.GetSongLyricsPaginatedHandler)
		api.GET("/songs/artists", h.GetSongsByArtistHandler)
		api.PUT("/songs/:id", h.UpdateSongHandler)
		api.DELETE("/songs/:id", h.DeleteSongHandler)
	}
}

// CreateSongHandler handles song creation
func (h *Handler) CreateSongHandler(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		h.logger.Printf("ERROR: Failed to parse request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Generate a time-based UUID
	timestamp := time.Now().UnixNano()
	song.ID = uuid.NewSHA1(uuid.NameSpaceOID, []byte(time.Unix(0, timestamp).String()))

	if err := h.repo.CreateSong(&song); err != nil {
		h.logger.Printf("ERROR: Failed to create song: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create song"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

// GetSongByIDHandler handles fetching a song by ID
func (h *Handler) GetSongByIDHandler(c *gin.Context) {
	id := c.Param("id")

	song, err := h.repo.GetSongByID(id)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch song ID %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}

// GetSongsWithFiltersHandler handles fetching songs with filters and pagination
func (h *Handler) GetSongsWithFiltersHandler(c *gin.Context) {
	var filters map[string]any
	if err := c.ShouldBindJSON(&filters); err != nil {
		h.logger.Printf("ERROR: Failed to parse filters: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter format"})
		return
	}

	limit := getIntQueryParam(c, "limit", 10)
	offset := getIntQueryParam(c, "offset", 0)

	songs, err := h.repo.GetSongsWithFilters(filters, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch songs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// GetSongsHandler handles fetching songs with filters and pagination
func (h *Handler) GetSongsHandler(c *gin.Context) {

	limit := getIntQueryParam(c, "limit", 10)
	offset := getIntQueryParam(c, "offset", 0)

	songs, err := h.repo.GetSongs(limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch songs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// GetSongLyricsPaginatedHandler handles fetching paginated lyrics
func (h *Handler) GetSongLyricsPaginatedHandler(c *gin.Context) {
	id := c.Param("id")
	limit := getIntQueryParam(c, "limit", 10)
	offset := getIntQueryParam(c, "offset", 0)

	lyrics, err := h.repo.GetSongLyricsPaginated(id, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch lyrics for song ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch lyrics"})
		return
	}

	c.JSON(http.StatusOK, lyrics)
}

// UpdateSongHandler handles updating a song
func (h *Handler) UpdateSongHandler(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		h.logger.Printf("ERROR: Failed to parse request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.repo.UpdateSong(&song); err != nil {
		h.logger.Printf("ERROR: Failed to update song: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update song"})
		return
	}

	c.JSON(http.StatusOK, song)
}

// DeleteSongHandler handles soft deleting a song
func (h *Handler) DeleteSongHandler(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.DeleteSong(id); err != nil {
		h.logger.Printf("ERROR: Failed to delete song ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete song"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "song deleted"})
}

// Helper function to get integer query parameters with defaults
func getIntQueryParam(c *gin.Context, key string, defaultValue int) int {
	val, err := c.GetQuery(key)
	if !err {
		return defaultValue
	}
	var intVal int
	if _, err := fmt.Sscanf(val, "%d", &intVal); err != nil {
		return defaultValue
	}
	return intVal
}

func (h *Handler) GetSongsByArtistHandler(c *gin.Context) {
	var request struct {
		Artist string `json:"artist"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Printf("ERROR: Error occured in parsing requst")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.logger.Printf("ERROR: Invalid limit parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		h.logger.Printf("ERROR: Invalid offset parameter: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		return
	}

	if request.Artist == "" {
		h.logger.Println("ERROR: Artist name is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "artist name is required"})
		return
	}

	songs, err := h.repo.GetSongsByArtist(request.Artist, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to get songs for artist %s: %v", request.Artist, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}
