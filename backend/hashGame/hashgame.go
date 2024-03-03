package hashgame

import (
	"fmt"

	"go.uber.org/zap"
)

var player1 = make([][2]int, 0)
var player2 = make([][2]int, 0)
var Draw bool
var PlayerWinner string

func HashGame(player bool, position [2]int) {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Saving the positions selected
	if player {
		player1 = append(player1, position)
	} else {
		player2 = append(player2, position)
	}

	// Mapping for search
	player1Mapped := map[[2]int]bool{}
	for _, item := range player1 {
		player1Mapped[item] = true
	}
	player2Mapped := map[[2]int]bool{}
	for _, item := range player2 {
		player2Mapped[item] = true
	}

	// Verifing if there's a winner
	if verifyWinner(player1Mapped) {
		Draw = false
		PlayerWinner = "Player 1"
		fmt.Println("Player 1 wins!")
	}
	if verifyWinner(player2Mapped) {
		Draw = false
		PlayerWinner = "Player 2"
		fmt.Println("Player 2 wins!")
	}
	if len(player1) == 5 && len(player2) == 4 && !verifyWinner(player1Mapped) {
		Draw = true
		PlayerWinner = "null"
		fmt.Println("Draw!")
	}

}

func verifyWinner(p map[[2]int]bool) bool {

	// Conditions for win the hash game !
	if (p[[2]int{1, 1}] && p[[2]int{1, 2}] && p[[2]int{1, 3}]) ||
		(p[[2]int{1, 1}] && p[[2]int{2, 2}] && p[[2]int{3, 3}]) ||
		(p[[2]int{1, 1}] && p[[2]int{2, 1}] && p[[2]int{3, 1}]) ||
		(p[[2]int{3, 1}] && p[[2]int{2, 2}] && p[[2]int{1, 3}]) ||
		(p[[2]int{3, 1}] && p[[2]int{3, 2}] && p[[2]int{3, 3}]) ||
		(p[[2]int{1, 3}] && p[[2]int{2, 3}] && p[[2]int{3, 3}]) ||
		(p[[2]int{2, 1}] && p[[2]int{2, 2}] && p[[2]int{2, 3}]) ||
		(p[[2]int{1, 2}] && p[[2]int{2, 2}] && p[[2]int{3, 2}]) {
		return true
	}

	return false

}
