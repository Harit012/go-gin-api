package main

import (
    "github.com/gin-gonic/gin"
    "go-gin-api/routes" // adjust if your module name is different
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Register routes from other packages
    routes.RegisterRoutes(r)

    // Register 404 handler
    r.NoRoute(func(c *gin.Context) {
        c.JSON(404, gin.H{
            "message": "No route available...",
        })
    })

    return r
}

func main() {
    r := SetupRouter()
    r.Run(":8080")
}

// to check Multithreading behavior.
// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	router := gin.Default()

// 	router.GET("/slow", func(c *gin.Context) {
// 		time.Sleep(5 * time.Second) // Simulate a long task
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Done",
// 		})
// 	})

// 	router.Run(":8080")
// }