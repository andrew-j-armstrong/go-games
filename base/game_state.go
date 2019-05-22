package base

import (
	"github.com/andrew-j-armstrong/go-extensions"
)

type GameState interface {
	Print()
	Save(filename string) error
	MakeMove(move interface{}) error
	IsGameOver() bool
	IsValidMove(move interface{}) bool
	GetPossibleMoves() *extensions.InterfaceSlice
	Clone() interface{}
	RegisterMoveListener(chan<- interface{})
	GetTurn() Turn
}
