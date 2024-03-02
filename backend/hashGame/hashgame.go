package hashgame

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

func HashGame(positionSelected [2]int) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info(strings.Trim(strings.Replace(fmt.Sprint(positionSelected), " ", ",", -1), "[]"))

}
