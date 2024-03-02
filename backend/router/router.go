package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() {
	// Router Initialization
	router := gin.Default()

	// CORS Configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	// Hash game routes
	hashRouter(router)

	// Router Run
	router.Run(":8081")
}
