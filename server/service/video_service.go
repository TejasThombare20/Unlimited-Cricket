package service

import (
	"context"
	"log"
	"time"
)

var SearchQuery = "cricket"

func (s *YoutubeService) StartBackgroundWorker(ctx context.Context) {

	log.Println("background data fectcing....")
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				if err := s.FetchAndStoreVideos(SearchQuery); err != nil {
					log.Printf("Error fetching videos: %v", err)
				}
			}
		}
	}()
}
