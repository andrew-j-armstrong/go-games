package players

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/andrew-j-armstrong/go-games/base"
)

func ChoosePlayer(game base.Game, gameState base.GameState, playerID base.PlayerID) base.Player {
	for {
		fmt.Printf("Choose player %s:\n", playerID.String())
		fmt.Printf("1: Human\n")
		fmt.Printf("2: Randy\n")
		fmt.Printf("3: Huey\n")
		fmt.Printf("4: Max\n")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			return NewHumanPlayer(game, gameState)
		case 2:
			return NewRandomPlayer(gameState)
		case 3:
			return NewHeuristicPlayer(gameState, game.ChooseHeuristic(playerID))
		case 4:
			difficulty := chooseExpectimaxDifficulty()
			duration := chooseExpectimaxDuration()
			expectimaxPlayer := NewExpectimaxPlayer(gameState, playerID, game.ChooseHeuristic(playerID), difficulty, duration)
			return expectimaxPlayer
		}

		fmt.Println("Invalid choice!")
	}
}

var playerDescriptionRegex = regexp.MustCompile(`(human)|(random)|(heuristic)/(simple|viability|viabilityextended)|(expectimax)/(simple|viability|viabilityextended)/(\d+)/(\d+)`)

func ParsePlayer(game base.Game, gameState base.GameState, playerID base.PlayerID, playerDescription string) (base.Player, error) {
	match := playerDescriptionRegex.FindStringSubmatch(playerDescription)
	if match == nil {
		log.Fatal("invalid player description", playerDescription)
	}

	if match[1] == "human" {
		fmt.Printf("Player %s: Human\n", playerID.String())
		return NewHumanPlayer(game, gameState), nil
	} else if match[2] == "random" {
		fmt.Printf("Player %s: Random\n", playerID.String())
		return NewRandomPlayer(gameState), nil
	} else if match[3] == "heuristic" {
		fmt.Printf("Player %s: Heuristic\n", playerID.String())
		return NewHeuristicPlayer(gameState, game.ParseHeuristic(match[4], playerID)), nil
	} else if match[5] == "expectimax" {
		fmt.Printf("Player %s: Expectimax\n", playerID.String())

		difficulty, err := parseExpectimaxDifficulty(match[7])
		if err != nil {
			return nil, err
		}

		heuristic := game.ParseHeuristic(match[6], playerID)
		maxDurationMilliseconds, _ := strconv.Atoi(match[8])
		maxDuration := time.Duration(maxDurationMilliseconds) * time.Millisecond

		expectimaxPlayer := NewExpectimaxPlayer(gameState, playerID, heuristic, difficulty, maxDuration)

		go expectimaxPlayer.Run()
		return expectimaxPlayer, nil
	} else {
		return nil, fmt.Errorf("unknown player type %s", playerDescription)
	}
}
