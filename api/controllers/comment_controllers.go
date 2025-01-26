package controllers

import (
	"cvwo-backend/api/dto"
	"cvwo-backend/api/models"
	"cvwo-backend/api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var newComment models.Comment
	
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	userId, _ := c.Get("userId")
	newComment.UserID = int(userId.(float64))
	log.Print(newComment.Content)
	if err := models.CreateComment(&newComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "user": newComment})
}

func GetComments(c *gin.Context){	
		pageNum, limitNum := utils.GetPaginationParam(c)
		postId, _ := strconv.Atoi(c.DefaultQuery("postId", ""))
		userId, _ := strconv.Atoi(c.DefaultQuery("userId", ""))
		comments, err := models.GetComments(postId, userId, pageNum, limitNum)
		var dtoComment []dto.CommentDTO
		
		for _, comment := range comments{
			dtoComment = append(dtoComment, utils.ConvertCommentToDTO(comment, userId))
		}
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching comments"})
			return
		}
	
		c.JSON(200, dtoComment)
}


func A(c *gin.Context){	
	c.JSON(200,nil)
}
