package router

import (
	"net/http"
	hashgame "playground/hashGame"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger, _ = zap.NewDevelopment()

type HashParams struct {
	Player   bool   `json:"player"`
	Position [2]int `json:"position"`
}

func hashRouter(router *gin.Engine) {

	hashRoutes := router.Group("/v1/hash")
	{
		hashRoutes.POST("/position", getPosition)
		hashRoutes.GET("/getResult", getResult)
	}

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
		"player":   HashPosition.Player,
		"position": HashPosition.Position,
	})

	hashgame.HashGame(HashPosition.Player, HashPosition.Position)

}

func getResult(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"draw":         hashgame.Draw,
		"playerWinner": hashgame.PlayerWinner,
	})

}
