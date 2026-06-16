package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	VideoID   uint      `json:"video_id" gorm:"not null;index"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	ParentID  *uint     `json:"parent_id" gorm:"index"` // for replies
	LikesCount int64    `json:"likes_count" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
	UpdatedAt time.Time `json:"updated_at"`

	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Replies   []Comment `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
}

type CreateCommentRequest struct {
	VideoID  uint   `json:"video_id" binding:"required"`
	Content  string `json:"content" binding:"required,min=1,max=1000"`
	ParentID *uint  `json:"parent_id"`
}
