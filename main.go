package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/db"
	"go-gin-api/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to DB
	db.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

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