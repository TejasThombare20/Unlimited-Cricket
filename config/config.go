package config

import (
	"fmt"
	"os"
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
}

func Load() (*Config, error) {
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
		DatabaseURL: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
	}, nil
}
