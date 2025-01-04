package main

import (
	"cvwo-backend/api/controllers"
	"cvwo-backend/api/models"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    models.ConnectDatabase()
    protected := router.Group("/protected")
	{
		usersGroup := protected.Group("/user")
		{ 
			usersGroup.POST("createAccount", controllers.CreateUser)   
		}
	}

    public := router.Group("/public")
    {
        usersGroup := public.Group("/user")
		{ 
            usersGroup.POST("login", controllers.Login)
			usersGroup.POST("createUser", controllers.CreateUser)   
		}
    }
    router.Run(":8080")

}

