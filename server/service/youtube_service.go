package service

import (
	"TejasThombare20/fampay/cache"
	"TejasThombare20/fampay/client"
	"TejasThombare20/fampay/config"
	"TejasThombare20/fampay/model"
	"TejasThombare20/fampay/repository"
	"context"
	"log"
	"time"
)

type YoutubeService struct {
	client     *client.YoutubeClient
	repository *repository.VideoRepository
	cache      *cache.VideoCache
}

func NewYoutubeService(client *client.YoutubeClient, repo *repository.VideoRepository, cache *cache.VideoCache) *YoutubeService {
	return &YoutubeService{
		client:     client,
		repository: repo,
		cache:      cache,
	}
}

func (s *YoutubeService) FetchAndStoreVideos(ctx context.Context, query string, cfg *config.Config) error {
	videos, err := s.client.SearchVideos(query, cfg)

	if err != nil {
		return err
	}

	log.Println("videos", videos)

	for _, video := range videos {
		// Convert to our model and store
		publishedAt, _ := time.Parse(time.RFC3339, video.Snippet.PublishedAt)

		videoModel := &model.Video{
			ID:           video.Id.VideoId,
			Title:        video.Snippet.Title,
			Description:  video.Snippet.Description,
			PublishedAt:  publishedAt,
			ThumbnailURL: video.Snippet.Thumbnails.Default.Url,
			CreatedAt:    time.Now(),
		}

		if err := s.repository.Create(videoModel); err != nil {
			log.Printf("Error storing video %s: %v", videoModel.ID, err)
			continue
		}

		// Invalidate cache for first page since we have new data
		// This will force a refresh on the next request
		s.cache.InvalidateFirstPage(ctx)

	}

	return nil
}

func (s *YoutubeService) GetData(ctx context.Context, page int, pageSize int) ([]model.Video, error) {

	videos, found := s.cache.GetVideos(ctx, page, pageSize)
	if found {
		log.Printf("Cache hit for page %d", page)
		return videos, nil
	}
	log.Printf("Cache miss for page %d, fetching from database", page)

	videos, err := s.repository.List(ctx, page, pageSize)

	if err != nil {
		return nil, err
	}

	if err := s.cache.SetVideos(ctx, page, pageSize, videos); err != nil {
		log.Printf("Failed to cache videos: %v", err)
	}

	return videos, nil

}
