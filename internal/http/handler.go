package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/ruziba3vich/music_lib/docs"
	"github.com/ruziba3vich/music_lib/internal/models"
	"github.com/ruziba3vich/music_lib/internal/repos"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

// @Summary Create a new song
// @Description Adds a new song to the database
// @Accept json
// @Tags songs
// @Produce json
// @Param song body models.Song true "Song object"
// @Success 201 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs [post]
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

	if err := h.repo.CreateSong(c, &song); err != nil {
		h.logger.Printf("ERROR: Failed to create song: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create song"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

// @Summary Get a song by ID
// @Description Fetches a song from the database using its ID
// @Produce json
// @Tags songs
// @Param id path string true "Song ID"
// @Success 200 {object} models.Song
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs/{id} [get]
func (h *Handler) GetSongByIDHandler(c *gin.Context) {
	id := c.Param("id")

	song, err := h.repo.GetSongByID(c, id)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch song ID %s: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}

// GetSongsWithFiltersHandler handles fetching songs with filters and pagination
// @Summary Get songs with filters and pagination
// @Description Fetches songs based on filters provided as query parameters
// @Tags songs
// @Accept json
// @Produce json
// @Param name query string false "Filter by song name"
// @Param artist query string false "Filter by artist name"
// @Param genre query string false "Filter by genre"
// @Param limit query int false "Limit the number of results" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]string "failed to fetch songs"
// @Router /api/songs/filtered [get]
func (h *Handler) GetSongsWithFiltersHandler(c *gin.Context) {
	filters := map[string]any{}

	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}
	if artist := c.Query("artist"); artist != "" {
		filters["artist"] = artist
	}
	if genre := c.Query("genre"); genre != "" {
		filters["genre"] = genre
	}

	limit := getIntQueryParam(c, "limit", 10)
	offset := getIntQueryParam(c, "offset", 0)

	songs, err := h.repo.GetSongsWithFilters(c, filters, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch songs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// @Summary Get all songs
// @Description Fetches a list of songs with optional pagination
// @Produce json
// @Tags songs
// @Param limit query int false "Limit the number of results" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]string "failed to fetch songs"
// @Router /api/songs [get]
func (h *Handler) GetSongsHandler(c *gin.Context) {

	limit := getIntQueryParam(c, "limit", 10)
	offset := getIntQueryParam(c, "offset", 0)

	songs, err := h.repo.GetSongs(c, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch songs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// @Summary Get song lyrics with pagination
// @Description Fetches paginated lyrics for a song by ID
// @Produce json
// @Tags songs
// @Param id path string true "Song ID"
// @Param limit query int false "Limit the number of results" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} string
// @Failure 500 {object} map[string]string "failed to fetch lyrics"
// @Router /api/songs/{id}/lyrics [get]
func (h *Handler) GetSongLyricsPaginatedHandler(c *gin.Context) {
	id := c.Param("id")
	limit := getIntQueryParam(c, "limit", 10)
	offset := getIntQueryParam(c, "offset", 0)

	lyrics, err := h.repo.GetSongLyricsPaginated(c, id, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to fetch lyrics for song ID %s: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch lyrics"})
		return
	}

	c.JSON(http.StatusOK, lyrics)
}

// UpdateSongHandler handles updating a song
// @Summary Update a song
// @Description Update the details of an existing song
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Song data"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Failed to update song"
// @Router /api/songs/{id} [put]
func (h *Handler) UpdateSongHandler(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		h.logger.Printf("ERROR: Failed to parse request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.repo.UpdateSong(c, &song); err != nil {
		h.logger.Printf("ERROR: Failed to update song: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update song"})
		return
	}

	c.JSON(http.StatusOK, song)
}

// @Summary Delete a song
// @Description Deletes a song by ID
// @Tags songs
// @Param id path string true "Song ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs/{id} [delete]
func (h *Handler) DeleteSongHandler(c *gin.Context) {
	id := c.Param("id")

	if err := h.repo.DeleteSong(c, id); err != nil {
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

// GetSongsByArtistHandler handles fetching songs by a specific artist
// @Summary Get songs by artist
// @Description Fetch songs by the given artist name with pagination
// @Tags songs
// @Accept json
// @Produce json
// @Param artist query string true "Artist name"
// @Param limit query int false "Limit (default: 10)"
// @Param offset query int false "Offset (default: 0)"
// @Success 200 {array} models.Song
// @Failure 400 {object} map[string]string "Invalid request parameters"
// @Failure 500 {object} map[string]string "Failed to fetch songs"
// @Router /api/songs/artists [get]
func (h *Handler) GetSongsByArtistHandler(c *gin.Context) {
	artist := c.Query("artist")
	if artist == "" {
		h.logger.Println("ERROR: Artist name is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "artist name is required"})
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

	songs, err := h.repo.GetSongsByArtist(c, artist, limit, offset)
	if err != nil {
		h.logger.Printf("ERROR: Failed to get songs for artist %s: %v", artist, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}
