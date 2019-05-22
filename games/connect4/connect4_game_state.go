package games_connect4

import (
	"github.com/carbon-12/go-connect4"
	"github.com/carbon-12/go-extensions"
	"github.com/carbon-12/go-games/base"
)

type Connect4GameState struct {
	gameState *connect4.GameState
}

func (gameState *Connect4GameState) GetTurn() base.Turn {
	switch gameState.gameState.GetTurn() {
	case connect4.Player1Turn:
		return base.Player1Turn
	case connect4.Player2Turn:
		return base.Player2Turn
	default:
		return base.GameOver
	}
}

func (gameState *Connect4GameState) RegisterMoveListener(moveListener chan<- interface{}) {
	gameMoveListener := make(chan connect4.Move)
	gameState.gameState.RegisterMoveListener(gameMoveListener)
	go func() {
		for move := range gameMoveListener {
			moveListener <- move
		}
	}()
}

func (gameState *Connect4GameState) IsValidMove(move interface{}) bool {
	switch m := move.(type) {
	case connect4.Move:
		return gameState.gameState.IsValidMove(m)
	default:
		return false
	}
}

func (gameState *Connect4GameState) GetPossibleMoves() *extensions.InterfaceSlice {
	moves := gameState.gameState.GetPossibleMoves()
	genericMoves := make(extensions.InterfaceSlice, 0, len(moves))
	for _, move := range moves {
		genericMoves = append(genericMoves, move)
	}
	return &genericMoves
}

func (gameState *Connect4GameState) String() string {
	return gameState.gameState.String()
}

func (gameState *Connect4GameState) Print() {
	gameState.gameState.Print()
}

func (gameState *Connect4GameState) IsGameOver() bool {
	return gameState.gameState.IsGameOver()
}

func (gameState *Connect4GameState) MakeMove(move interface{}) error {
	return gameState.gameState.MakeMove(move.(connect4.Move))
}

func (gameState *Connect4GameState) GetGameState() *connect4.GameState {
	return gameState.gameState
}

func NewConnect4GameState() *Connect4GameState {
	return &Connect4GameState{gameState: connect4.NewGame()}
}

func GetConnect4GameState(gameState *connect4.GameState) *Connect4GameState {
	return &Connect4GameState{gameState: gameState}
}

func (gameState *Connect4GameState) Clone() interface{} {
	return &Connect4GameState{gameState: gameState.gameState.Clone()}
}

func (gameState *Connect4GameState) Save(filename string) error {
	return gameState.gameState.Save(filename)
}

func ParseConnect4GameState(gameDescription string) (*Connect4GameState, error) {
	gameState, err := connect4.ParseGame(gameDescription)
	if err != nil {
		return nil, err
	}

	return &Connect4GameState{gameState: gameState}, nil
}

func LoadConnect4GameState(filename string) (*Connect4GameState, error) {
	gameState, err := connect4.LoadGame(filename)
	if err != nil {
		return nil, err
	}

	return &Connect4GameState{gameState: gameState}, nil
}
