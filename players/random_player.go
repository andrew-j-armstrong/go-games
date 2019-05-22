package players

import (
	"math/rand"

	"github.com/andrew-j-armstrong/go-games/base"
)

type RandomPlayer struct {
	gameState base.GameState
}

func (player *RandomPlayer) GetNextMove() interface{} {
	moves := player.gameState.GetPossibleMoves()
	return (*moves)[rand.Intn(len(*moves))]
}

func NewRandomPlayer(gameState base.GameState) *RandomPlayer {
	return &RandomPlayer{gameState: gameState}
}
