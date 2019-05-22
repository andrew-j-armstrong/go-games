package games_reversi

import (
	"github.com/carbon-12/go-games/base"
	"github.com/carbon-12/go-reversi"
)

type ReversiStandardHeuristic struct {
	heuristic *reversi.StandardHeuristic
}

func (heuristic *ReversiStandardHeuristic) Heuristic(gameState base.GameState) float64 {
	return heuristic.heuristic.Heuristic(gameState.(*ReversiGameState).GetGameState())
}

func NewReversiStandardHeuristic(playerID base.PlayerID) *ReversiStandardHeuristic {
	var reversiPlayerID reversi.PlayerID
	switch playerID {
	case base.Player1:
		reversiPlayerID = reversi.Player1
	case base.Player2:
		reversiPlayerID = reversi.Player2
	default:
		return nil
	}

	return &ReversiStandardHeuristic{reversi.NewStandardHeuristic(reversiPlayerID)}
}
