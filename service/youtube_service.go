package service

import (
	"TejasThombare20/fampay/client"
	"TejasThombare20/fampay/model"
	"TejasThombare20/fampay/repository"
	"context"
	"log"
	"time"
)

type YoutubeService struct {
	client     *client.YoutubeClient
	repository *repository.VideoRepository
}

func NewYoutubeService(client *client.YoutubeClient, repo *repository.VideoRepository) *YoutubeService {
	return &YoutubeService{
		client:     client,
		repository: repo,
	}
}

func (s *YoutubeService) FetchAndStoreVideos(query string) error {
	videos, err := s.client.SearchVideos(query)
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

		s.repository.Create(videoModel)
	}

	return nil
}

func (s *YoutubeService) GetData(ctx context.Context, page int, pageSize int) ([]model.Video, error) {

	return s.repository.List(ctx, page, pageSize)
}
