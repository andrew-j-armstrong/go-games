package base

type Turn int

const (
	Player1Turn Turn = iota
	Player2Turn
	NonPlayerTurn
	GameOver
)
