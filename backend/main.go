package main

import (
	"chess/engine"
	"fmt"
)

func main () {
	fmt.Println(engine.Board[31])
	fmt.Println(engine.GetAlgebraic(31))
}