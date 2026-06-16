package main

import (
	"log"
	"video-platform/backend/internal/config"
	"video-platform/backend/internal/database"
	"video-platform/backend/internal/handlers"
	"video-platform/backend/internal/middleware"
	"video-platform/backend/internal/services"
	"video-platform/backend/pkg/logger"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize logger
	logger.Init()
	logger.Info.Println("Starting Video Platform API...")

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewPostgresDB(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize Redis
	redisClient, err := database.NewRedisClient(&cfg.Redis)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// Initialize services
	authService := services.NewAuthService(db)
	videoService := services.NewVideoService(db, redisClient)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, cfg)
	videoHandler := handlers.NewVideoHandler(videoService)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		BodyLimit: 2 * 1024 * 1024 * 1024, // 2GB for video uploads
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(fiberLogger.New())
	app.Use(middleware.CORS())

	// Initialize auth middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWT.Secret)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"service": "video-platform-api",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Get("/me", authMiddleware.Protected(), authHandler.GetMe)

	// Video routes
	videos := api.Group("/videos")
	videos.Get("/", videoHandler.GetVideosList)           // Public: list videos
	videos.Get("/:id", videoHandler.GetVideo)             // Public: get video details
	videos.Post("/", authMiddleware.Protected(), videoHandler.UploadVideo)     // Protected: upload video
	videos.Delete("/:id", authMiddleware.Protected(), videoHandler.DeleteVideo) // Protected: delete video

	// Start server
	port := cfg.Server.Port
	logger.Info.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
