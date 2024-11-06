package main

import (
	"chess/engine"
	"fmt"
)

func main() {
	// fmt.Println(engine.Board[39])
	// fmt.Println(engine.GetAlgebraic(39))

	fmt.Println(engine.GetLegalMoves(1, 42))
}