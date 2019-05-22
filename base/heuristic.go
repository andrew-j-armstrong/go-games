package base

type Heuristic interface {
	Heuristic(gameState GameState) float64
}
