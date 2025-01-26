package dto

import (
	"time"
)

type PostDTO struct{
	ID       int   `json:"id" gorm:"primary_key"`
	UserID int `json:"user_id"`
	Username string `json:"username"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	PrimaryTag string `json:"primary_tag"`
	IsUpvoted bool `json:"is_upvoted"`
	IsDownvoted bool `json:"is_downvoted"`
	Comments []CommentDTO `json:"comments"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`
}

type CommentDTO struct{
	ID             int      `json:"id" gorm:"primaryKey"`
    PostID         int      `json:"post_id"`
    UserID         int      `json:"user_id"`
    ParentCommentID *int    `json:"parent_comment_id"` // Pointer to allow null values
    Content        string    `json:"content"`
    CreatedAt      time.Time `json:"created_at"`
    Username         string      `json:"user"`
    Replies        []CommentDTO `json:"replies" gorm:"foreignKey:ParentCommentID"`
	IsUpvoted bool `json:"is_upvoted"`
	IsDownvoted bool `json:"is_downvoted"`
	Upvotes int `json:"upvotes"`
	Downvotes int `json:"downvotes"`
}