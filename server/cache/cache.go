package cache

import (
	"TejasThombare20/fampay/config"
	"TejasThombare20/fampay/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type VideoCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewVideoCache(redisURL string, cfg *config.Config) (*VideoCache, error) {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	return &VideoCache{
		client: redis.NewClient(opt),
		ttl:    time.Minute * time.Duration(cfg.FetchTime), // Cache for fetchtime(env variable) minutes
	}, nil
}

// GetVideos tries to get videos from cache
func (c *VideoCache) GetVideos(ctx context.Context, page, pageSize int) ([]model.Video, bool) {
	key := c.getCacheKey(page, pageSize)
	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, false
	}

	var videos []model.Video
	if err := json.Unmarshal(data, &videos); err != nil {
		return nil, false
	}

	return videos, true
}

// set the video cache into redis
func (c *VideoCache) SetVideos(ctx context.Context, page, pageSize int, videos []model.Video) error {
	key := c.getCacheKey(page, pageSize)
	data, err := json.Marshal(videos)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, data, c.ttl).Err()
}

func (c *VideoCache) InvalidateFirstPage(ctx context.Context) {
	// Remove cache for common page sizes
	pageSizes := []int{10, 20, 50}
	for _, size := range pageSizes {
		key := c.getCacheKey(1, size)
		c.client.Del(ctx, key)
	}
}

// cache key structure
func (c *VideoCache) getCacheKey(page, pageSize int) string {
	return fmt.Sprintf("videos:page:%d:size:%d", page, pageSize)
}
