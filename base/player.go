package base

type PlayerID int

const (
	Player1 PlayerID = iota
	Player2
)

func (playerID PlayerID) String() string {
	switch playerID {
	case Player1:
		return "1"
	case Player2:
		return "2"
	}

	return "<Unknown>"
}

type Player interface {
	GetNextMove() interface{}
}
