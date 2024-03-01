package router

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HashParams struct {
	Position [2]int `json:"position"`
}

func GetRouter() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	router.POST("/position", func(ctx *gin.Context) {
		var HashPosition HashParams
		err := ctx.BindJSON(&HashPosition)
		if err != nil {
			fmt.Print("err")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": HashPosition.Position,
		})

	})

	router.Run(":8081")
}
