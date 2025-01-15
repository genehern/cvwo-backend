package models

import (
	"time"

	"gorm.io/gorm"
)
type Post struct {
	ID       int   `json:"id" gorm:"primary_key"`
	UserID int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Upvote int `json:"upvote"`
	Downvote int `json:"downvote"`
	PrimaryTag string `json:"primary_tag"`
	SecondaryTag string `json:"secondary_tag"`
}

func CreatePost(db *gorm.DB, newPost *Post) error {
	if err := db.Create(&newPost).Error; err != nil {
		return err
	}
	return nil
}

func DeletePost(db *gorm.DB, postId int) error{
	if err := db.Delete(&Post{}, postId).Error; err!= nil{
		return err
	}
	return nil
}

func GetPostByTags(db *gorm.DB, pageNum int, limitNum int) ([]Post, error) {
	var posts []Post
	query := db.Debug() 

	if err := query.Offset((pageNum - 1) * limitNum).Limit(limitNum).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

