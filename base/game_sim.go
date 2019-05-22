package base

import "log"

type GameSim struct {
	gameState GameState
	player1   Player
	player2   Player
}

func (gameSim *GameSim) Run(printAfterMove bool) {
	for {
		if gameSim.gameState.IsGameOver() {
			break
		}

		var move interface{}
		switch gameSim.gameState.GetTurn() {
		case Player1Turn:
			move = gameSim.player1.GetNextMove()
		case Player2Turn:
			move = gameSim.player2.GetNextMove()
		default:
			log.Fatal("Invalid Turn!")
		}

		err := gameSim.gameState.MakeMove(move)

		if err != nil {
			log.Fatal(err)
		}

		if printAfterMove {
			gameSim.gameState.Print()
		}
	}
}

func NewGameSim(gameState GameState, player1 Player, player2 Player) *GameSim {
	return &GameSim{gameState: gameState, player1: player1, player2: player2}
}
