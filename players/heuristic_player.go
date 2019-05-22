package players

import (
	"math"

	"github.com/carbon-12/go-games/base"
)

type HeuristicPlayer struct {
	gameState base.GameState
	heuristic base.Heuristic
}

func (player *HeuristicPlayer) GetNextMove() interface{} {
	var bestMove interface{}
	var bestMoveValue float64 = -math.MaxFloat64
	for _, move := range *player.gameState.GetPossibleMoves() {
		gameState := player.gameState.Clone().(base.GameState)
		gameState.MakeMove(move)

		heuristic := player.heuristic.Heuristic(gameState)
		if heuristic > bestMoveValue {
			bestMove = move
			bestMoveValue = heuristic
		}
	}
	return bestMove
}

func NewHeuristicPlayer(gameState base.GameState, heuristic base.Heuristic) *HeuristicPlayer {
	return &HeuristicPlayer{gameState: gameState, heuristic: heuristic}
}
