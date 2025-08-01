package main

import (
	"mall/internal/db"
	"mall/internal/redis"
	"mall/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	redis.Init()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
