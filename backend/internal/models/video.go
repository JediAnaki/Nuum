package models

import (
	"time"
)

type VideoStatus string

const (
	StatusUploading  VideoStatus = "uploading"
	StatusProcessing VideoStatus = "processing"
	StatusReady      VideoStatus = "ready"
	StatusFailed     VideoStatus = "failed"
)

type Video struct {
	ID          uint        `json:"id" gorm:"primaryKey"`
	UserID      uint        `json:"user_id" gorm:"not null;index"`
	Title       string      `json:"title" gorm:"not null"`
	Description string      `json:"description" gorm:"type:text"`
	Thumbnail   string      `json:"thumbnail"`
	Duration    int         `json:"duration"` // seconds
	Status      VideoStatus `json:"status" gorm:"default:'uploading'"`
	ViewsCount  int64       `json:"views_count" gorm:"default:0;index"`
	LikesCount  int64       `json:"likes_count" gorm:"default:0"`

	// Storage paths
	OriginalPath string `json:"-" gorm:"not null"`
	ProcessedVideos []ProcessedVideo `json:"processed_videos,omitempty" gorm:"foreignKey:VideoID"`

	// Metadata
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Codec       string `json:"codec"`
	Bitrate     int    `json:"bitrate"`

	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`

	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments  []Comment `json:"comments,omitempty" gorm:"foreignKey:VideoID"`
}

type ProcessedVideo struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	VideoID    uint   `json:"video_id" gorm:"not null;index"`
	Quality    string `json:"quality"` // 360p, 480p, 720p, 1080p
	Path       string `json:"path" gorm:"not null"`
	Size       int64  `json:"size"` // bytes
	Bitrate    int    `json:"bitrate"`
	CreatedAt  time.Time `json:"created_at"`
}

type UploadVideoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description"`
}

type VideoListResponse struct {
	Videos     []Video `json:"videos"`
	TotalCount int64   `json:"total_count"`
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
}
