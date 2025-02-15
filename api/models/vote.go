package models

import (
	"log"

	"gorm.io/gorm"
)
type PostVote struct {
    ID             int      `json:"id" gorm:"primaryKey"`
    PostID         int      `json:"post_id" gorm:"foreignKey:PostID"`
    UserID         int      `json:"user_id" "`
    User           User     `json:"user" gorm:"foreignKey:UserID"`
    Upvote         bool     `json:"upvote"`
    Downvote       bool     `json:"downvote"`
}

type CommentVote struct {
    ID             int      `json:"id" gorm:"primaryKey"`
    CommentID      int      `json:"comment_id" gorm:"foreignKey:CommentID"`
    UserID         int      `json:"user_id"`
    User           User     `json:"user" gorm:"foreignKey:UserID"`
    Upvote         bool     `json:"upvote"`
    Downvote       bool     `json:"downvote"`
}

func CreatePostVote(vote *PostVote) error {
    var existingVote PostVote
    log.Print(vote)
    if err := DB.Debug().Where("post_id = ? AND user_id = ?", vote.PostID, vote.UserID).First(&existingVote).Error; err == nil {
        if err := DB.Debug().Where("post_id = ? AND user_id = ?", vote.PostID, vote.UserID).Select("Upvote", "Downvote").Updates(PostVote{Upvote:vote.Upvote, Downvote: vote.Downvote}).Error; err != nil {
            return err
        }
    } else if (err == gorm.ErrRecordNotFound){
        if err := DB.Debug().Create(&vote).Error; err != nil {
            return err
        }
    }
    return nil
}


func DeletePostVote(postID int, userID int) error {
    if err := DB.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&PostVote{}).Error; err != nil {
        return err
    }
    return nil
}

