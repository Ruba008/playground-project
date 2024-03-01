package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HashParams struct {
	Position [2]int `json:"position"`
}

func GetRouter() {
	router := gin.Default()

	router.GET("/position", func(ctx *gin.Context) {
		var hashParams HashParams
		err := ctx.BindJSON(&hashParams)
		if err != nil {
			fmt.Print("err")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": hashParams.Position,
		})

	})
	router.Run(":8081")
}
