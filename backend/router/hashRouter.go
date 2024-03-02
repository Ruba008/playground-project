package router

import (
	"net/http"
	hashgame "playground/hashGame"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger, _ = zap.NewDevelopment()

type HashParams struct {
	Player   string `json:"player"`
	Position [2]int `json:"position"`
}

func hashRouter(router *gin.Engine) {

	hashRoutes := router.Group("/v1/hash")
	{
		hashRoutes.POST("/position", getPosition)
	}

	router.POST("/position", func(ctx *gin.Context) {
		var HashPosition HashParams
		err := ctx.BindJSON(&HashPosition)
		if err != nil {
			logger.Info(err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": HashPosition.Position,
		})

	})

}

func getPosition(ctx *gin.Context) {
	defer logger.Sync()

	// Body
	var HashPosition HashParams
	err := ctx.BindJSON(&HashPosition)
	if err != nil {
		logger.Info(err.Error())
	}

	// Response return
	ctx.JSON(http.StatusOK, gin.H{
		"message": HashPosition.Position,
	})

	hashgame.HashGame(HashPosition.Position)

}
