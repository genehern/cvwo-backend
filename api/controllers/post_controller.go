package controllers

import (
	"cvwo-backend/api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var newPost models.Post
	
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	userId, _ := c.Get("userId")

	newPost.UserID = int(userId.(float64))
	log.Print(newPost)
	if err := models.CreatePost(models.DB, &newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "user": newPost})
}

func GetPosts(c *gin.Context){	
		log.Print("re")
		pageNumStr := c.DefaultQuery("pageNum", "1")
		limitNumStr := c.DefaultQuery("limitNum", "10")
		pageNum, _ := strconv.Atoi(pageNumStr)
		limitNum, _ := strconv.Atoi(limitNumStr)
		posts, err := models.GetPostByTags(models.DB, pageNum, limitNum)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error fetching posts"})
			return
		}

		c.JSON(200, posts)
}


