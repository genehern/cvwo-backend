package routes

import (
	"cvwo-backend/api/controllers"
	"cvwo-backend/api/middlewares"

	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(router *gin.Engine) {
	protected := router.Group("/protected")
	{
		protected.Use(middlewares.AuthMiddleware())

		PostRoutes(protected)
		CommentRoutes(protected)
	}
}

func PublicRoutes(router *gin.Engine) {
	public := router.Group("/public")


	UserRoutes(public)
	PostRoutes(public)
	CommentRoutes(public)
}

func PostRoutes(group *gin.RouterGroup) {
	postsGroup := group.Group("/posts")
	{
		postsGroup.POST("/createPost", controllers.CreatePost)
		postsGroup.GET("/", controllers.GetPosts)
	}
}

func CommentRoutes(group *gin.RouterGroup) {
	commentsGroup := group.Group("/comments")
	{
		commentsGroup.POST("/createComment", controllers.CreateComment)
		commentsGroup.GET("/", controllers.GetComments)
	}
}

func UserRoutes(group *gin.RouterGroup) {
	usersGroup := group.Group("/user")
	{
		usersGroup.POST("/login", controllers.Login)
		usersGroup.POST("/createUser", controllers.CreateUser)
	}
}

