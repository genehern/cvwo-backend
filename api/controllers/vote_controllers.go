package controllers

import (
	"cvwo-backend/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePostVote(c *gin.Context) {
    var vote models.PostVote
    if err := c.ShouldBindJSON(&vote); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    userId, _:= c.Get("userId")
    vote.UserID = int(userId.(float64))
    if err := models.CreatePostVote(&vote); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "CommentVote created successfully", "vote": vote})
}

func DeletePostVote(c *gin.Context) {
    postId, err := strconv.Atoi(c.Param("postId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    userId, err := strconv.Atoi(c.Param("userId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    if err := models.DeletePostVote(postId, userId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "CommentVote deleted successfully"})
}