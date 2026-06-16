package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"video-platform/worker/internal/processor"
	"video-platform/worker/internal/queue"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Starting Video Processing Worker...")

	// Load configuration from environment
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "video_platform")

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")

	// Connect to PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	// Initialize video processor
	outputDir := getEnv("OUTPUT_DIR", "./processed_videos")
	videoProcessor := processor.NewVideoProcessor(outputDir)

	// Initialize queue consumer
	redisQueue := queue.NewRedisQueue(
		redisClient,
		"video:processing",
		"video-workers",
		"worker-1",
	)

	if err := redisQueue.Initialize(ctx); err != nil {
		log.Fatalf("Failed to initialize queue: %v", err)
	}

	// Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received shutdown signal, stopping worker...")
		cancel()
	}()

	// Start consuming tasks
	log.Println("Worker is ready to process videos...")
	err = redisQueue.Consume(ctx, func(task queue.VideoProcessingTask) error {
		return processVideo(db, videoProcessor, task)
	})

	if err != nil && err != context.Canceled {
		log.Fatalf("Worker error: %v", err)
	}

	log.Println("Worker stopped gracefully")
}

func processVideo(db *gorm.DB, processor *processor.VideoProcessor, task queue.VideoProcessingTask) error {
	log.Printf("Starting processing for video ID: %d", task.VideoID)

	// Get video from database
	var video struct {
		ID           uint
		OriginalPath string
		UserID       uint
	}

	if err := db.Table("videos").Where("id = ?", task.VideoID).First(&video).Error; err != nil {
		return fmt.Errorf("failed to fetch video: %w", err)
	}

	// Update status to processing
	if err := db.Table("videos").Where("id = ?", task.VideoID).Update("status", "processing").Error; err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	// Process video
	results, err := processor.Process(processor.ProcessingOptions{
		InputPath: video.OriginalPath,
		OutputDir: "./processed_videos",
		VideoID:   video.ID,
	})

	if err != nil {
		// Update status to failed
		db.Table("videos").Where("id = ?", task.VideoID).Update("status", "failed")
		return fmt.Errorf("failed to process video: %w", err)
	}

	// Save processed videos to database
	for _, result := range results {
		processedVideo := map[string]interface{}{
			"video_id": video.ID,
			"quality":  result.Quality,
			"path":     result.Path,
			"size":     result.Size,
			"bitrate":  result.Bitrate,
		}

		if err := db.Table("processed_videos").Create(processedVideo).Error; err != nil {
			log.Printf("Failed to save processed video %s: %v", result.Quality, err)
		}
	}

	// Generate thumbnail
	thumbnailPath := fmt.Sprintf("./thumbnails/video_%d.jpg", video.ID)
	if err := processor.GenerateThumbnail(video.OriginalPath, thumbnailPath, ""); err != nil {
		log.Printf("Failed to generate thumbnail: %v", err)
	} else {
		db.Table("videos").Where("id = ?", task.VideoID).Update("thumbnail", thumbnailPath)
	}

	// Update status to ready
	if err := db.Table("videos").Where("id = ?", task.VideoID).Update("status", "ready").Error; err != nil {
		return fmt.Errorf("failed to update status to ready: %w", err)
	}

	log.Printf("Successfully processed video ID: %d with %d quality variants", task.VideoID, len(results))
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
