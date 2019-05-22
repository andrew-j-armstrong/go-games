package players

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/andrew-j-armstrong/go-expectimax"
	"github.com/andrew-j-armstrong/go-extensions"
	"github.com/andrew-j-armstrong/go-games/base"
)

func getExpectimaxHeuristic(heuristic base.Heuristic) expectimax.ExpectimaxHeuristic {
	return func(expectimaxGame expectimax.Game) float64 {
		return heuristic.Heuristic(expectimaxGame.(base.GameState))
	}
}

type Expectimax interface {
	RunExpectimax()
	IsCurrentlySearching() bool
	GetNextMoveValues() *extensions.ValueMap
}

type ExpectimaxPlayer struct {
	gameState      base.GameState
	player         base.PlayerID
	expectimaxBase Expectimax
	difficulty     float64
	lastChoiceTime time.Time
	maxSearchTime  time.Duration
}

func (player *ExpectimaxPlayer) Run() {
	player.lastChoiceTime = time.Now()
	player.expectimaxBase.RunExpectimax()
}

func (player *ExpectimaxPlayer) IsReadyToMakeMove() bool {
	if !player.expectimaxBase.IsCurrentlySearching() {
		return true
	}

	nextMoves := player.expectimaxBase.GetNextMoveValues()

	// Check the moves to determine whether there's enough difference or we should wait longer
	if len(*nextMoves) <= 1 {
		return true
	}

	moveValues := make([]float64, 0, len(*nextMoves))
	for _, value := range *nextMoves {
		moveValues = append(moveValues, value)
	}

	sort.Float64s(moveValues)

	return moveValues[0] >= 1.0 || moveValues[len(moveValues)-1] <= -1.0 || moveValues[len(moveValues)-1]-moveValues[len(moveValues)-2] >= 0.5*player.difficulty/100.0
}

func (player *ExpectimaxPlayer) buildSelectionWheel(moves *extensions.InterfaceSlice, getMoveValue func(interface{}) float64, playerDifficulty float64) *extensions.ValueMap {
	powerBase := math.Pow(100.0/(100.0-playerDifficulty), playerDifficulty/25.0)
	selectionWheel := extensions.ValueMap{}
	for _, move := range *moves {
		wheelValue := math.Pow(powerBase, 10.0*getMoveValue(move))
		selectionWheel[move] = wheelValue
	}

	return &selectionWheel
}

func (player *ExpectimaxPlayer) GetNextMove() interface{} {
	player.lastChoiceTime = time.Now()

	for time.Since(player.lastChoiceTime) < player.maxSearchTime && !player.IsReadyToMakeMove() {
		time.Sleep(time.Duration(50) * time.Millisecond)
	}

	nextMoves := player.expectimaxBase.GetNextMoveValues()

	if len(*nextMoves) == 0 {
		log.Fatal("No next moves!")
	}

	var nextMove interface{}
	if player.difficulty == 100.0 {
		nextMove = nextMoves.GetBestKey()
	} else {
		selectionWheel := player.buildSelectionWheel(nextMoves.GetKeys(), func(move interface{}) float64 { return (*nextMoves)[move] }, player.difficulty)
		nextMove = selectionWheel.SelectFromWheel()
	}

	player.lastChoiceTime = time.Now()

	return nextMove
}

func (player *ExpectimaxPlayer) calculateChildLikelihoodMap(getChildValue func(interface{}) float64, childLikelihood *extensions.ValueMap, playerDifficulty float64, minSpread float64) {
	selectionWheel := player.buildSelectionWheel(childLikelihood.GetKeys(), getChildValue, playerDifficulty)
	totalWheelValue := selectionWheel.GetTotalValue()

	minSpreadWeight := minSpread / float64(len(*selectionWheel))
	wheelValueMultiplier := (1.0 - minSpread) / totalWheelValue
	for move, wheelValue := range *selectionWheel {
		(*childLikelihood)[move] = minSpreadWeight + wheelValue*wheelValueMultiplier
	}
}

func (player *ExpectimaxPlayer) calculateChildLikelihood(getGameState func() expectimax.Game, getChildValue func(interface{}) float64, childLikelihood *extensions.ValueMap) {
	expectimaxGameState := getGameState()
	if expectimaxGameState == nil {
		return
	}

	gameState := expectimaxGameState.(base.GameState)

	if gameState.GetTurn() == base.Player1Turn && player.player == base.Player2 || gameState.GetTurn() == base.Player2Turn && player.player == base.Player1 {
		player.calculateChildLikelihoodMap(func(move interface{}) float64 { return -getChildValue(move) }, childLikelihood, 99, 0.02)
	} else {
		if player.difficulty == 100.0 {
			bestMove := childLikelihood.GetKeys().GetBestEntry(getChildValue)

			for move := range *childLikelihood {
				if move == bestMove {
					(*childLikelihood)[move] = 1.0
				} else {
					(*childLikelihood)[move] = 0.0
				}
			}
		} else {
			player.calculateChildLikelihoodMap(getChildValue, childLikelihood, player.difficulty, 0.0)
		}
	}
}

func NewExpectimaxPlayer(gameState base.GameState, playerID base.PlayerID, heuristic base.Heuristic, difficulty float64, maxSearchTime time.Duration) *ExpectimaxPlayer {
	player := &ExpectimaxPlayer{gameState, playerID, nil, difficulty, time.Time{}, maxSearchTime}
	player.expectimaxBase = expectimax.NewExpectimax(gameState, getExpectimaxHeuristic(heuristic), player.calculateChildLikelihood, 1000000)
	go player.Run()
	return player
}

func parseExpectimaxDifficulty(difficultyDescription string) (float64, error) {
	difficulty, err := strconv.ParseFloat(difficultyDescription, 64)

	if err != nil {
		return 0.0, err
	}

	if difficulty < 0 || difficulty > 100 {
		return 0.0, errors.New(fmt.Sprintf("invalid expectimax difficulty: %f is not between 0 and 100", difficulty))
	}

	return difficulty, nil
}

func chooseExpectimaxDifficulty() float64 {
	for {
		fmt.Printf("Choose difficulty (0-100): ")
		var choice float64
		fmt.Scan(&choice)

		if choice >= 0 && choice <= 100 {
			return choice
		}

		fmt.Println("Invalid choice!")
	}
}

func chooseExpectimaxDuration() time.Duration {
	for {
		fmt.Printf("Choose duration (ms): ")
		var choice float64
		fmt.Scan(&choice)

		if choice >= 0 && choice <= 30000 {
			return time.Duration(choice) * time.Millisecond
		}

		fmt.Println("Invalid choice!")
	}
}
