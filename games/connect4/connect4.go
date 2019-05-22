package games_connect4

import (
	"fmt"
	"log"

	"github.com/andrew-j-armstrong/go-connect4"
	"github.com/andrew-j-armstrong/go-games/base"
)

type Connect4 struct {
}

func (Connect4) ParseMove(moveDefinition string) (base.Move, error) {
	return connect4.ParseMove(moveDefinition)
}

func (Connect4) NewGameState() base.GameState {
	return NewConnect4GameState()
}

func (Connect4) LoadGameState(filename string) (base.GameState, error) {
	return LoadConnect4GameState(filename)
}

func (Connect4) ChooseHeuristic(playerID base.PlayerID) base.Heuristic {
	for {
		fmt.Printf("Choose heurstic for player %d:\n", playerID)
		fmt.Printf("1: Simple\n")
		fmt.Printf("2: Viability\n")
		fmt.Printf("3: Viability Extended\n")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			return NewConnect4SimpleHeuristic(playerID)
		case 2:
			return NewConnect4ViabilityHeuristic(playerID)
		case 3:
			return NewConnect4ViabilityExtendedHeuristic(playerID)
		}

		fmt.Println("Invalid choice!")
	}
}

func (Connect4) ParseHeuristic(heuristicDescription string, playerID base.PlayerID) base.Heuristic {
	switch heuristicDescription {
	case "simple":
		return NewConnect4SimpleHeuristic(playerID)
	case "viability":
		return NewConnect4ViabilityHeuristic(playerID)
	case "viabilityextended":
		return NewConnect4ViabilityExtendedHeuristic(playerID)
	default:
		log.Fatalf("invalid heuristic description %s", heuristicDescription)
		return nil
	}
}
