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
        protected.Use(middlewares.AuthMiddleware())
		postGroup := protected.Group("/post")
		{ 
			postGroup.POST("createPost", controllers.CreatePost)
		}
	}
    
    public := router.Group("/public")
    {
        usersGroup := public.Group("/user")
		{ 
            usersGroup.POST("login", controllers.Login)
			usersGroup.POST("createUser", controllers.CreateUser)   
		}
        postsGroup := public.Group("/post")
        {
            postsGroup.GET("", controllers.GetPosts)
        }
    }
   router.Run(":3000")

}

