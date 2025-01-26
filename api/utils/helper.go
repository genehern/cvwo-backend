package utils

import (
	"cvwo-backend/api/dto"
	"cvwo-backend/api/models"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParam(c *gin.Context)(int, int){
	pageNumStr := c.DefaultQuery("pageNum", "1")
	limitNumStr := c.DefaultQuery("limitNum", "10")
	pageNum, _ := strconv.Atoi(pageNumStr)
	limitNum, _ := strconv.Atoi(limitNumStr)
	return pageNum, limitNum
}

func ConvertCommentToDTO(comment models.Comment, userID int) dto.CommentDTO{
	var commentDto dto.CommentDTO
	log.Print(comment);
	commentDto.ID = comment.ID	
    commentDto.PostID = comment.PostID
    commentDto.UserID = comment.UserID
    commentDto.ParentCommentID = comment.ParentCommentID
    commentDto.Content = comment.Content
    commentDto.CreatedAt = comment.CreatedAt
    commentDto.Username = comment.User.Username

	var repliesDTO []dto.CommentDTO
    for _, reply := range comment.Replies {
        repliesDTO = append(repliesDTO, ConvertCommentToDTO(reply, userID)) 
    }
    commentDto.Replies = repliesDTO

    for _, vote := range comment.Votes {
        if vote.UserID == userID {
            commentDto.IsUpvoted = vote.Upvote
            commentDto.IsDownvoted = vote.Downvote
        }
        if vote.Upvote {
            commentDto.Upvotes++
        }
        if vote.Downvote {
            commentDto.Downvotes++
        }
    }


    return commentDto
}

