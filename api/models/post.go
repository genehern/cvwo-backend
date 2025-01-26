package models

import (
	"time"
)
type Post struct {
	ID       int   `json:"id" gorm:"primary_key"`
	UserID int `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	PrimaryTag string `json:"primary_tag"`
}

func CreatePost( newPost *Post) error {
	if err := DB.Create(&newPost).Error; err != nil {
		return err
	}
	return nil
}

func DeletePost( postId int) error{
	if err := DB.Delete(&Post{}, postId).Error; err!= nil{
		return err
	}
	return nil
}

func GetPost( pageNum int, limitNum int) ([]Post, error) {
	var posts []Post

	if err := DB.Offset((pageNum - 1) * limitNum).Limit(limitNum).Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

