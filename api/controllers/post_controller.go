package controllers

import (
	"cvwo-backend/api/models"
	"cvwo-backend/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var newPost models.Post
	
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	userId, _:= c.Get("userId")

	newPost.UserID = int(userId.(float64))
	if err := models.CreatePost(&newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "user": newPost})
}

func GetPosts(c *gin.Context){	
	
		pageNum, limitNum := utils.GetPaginationParam(c)
		posts, err := models.GetPost(pageNum, limitNum)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching posts"})
			return
		}

		userId, exist := c.Get("userId")
		if !exist{
			userId = 0;
		}
		value, _ := userId.(int)

		res := utils.VotesAggregation(posts, value)
		c.JSON(200, res)
}



