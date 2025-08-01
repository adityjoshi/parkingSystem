package main

import (
	"github.com/adityjoshi/parkingSystem.git/internal/redis"

	//"github.com/adityjoshi/parkingSystem.git/internal/routes"

	"github.com/adityjoshi/parkingSystem.git/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	redis.Init()

	r := gin.Default()
	//routes.SetupRoutes(r)

	r.Run(":8080")
}
