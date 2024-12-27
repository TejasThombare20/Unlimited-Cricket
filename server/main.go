package main

import (
	"TejasThombare20/fampay/cache"
	"TejasThombare20/fampay/client"
	"TejasThombare20/fampay/config"
	"TejasThombare20/fampay/controller"
	"TejasThombare20/fampay/middleware"
	"TejasThombare20/fampay/repository"
	"TejasThombare20/fampay/route"
	"TejasThombare20/fampay/service"
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func main() {

	enviroment := os.Getenv("GO_ENVIROMENT")

	//we don't need to explicitly load env file in production environment
	if enviroment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}
	//loading all required  configs from environment variable
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize YouTube client
	ytClient, err := client.NewYoutubeClient(cfg.YoutubeAPIKeys)
	if err != nil {
		log.Fatalf("Failed to create YouTube client: %v", err)
	}

	// Initalize the Postgres DB
	db, err := config.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Error while connecting with database : ", err)
	}
	defer db.Close()

	repo := repository.NewVideoRepository(db)

	// Initialize the Redis server
	videoCache, err := cache.NewVideoCache(cfg.RedisURL, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Initialize YouTube service and controller with dependencies.
	ytService := service.NewYoutubeService(ytClient, repo, videoCache)
	ytController := controller.NewVideoController(ytService)

	// Start background worker
	ctx := context.Background()
	ytService.StartBackgroundWorker(ctx, cfg)

	router := gin.Default()
	router.Use(cors.Default())

	// Middleware for handling  Rate limiting
	rateLimiter := middleware.NewIPRateLimiter(rate.Limit(cfg.RPS), cfg.BurstTime)
	router.Use(rateLimiter.RateLimit())

	route.SetupRoutes(router, ytController)
	router.Run(":8080")

}
