package hashgame

import (
	"strconv"

	"go.uber.org/zap"
)

func HashGame(player bool, position [2]int) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info(strconv.FormatBool(player))

}
