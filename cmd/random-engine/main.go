package main

import (
	"github.com/corentings/chess/v2/internal/uciserver"
)

func main() {
	uciserver.Run(NewRandomEngine())
}
