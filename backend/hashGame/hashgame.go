package hashgame

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

var player1 = make([][2]int, 0)
var player2 = make([][2]int, 0)

func HashGame(player bool, position [2]int) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var player2 [][2]int

	if player {
		player1 = append(player1, position)
	} else {
		player2 = append(player2, position)
	}

	fmt.Println(player1)
	for index, item := range player1 {
		logger.Info("Pair values", zap.Int("First", item[0]), zap.Int("Second", item[1]))
		logger.Info(strconv.Itoa(index))
	}
}
