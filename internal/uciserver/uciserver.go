package uciserver

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	chess "github.com/corentings/chess/v2"
)

// Engine defines the methods an engine must implement to be run by the UCI server.
type Engine interface {
	Name() string
	Author() string
	ChooseMove(game *chess.Game) *chess.Move
}

// Run starts the UCI protocol loop on stdin/stdout for the given Engine.
func Run(engine Engine) {
	scanner := bufio.NewScanner(os.Stdin)
	var game *chess.Game = chess.NewGame()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]

		switch cmd {
		case "uci":
			fmt.Printf("id name %s\n", engine.Name())
			fmt.Printf("id author %s\n", engine.Author())
			fmt.Println("uciok")

		case "isready":
			fmt.Println("readyok")

		case "ucinewgame":
			game = chess.NewGame()

		case "position":
			game = parsePosition(parts[1:])

		case "go":
			move := engine.ChooseMove(game)
			if move == nil {
				fmt.Println("bestmove (none)")
			} else {
				fmt.Printf("bestmove %s\n", chess.UCINotation{}.Encode(nil, move))
			}

		case "quit":
			return
		}
	}
}

// parsePosition parses:
//   - "position startpos [moves e2e4 e7e5 ...]"
//   - "position fen <fenstring> [moves ...]"
func parsePosition(args []string) *chess.Game {
	if len(args) == 0 {
		return chess.NewGame()
	}

	var game *chess.Game
	movesIdx := -1

	if args[0] == "startpos" {
		game = chess.NewGame()
		for i, arg := range args {
			if arg == "moves" {
				movesIdx = i + 1
				break
			}
		}
	} else if args[0] == "fen" {
		var fenTokens []string
		for i := 1; i < len(args); i++ {
			if args[i] == "moves" {
				movesIdx = i + 1
				break
			}
			fenTokens = append(fenTokens, args[i])
		}
		fenStr := strings.Join(fenTokens, " ")
		fenFunc, err := chess.FEN(fenStr)
		if err != nil {
			game = chess.NewGame()
		} else {
			game = chess.NewGame(fenFunc)
		}
	} else {
		game = chess.NewGame()
	}

	if movesIdx != -1 && movesIdx < len(args) {
		for _, mStr := range args[movesIdx:] {
			_ = game.PushNotationMove(mStr, chess.UCINotation{}, nil)
		}
	}

	return game
}
