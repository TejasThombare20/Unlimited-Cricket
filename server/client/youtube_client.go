package client

import (
	"TejasThombare20/fampay/config"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeClient struct {
	service *youtube.Service
	keys    struct {
		Key1 string
		Key2 string
		Key3 string
		Key4 string
		Key5 string
	}
	currentKey string
}

func NewYoutubeClient(keys struct {
	Key1 string
	Key2 string
	Key3 string
	Key4 string
	Key5 string
}) (*YoutubeClient, error) {
	// 1. Collect the keys into a slice
	allKeys := []string{keys.Key1, keys.Key2, keys.Key3, keys.Key4, keys.Key5}

	// 2. Find the first non-empty key
	var firstKey string
	for _, k := range allKeys {
		if k != "" {
			firstKey = k
			break
		}
	}

	// 3. If no key is found, throw an error
	if firstKey == "" {
		return nil, errors.New("no API keys provided")
	}

	// 4. Create and initialize the YoutubeClient
	client := &YoutubeClient{
		keys:       keys,
		currentKey: firstKey,
	}

	// 5. Instantiate a service using the first valid key
	service, err := client.createService()
	if err != nil {
		return nil, err
	}
	client.service = service

	return client, nil
}

func (c *YoutubeClient) createService() (*youtube.Service, error) {
	ctx := context.Background()
	return youtube.NewService(ctx, option.WithAPIKey(c.currentKey))
}

func (c *YoutubeClient) rotateKey() error {

	// 1. Put all potential API keys into a slice
	keysList := []string{
		c.keys.Key1,
		c.keys.Key2,
		c.keys.Key3,
		c.keys.Key4,
		c.keys.Key5,
	}

	// 2. Find the index of the current key in the slice
	currentIndex := -1
	for i, k := range keysList {
		if k == c.currentKey {
			currentIndex = i
			break
		}
	}

	// If the current key is not found, start with index 0
	if currentIndex == -1 {
		currentIndex = 0
	}

	// 3. Round-robin: move to the next key
	startIndex := (currentIndex + 1) % len(keysList)
	i := startIndex

	for {
		if keysList[i] != "" {
			// 4. Found a valid key (not empty)
			c.currentKey = keysList[i]
			break
		}

		// Move to the next index in a circular fashion
		i = (i + 1) % len(keysList)

		// If we've looped all the way back to startIndex,
		// it means no key is left
		if i == startIndex {
			return errors.New("no more API keys available")
		}
	}

	// 5. Create a new service with the newly selected key
	service, err := c.createService()
	if err != nil {
		return err
	}
	c.service = service

	return nil
}

func (c *YoutubeClient) SearchVideos(query string, cfg *config.Config) ([]*youtube.SearchResult, error) {
	log.Println("api key", c.currentKey)
	call := c.service.Search.List([]string{"id", "snippet"}).
		Q(query).
		Type("video").
		Order("date").
		PublishedAfter(time.Now().Add(-time.Duration(cfg.FetchTime) * time.Minute).Format(time.RFC3339)).
		MaxResults(50)

	response, err := call.Do()

	println("response: ", response)
	if err != nil {
		// Check if error is due to quota exhaustion
		if c.isQuotaExceeded(err) {
			log.Printf("Quota exceeded for key: %s. Attempting to rotate key...", c.currentKey)
			// Try rotating to next key
			if rotateErr := c.rotateKey(); rotateErr != nil {
				return nil, fmt.Errorf("quota exceeded and failed to rotate key: %v", rotateErr)
			}
			// Retry the search with new key
			log.Printf("Retrying with new key: %s", c.currentKey)
			return c.SearchVideos(query, cfg)
		}
		return nil, err
	}

	return response.Items, nil
}

func (c *YoutubeClient) isQuotaExceeded(err error) bool {
	if err == nil {
		return false
	}

	// Type assert to googleapi.Error
	gerr, ok := err.(*googleapi.Error)
	if !ok {
		return false
	}

	// Check if it's a 403 error and contains quota exceeded message
	return gerr.Code == 403 && strings.Contains(gerr.Message, "quota")
}
