package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	YoutubeAPIKeys struct {
		Key1 string
		Key2 string
		Key3 string
		Key4 string
		Key5 string
	}
	SearchQuery string
	DatabaseURL string
	RedisURL    string
	FetchTime   int
	RPS         int
	BurstTime   int
}

func Load() (*Config, error) {

	fetchtime, err := strconv.Atoi(os.Getenv("FETCH_TIME"))

	if err != nil {
		log.Printf("Invalid FETCH_TIME: %v", err)
		fetchtime = 5
	}

	fetchRPS, err := strconv.Atoi(os.Getenv("RPS"))

	if err != nil {
		log.Printf("Invalid FETCH_TIME: %v", err)
		fetchRPS = 5
	}

	fetchBurstTime, err := strconv.Atoi(os.Getenv("BURST_TIME"))

	if err != nil {
		log.Printf("Invalid FETCH_TIME: %v", err)
		fetchBurstTime = 10
	}

	return &Config{
		YoutubeAPIKeys: struct {
			Key1 string
			Key2 string
			Key3 string
			Key4 string
			Key5 string
		}{
			Key1: os.Getenv("YOUTUBE_API_KEY_1"),
			Key2: os.Getenv("YOUTUBE_API_KEY_2"),
			Key3: os.Getenv("YOUTUBE_API_KEY_3"),
			Key4: os.Getenv("YOUTUBE_API_KEY_4"),
			Key5: os.Getenv("YOUTUBE_API_KEY_5"),
		},
		SearchQuery: "cricket",
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
		FetchTime:   fetchtime,
		RPS:         fetchRPS,
		BurstTime:   fetchBurstTime,
	}, nil
}
