package main

import (
	"math/rand"

	chess "github.com/corentings/chess/v2"
)

// RandomEngine implements the uciserver.Engine interface by picking random legal moves.
type RandomEngine struct{}

func NewRandomEngine() *RandomEngine {
	return &RandomEngine{}
}

func (e *RandomEngine) Name() string {
	return "RandomEngine"
}

func (e *RandomEngine) Author() string {
	return "Stefan"
}

// ChooseMove selects a random legal move from the current game state.
func (e *RandomEngine) ChooseMove(game *chess.Game) *chess.Move {
	validMoves := game.ValidMoves()
	if len(validMoves) == 0 {
		return nil
	}
	move := validMoves[rand.Intn(len(validMoves))]
	return &move
}
