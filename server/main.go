package main

import (
	"TejasThombare20/fampay/client"
	"TejasThombare20/fampay/config"
	"TejasThombare20/fampay/controller"
	"TejasThombare20/fampay/repository"
	"TejasThombare20/fampay/route"
	"TejasThombare20/fampay/service"
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize YouTube client
	ytClient, err := client.NewYoutubeClient(cfg.YoutubeAPIKeys)
	if err != nil {
		log.Fatalf("Failed to create YouTube client: %v", err)
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error while connecting with database", err)
	}
	defer db.Close()
	repo := repository.NewVideoRepository(db)

	ytService := service.NewYoutubeService(ytClient, repo)

	ytController := controller.NewVideoController(ytService)

	// Start background worker
	ctx := context.Background()
	ytService.StartBackgroundWorker(ctx)

	router := gin.Default()

	router.Use(cors.Default())

	route.SetupRoutes(router, ytController)
	router.Run(":8080")

}
