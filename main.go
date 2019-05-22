package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/andrew-j-armstrong/go-games/base"
	connect4 "github.com/andrew-j-armstrong/go-games/games/connect4"
	reversi "github.com/andrew-j-armstrong/go-games/games/reversi"
	"github.com/andrew-j-armstrong/go-games/players"
)

func parseGame(gameName string) (base.Game, error) {
	switch strings.ToLower(gameName) {
	case "connect4":
		return connect4.Connect4{}, nil
	case "reversi":
		return reversi.Reversi{}, nil
	}

	return nil, errors.New(fmt.Sprintf("invalid game name '%s'", gameName))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	runGamePtr := flag.Bool("run", false, "Run a game")
	loadGamePtr := flag.String("load", "", "Load a game from <filename>")
	//saveGamePtr := flag.String("save", "", "Save the game after the move to <filename>")
	player1Ptr := flag.String("player1", "", "Player 1 player description")
	player2Ptr := flag.String("player2", "", "Player 2 player description")
	makeMovePtr := flag.String("makeMove", "", "Make this move")
	//nextPlayerPtr := flag.String("nextPlayer", "", "Next player description")
	gamePtr := flag.String("game", "", "Game to play (connect4)")
	printAfterMovePtr := flag.Bool("printAfterMove", false, "Print the game state after each move")

	flag.Parse()

	var game base.Game
	var err error
	game, err = parseGame(*gamePtr)
	if err != nil {
		log.Fatal(err)
	}

	var gameState base.GameState
	var player1 base.Player
	var player2 base.Player

	if (*loadGamePtr) == "" {
		gameState = game.NewGameState()
	} else {
		// Load from file
		gameState, err = game.LoadGameState(*loadGamePtr)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Loaded game:")
		gameState.Print()
	}

	if gameState.IsGameOver() {
		gameState.Print()
		return
	}

	if *makeMovePtr != "" {
		move, err := game.ParseMove(*makeMovePtr)
		if err != nil {
			log.Fatalf("invalid move: %s", err)
		}
		if !gameState.IsValidMove(move) {
			log.Fatalf("invalid move: %s", move.String())
		}

		err = gameState.MakeMove(move)
		if err != nil {
			log.Fatal(err)
		}

		gameState.Print()
	}

	if gameState.IsGameOver() {
		gameState.Print()
		return
	}

	if *runGamePtr {
		if *player1Ptr == "" {
			player1 = players.ChoosePlayer(game, gameState, base.Player1)
		} else {
			player1, err = players.ParsePlayer(game, gameState, base.Player1, *player1Ptr)
			if err != nil {
				log.Fatal(err)
			}
		}

		if *player2Ptr == "" {
			player2 = players.ChoosePlayer(game, gameState, base.Player2)
		} else {
			player2, err = players.ParsePlayer(game, gameState, base.Player2, *player2Ptr)
			if err != nil {
				log.Fatal(err)
			}
		}

		gameSim := base.NewGameSim(gameState, player1, player2)
		gameSim.Run(*printAfterMovePtr)
		gameState.Print()
	} /*else if *nextPlayerPtr != "" || (game.GetTurn(gameState) == Player1Turn && *player1Ptr != "") || (game.GetTurn(gameState) == Player2Turn && *player2Ptr != "") {
		// Perform one move
		var player Player
		switch game.GetTurn(gameState) {
		case Player1Turn:
			if *nextPlayerPtr != "" {
				player = game.ParsePlayer(gameState, Player1, *nextPlayerPtr)
			} else if *player1Ptr == "" {
				player = game.ChoosePlayer(gameState, Player1)
			} else {
				player = game.ParsePlayer(gameState, Player1, *player1Ptr)
			}
		case Player2Turn:
			if *nextPlayerPtr != "" {
				player = game.ParsePlayer(gameState, Player2, *nextPlayerPtr)
			} else if *player2Ptr == "" {
				player = game.ChoosePlayer(gameState, Player2)
			} else {
				player = game.ParsePlayer(gameState, Player2, *player2Ptr)
			}
		}

		gameState.MakeMove(player.GetNextMove())
		gameState.Print()

		if *saveGamePtr != "" {
			gameState.Save(*saveGamePtr)
		}
	} else {
		gameState.Print()
	}*/
}
