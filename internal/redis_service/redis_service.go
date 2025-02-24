package redisservice

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ruziba3vich/music_lib/internal/models"
	"github.com/ruziba3vich/music_lib/pkg/config"
)

// RedisService handles caching songs in Redis.
type RedisService struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisService initializes a RedisService with TTL from config.
func NewRedisService(client *redis.Client, cfg *config.Config) *RedisService {
	return &RedisService{
		client: client,
		ttl:    time.Duration(cfg.RedisTTL) * time.Second, // Read TTL from config
	}
}

// AddSong caches a song in Redis with an expiration time.
func (r *RedisService) AddSong(ctx context.Context, song *models.Song) error {
	data, err := json.Marshal(song)
	if err != nil {
		return fmt.Errorf("failed to marshal song: %v", err)
	}

	key := fmt.Sprintf("song:%s", song.ID.String())
	return r.client.Set(ctx, key, data, r.ttl).Err()
}

// GetSong retrieves a song from Redis by ID.
func (r *RedisService) GetSong(ctx context.Context, songID string) (*models.Song, error) {
	key := fmt.Sprintf("song:%s", songID)
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get song from redis: %v", err)
	}

	var song models.Song
	if err := json.Unmarshal([]byte(data), &song); err != nil {
		return nil, fmt.Errorf("failed to unmarshal song: %v", err)
	}

	return &song, nil
}

func (r *RedisService) DeleteSong(ctx context.Context, songID string) error {
	key := fmt.Sprintf("song:%s", songID)
	deleted, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to delete song from redis: %s", err.Error())
	}
	if deleted == 0 {
		return fmt.Errorf("song with ID %s not found in cache", songID)
	}
	return nil
}
