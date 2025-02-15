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

//Gets userid from coookie. returns -1 if user is not logged in
func GetUserId(c *gin.Context) int {
    userId, exists := c.Get("userId") 
		if(!exists){
			return -1
		} else{
			userIdStr, _ := userId.(string)		
		res, _ := strconv.Atoi(userIdStr)
		return res
		}
}

func ConvertCommentToDTO(comment models.Comment, userID int) dto.CommentDTO{
	var commentDto dto.CommentDTO
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
    if(comment.Replies == nil){
        commentDto.Replies = []dto.CommentDTO{}
    }
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

//check if is voted and number of votes
func VotesAggregation(posts []models.Post, userID int) []dto.PostDTO{
    res := make([]dto.PostDTO, len(posts))
     transformPostToDTO := func(post models.Post) dto.PostDTO {
        // Calculate upvotes and downvotes
        upvotes := 0
        downvotes := 0
        isUpvoted := false
        isDownvoted := false
    
        for _, vote := range post.Votes {
            if vote.Upvote {
                upvotes++
                if vote.UserID == userID  {
                    isUpvoted = true
                    log.Print(true)
                }
            }
            if vote.Downvote {
                downvotes++
                if vote.UserID == userID {
                    isDownvoted = true
                }
            }
        }
    
        // Transform Post to PostDTO
        return dto.PostDTO{
            ID:          post.ID,
            UserID:      post.UserID,
            Username:    post.User.Username, 
            Title:       post.Title,
            Content:     post.Content,
            CreatedAt:   post.CreatedAt,
            PrimaryTag:  post.PrimaryTag,
            IsUpvoted:   isUpvoted,
            IsDownvoted: isDownvoted,
            Upvotes:     upvotes,
            Downvotes:   downvotes,
        }
    }

    for i, post := range posts {
		res[i] = transformPostToDTO(post)
	}

    return res;

}

