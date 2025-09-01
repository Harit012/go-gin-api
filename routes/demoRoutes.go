package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterDemoRoutes(rg *gin.RouterGroup) {
    demo := rg.Group("/demo")

    demo.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong from demo"})
    })
}