package base

type Game interface {
	ParseMove(moveDefinition string) (Move, error)
	NewGameState() GameState
	LoadGameState(filename string) (GameState, error)
	ChooseHeuristic(PlayerID) Heuristic
	ParseHeuristic(string, PlayerID) Heuristic
}
