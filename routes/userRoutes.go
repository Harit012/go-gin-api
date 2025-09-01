package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
    user := rg.Group("/user")

    user.GET("/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{"user_id": id})
    })

    user.POST("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User created"})
    })
}