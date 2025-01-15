package models

import (
	"time"
)

type Comment struct {
    ID             uint      `json:"id" gorm:"primaryKey"`
    PostID         uint      `json:"post_id"`
    UserID         uint      `json:"user_id"`
    ParentCommentID *uint    `json:"parent_comment_id"` // Pointer to allow null values
    Content        string    `json:"content"`
    CreatedAt      time.Time `json:"created_at"`
    User           User      `json:"user" gorm:"foreignKey:UserID"`
    Replies        []Comment `json:"replies" gorm:"foreignKey:ParentCommentID"`
}