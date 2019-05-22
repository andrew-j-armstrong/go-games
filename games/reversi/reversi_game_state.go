package games_reversi

import (
	"github.com/carbon-12/go-extensions"
	"github.com/carbon-12/go-games/base"
	"github.com/carbon-12/go-reversi"
)

type ReversiGameState struct {
	gameState *reversi.GameState
}

func (gameState *ReversiGameState) GetTurn() base.Turn {
	switch gameState.gameState.GetTurn() {
	case reversi.Player1Turn:
		return base.Player1Turn
	case reversi.Player2Turn:
		return base.Player2Turn
	default:
		return base.GameOver
	}
}

func (gameState *ReversiGameState) RegisterMoveListener(moveListener chan<- interface{}) {
	gameMoveListener := make(chan reversi.Move)
	gameState.gameState.RegisterMoveListener(gameMoveListener)
	go func() {
		for move := range gameMoveListener {
			moveListener <- move
		}
	}()
}

func (gameState *ReversiGameState) IsValidMove(move interface{}) bool {
	switch m := move.(type) {
	case reversi.Move:
		return gameState.gameState.IsValidMove(m)
	default:
		return false
	}
}

func (gameState *ReversiGameState) GetPossibleMoves() *extensions.InterfaceSlice {
	moves := gameState.gameState.GetPossibleMoves()
	genericMoves := make(extensions.InterfaceSlice, 0, len(moves))
	for _, move := range moves {
		genericMoves = append(genericMoves, move)
	}
	return &genericMoves
}

func (gameState *ReversiGameState) String() string {
	return gameState.gameState.String()
}

func (gameState *ReversiGameState) Print() {
	gameState.gameState.Print()
}

func (gameState *ReversiGameState) IsGameOver() bool {
	return gameState.gameState.IsGameOver()
}

func (gameState *ReversiGameState) MakeMove(move interface{}) error {
	return gameState.gameState.MakeMove(move.(reversi.Move))
}

func (gameState *ReversiGameState) GetGameState() *reversi.GameState {
	return gameState.gameState
}

func NewReversiGameState() *ReversiGameState {
	return &ReversiGameState{gameState: reversi.NewGame()}
}

func GetReversiGameState(gameState *reversi.GameState) *ReversiGameState {
	return &ReversiGameState{gameState: gameState}
}

func (gameState *ReversiGameState) Clone() interface{} {
	return &ReversiGameState{gameState: gameState.gameState.Clone()}
}

func (gameState *ReversiGameState) Save(filename string) error {
	return gameState.gameState.Save(filename)
}

func ParseReversiGameState(gameDescription string) (*ReversiGameState, error) {
	gameState, err := reversi.ParseGame(gameDescription)
	if err != nil {
		return nil, err
	}

	return &ReversiGameState{gameState: gameState}, nil
}

func LoadReversiGameState(filename string) (*ReversiGameState, error) {
	gameState, err := reversi.LoadGame(filename)
	if err != nil {
		return nil, err
	}

	return &ReversiGameState{gameState: gameState}, nil
}
