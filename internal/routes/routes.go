package routes

import (
	"github.com/adityjoshi/parkingSystem.git/internal/handler"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all routes for the application.
func SetupRoutes(r *gin.Engine) {
	r.POST("/register", handler.RegisterVehicle)
}
