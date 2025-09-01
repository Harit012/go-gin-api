package routes

import (
	"go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
    userRoutes := rg.Group("/user")

    userRoutes.POST("/", controllers.CreateUser);

	userRoutes.GET("/", controllers.GetUsers);
    // Add more user routes here
}
