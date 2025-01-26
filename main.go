package main

import (
	"cvwo-backend/api/controllers"
	"cvwo-backend/api/middlewares"
	"cvwo-backend/api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // React app URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	models.ConnectDatabase()

	protected := router.Group("/protected")
	{
		protected.Use(middlewares.AuthMiddleware()) // Apply middleware
		
		postGroup := protected.Group("/posts")
		{
			postGroup.POST("", controllers.CreatePost)
		}

		voteGroup := protected.Group("/votes")
		{
			voteGroup.POST("/postVote", controllers.CreatePostVote)
			//voteGroup.POST("/CommentVote", controllers.CreateCommentVote)
			voteGroup.DELETE("/postVote", controllers.DeletePostVote)
			//voteGroup.DELETE("/CommentVote", controllers.DeleteCommentVote)
		}

		commentsGroup := protected.Group("/comments")
		{
			commentsGroup.POST("", controllers.CreateComment)
		}
	}

	public := router.Group("/public")
	{
		usersGroup := public.Group("/user")
		{
			usersGroup.POST("login", controllers.Login)
			usersGroup.POST("createUser", controllers.CreateUser)
		}

		postsGroup := public.Group("/posts")
		{
			postsGroup.GET("", controllers.GetPosts)
		}

		commentsGroup := public.Group("/comments")
		{
			commentsGroup.GET("", controllers.GetComments)
		}
	}

	router.Run(":3000")
}