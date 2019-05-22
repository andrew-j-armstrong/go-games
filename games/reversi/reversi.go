package games_reversi

import (
	"fmt"
	"log"

	"github.com/andrew-j-armstrong/go-games/base"
	"github.com/andrew-j-armstrong/go-reversi"
)

type Reversi struct {
}

func (Reversi) ParseMove(moveDefinition string) (base.Move, error) {
	return reversi.ParseMove(moveDefinition)
}

func (Reversi) NewGameState() base.GameState {
	return NewReversiGameState()
}

func (Reversi) LoadGameState(filename string) (base.GameState, error) {
	return LoadReversiGameState(filename)
}

func (Reversi) ChooseHeuristic(playerID base.PlayerID) base.Heuristic {
	for {
		fmt.Printf("Choose heurstic for player %d:\n", playerID)
		fmt.Printf("1: Standard\n")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			return NewReversiStandardHeuristic(playerID)
		}

		fmt.Println("Invalid choice!")
	}
}

func (Reversi) ParseHeuristic(heuristicDescription string, playerID base.PlayerID) base.Heuristic {
	switch heuristicDescription {
	case "standard":
		return NewReversiStandardHeuristic(playerID)
	default:
		log.Fatalf("invalid heuristic description %s", heuristicDescription)
		return nil
	}
}
