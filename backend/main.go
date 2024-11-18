package main

import (
	"chess/engine"
	"chess/utils"
	"flag"
	"fmt"
	"os"
)


var TestBoard = []int {

	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

	99, 4, 2, 3, 5, 6, 3, 2, 4, 99,

	99, 1, 1, 1, 0, 1, 1, 1, 1, 99,

	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,

	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,

	99, 0, 0, 0, 1, -1, 0, 0, 0, 99,

	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,

	99, -1, -1, -1, -1, 0, -1, -1, -1, 99,

	99, -4, -2, -3, -5, -6, -3, -2, -4, 99,

	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

	};

func main() {
	gameFilePath := flag.String("game", "", "Chess game file")
	flag.Parse()
	gameFile, err := os.ReadFile(*gameFilePath)

	if err != nil {
		panic(err)
	}


	moves, result := utils.GetMovesFromPGN(string(gameFile))
	var boardState utils.BoardState

	boardState.Board = engine.Board
	boardState.ToMove = "white"
	fmt.Println(result)

	fmt.Println(boardState)

	for {
		if (boardState.CurrentMoveNum != len(moves)) {
			currentMove := moves[boardState.CurrentMoveNum]
			fmt.Println(currentMove)			
		}
	}
	
}