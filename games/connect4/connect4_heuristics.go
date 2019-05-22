package games_connect4

import (
	"github.com/carbon-12/go-connect4"
	"github.com/carbon-12/go-games/base"
)

type Connect4SimpleHeuristic struct {
	heuristic *connect4.SimpleHeuristic
}

func (heuristic *Connect4SimpleHeuristic) Heuristic(gameState base.GameState) float64 {
	return heuristic.heuristic.Heuristic(gameState.(*Connect4GameState).GetGameState())
}

func NewConnect4SimpleHeuristic(playerID base.PlayerID) *Connect4SimpleHeuristic {
	var connect4PlayerID connect4.PlayerID
	switch playerID {
	case base.Player1:
		connect4PlayerID = connect4.Player1
	case base.Player2:
		connect4PlayerID = connect4.Player2
	default:
		return nil
	}

	return &Connect4SimpleHeuristic{connect4.NewSimpleHeuristic(connect4PlayerID)}
}

type Connect4ViabilityHeuristic struct {
	heuristic *connect4.ViabilityHeuristic
}

func (heuristic *Connect4ViabilityHeuristic) Heuristic(gameState base.GameState) float64 {
	return heuristic.heuristic.Heuristic(gameState.(*Connect4GameState).GetGameState())
}

func NewConnect4ViabilityHeuristic(playerID base.PlayerID) *Connect4ViabilityHeuristic {
	var connect4PlayerID connect4.PlayerID
	switch playerID {
	case base.Player1:
		connect4PlayerID = connect4.Player1
	case base.Player2:
		connect4PlayerID = connect4.Player2
	default:
		return nil
	}

	return &Connect4ViabilityHeuristic{connect4.NewViabilityHeuristic(connect4PlayerID)}
}

type Connect4ViabilityExtendedHeuristic struct {
	heuristic *connect4.ViabilityExtendedHeuristic
}

func (heuristic *Connect4ViabilityExtendedHeuristic) Heuristic(gameState base.GameState) float64 {
	return heuristic.heuristic.Heuristic(gameState.(*Connect4GameState).GetGameState())
}

func NewConnect4ViabilityExtendedHeuristic(playerID base.PlayerID) *Connect4ViabilityExtendedHeuristic {
	var connect4PlayerID connect4.PlayerID
	switch playerID {
	case base.Player1:
		connect4PlayerID = connect4.Player1
	case base.Player2:
		connect4PlayerID = connect4.Player2
	default:
		return nil
	}

	return &Connect4ViabilityExtendedHeuristic{connect4.NewViabilityExtendedHeuristic(connect4PlayerID)}
}
