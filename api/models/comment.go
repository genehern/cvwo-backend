package models

import (
	"time"
)

type Comment struct {
    ID             int      `json:"id" gorm:"primaryKey"`
    PostID         int      `json:"post_id"`
    UserID         int      `json:"user_id"`
    ParentCommentID *int    `json:"parent_comment_id"` // Pointer to allow null values
    Content        string    `json:"content"`
    CreatedAt      time.Time `json:"created_at"`
    User           User      `json:"user" gorm:"foreignKey:UserID"`
    Replies        []Comment `json:"replies" gorm:"foreignKey:ParentCommentID"`
    Votes          []CommentVote `gorm:"foreignKey:CommentID"`
}

func CreateComment(comment *Comment)error{
    if err := DB.Debug().Model(&Comment{}).Create(&comment).Error; err!= nil{
        return err
    }
    return nil
} 

func DeleteComment(id uint)error{
    if err := DB.Where("ID= ?", id).Delete(&Comment{}).Error; err!= nil{
        return err
    }
    return nil
} 

func GetComments( postID int, userID int, pageNum int, limitNum int) ([]Comment, error) {
	var comments []Comment

	if err := DB.Offset((pageNum - 1) * limitNum).Limit(limitNum).
    Where("post_id = ? and parent_comment_id is NULL", postID).
    Preload("Replies").
    Find(&comments).Error; err != nil {
    return nil, err
}

var loadUsersRecursively func(comments []Comment)
  loadUsersRecursively = func(comments []Comment) {
	for i := range comments {
		DB.Model(&User{}).Select("id", "username").Where("id = ?", comments[i].UserID).First(&comments[i].User)
        DB.Model(&CommentVote{}).Where("comment_id = ?", comments[i].ID).First(&comments[i].Votes)
		loadUsersRecursively(comments[i].Replies)
	}
}

    loadUsersRecursively(comments)
	return comments, nil
}





