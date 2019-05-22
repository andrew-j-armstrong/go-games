package players

import (
	"fmt"

	"github.com/carbon-12/go-games/base"
)

type HumanPlayer struct {
	game      base.Game
	gameState base.GameState
}

func (player *HumanPlayer) GetNextMove() interface{} {
	player.gameState.Print()

	var move base.Move
	for {
		print("Move: ")

		var moveString string
		_, err := fmt.Scan(&moveString)
		if err != nil {
			fmt.Println(err)
			continue
		}

		move, err = player.game.ParseMove(moveString)

		if err != nil {
			fmt.Printf("Invalid Move: %s\n", err)
			continue
		}

		if !player.gameState.IsValidMove(move) {
			fmt.Println("Invalid Move!")
			continue
		}

		break
	}

	return move
}

func NewHumanPlayer(game base.Game, gameState base.GameState) *HumanPlayer {
	return &HumanPlayer{game: game, gameState: gameState}
}
