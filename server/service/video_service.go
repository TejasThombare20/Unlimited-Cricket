package service

import (
	"TejasThombare20/fampay/config"
	"context"
	"log"
	"time"
)

var SearchQuery = "cricket"

func (s *YoutubeService) StartBackgroundWorker(ctx context.Context, cfg *config.Config) {

	log.Println("background data fectcing....")
	ticker := time.NewTicker(time.Duration(cfg.FetchTime) * time.Minute)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				if err := s.FetchAndStoreVideos(ctx, SearchQuery, cfg); err != nil {
					log.Printf("Error fetching videos: %v", err)
				}
			}
		}
	}()
}
