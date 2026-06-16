package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"video-platform/backend/internal/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type VideoService struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewVideoService(db *gorm.DB, redis *redis.Client) *VideoService {
	return &VideoService{
		db:    db,
		redis: redis,
	}
}

func (s *VideoService) CreateVideo(userID uint, req *models.UploadVideoRequest, originalPath string) (*models.Video, error) {
	video := &models.Video{
		UserID:       userID,
		Title:        req.Title,
		Description:  req.Description,
		OriginalPath: originalPath,
		Status:       models.StatusUploading,
	}

	if err := s.db.Create(video).Error; err != nil {
		return nil, err
	}

	return video, nil
}

func (s *VideoService) GetVideoByID(videoID uint) (*models.Video, error) {
	var video models.Video
	if err := s.db.Preload("User").Preload("ProcessedVideos").First(&video, videoID).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (s *VideoService) GetVideosList(page, pageSize int, userID *uint) (*models.VideoListResponse, error) {
	var videos []models.Video
	var totalCount int64

	query := s.db.Model(&models.Video{}).Where("status = ?", models.StatusReady)

	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	// Count total
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := query.
		Preload("User").
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&videos).Error; err != nil {
		return nil, err
	}

	return &models.VideoListResponse{
		Videos:     videos,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

func (s *VideoService) UpdateVideoStatus(videoID uint, status models.VideoStatus) error {
	return s.db.Model(&models.Video{}).Where("id = ?", videoID).Update("status", status).Error
}

func (s *VideoService) IncrementViews(videoID uint) error {
	ctx := context.Background()

	// Use Redis for real-time view counting
	key := fmt.Sprintf("video:views:%d", videoID)

	// Increment in Redis
	count, err := s.redis.Incr(ctx, key).Result()
	if err != nil {
		return err
	}

	// Sync to DB every 100 views or every 5 minutes (handled by a background job)
	if count%100 == 0 {
		return s.db.Model(&models.Video{}).Where("id = ?", videoID).Update("views_count", count).Error
	}

	return nil
}

func (s *VideoService) QueueVideoForProcessing(videoID uint) error {
	ctx := context.Background()

	// Add to Redis Stream for worker to pick up
	message := map[string]interface{}{
		"video_id":   videoID,
		"queued_at":  time.Now().Unix(),
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return s.redis.XAdd(ctx, &redis.XAddArgs{
		Stream: "video:processing",
		Values: map[string]interface{}{
			"data": string(data),
		},
	}).Err()
}

func (s *VideoService) DeleteVideo(videoID, userID uint) error {
	// Verify ownership
	var video models.Video
	if err := s.db.First(&video, videoID).Error; err != nil {
		return err
	}

	if video.UserID != userID {
		return fmt.Errorf("unauthorized")
	}

	// Delete from database (cascade will handle related records)
	return s.db.Delete(&video).Error
}
