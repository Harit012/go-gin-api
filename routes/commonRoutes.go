package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes is the central route registrar
func RegisterRoutes(router *gin.Engine) {
    // Create an API group
    api := router.Group("/api")

    // Register individual route groups
    RegisterUserRoutes(api)
    RegisterDemoRoutes(api)
}