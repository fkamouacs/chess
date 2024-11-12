package main

import (
	"chess/engine"
	"chess/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(engine.Board[39])
	// fmt.Println(engine.GetAlgebraic(39))

	args := os.Args

	data, err := os.ReadFile(args[1])

	if err != nil {
		panic(err)
	}

	intBoardArr := StringToIntArray(string(data))

	board := PopulateBoard(intBoardArr)

	boardState := utils.BoardState{Board: board}

	fmt.Println(engine.GetLegalMoves(boardState, 34))

	fmt.Println(engine.GetPieceAttackSquares(boardState, 33))
}


func StringToIntArray(str string) []int {
	var result []int
	words := strings.Fields(str)
	for _, word := range words {
		num, err := strconv.Atoi(word)

		if err != nil {
			panic(err)
		}

		result = append(result, num)
	}
	
	return result
}



func PopulateBoard(board []int) []int{
	var result = []int {

		99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
		
		99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
		
		99, 0, 1, 2, 3, 4, 5, 6, 7, 99,
		
		99, 8, 9, 10, 11, 12, 13, 14, 15, 99,
		
		99, 16, 17, 18, 19, 20, 21, 22, 23, 99,
		
		99, 24, 25, 26, 27, 28, 29, 30, 31, 99,
		
		99, 32, 33, 34, 35, 36, 37, 38, 39, 99,
		
		99, 40, 41, 42, 43, 44, 45, 46, 47, 99,
		
		99, 48, 49, 50, 51, 52, 53, 54, 55, 99,
		
		99, 56, 57, 58, 59, 60, 61, 62, 63, 99,
		
		99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
		
		99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
		
		};
		j := 21
		for i := 0; i < len(board); i++  {

			if j % 10 == 9 {
				j += 2
			}
			result[j] = board[i] 
			j++
		}
		return result
}


func PrintBoard(board []int) {
	temp_board := utils.RemoveElement(board, 99)
	
	fmt.Print("   a  b  c  d  e  f  g  h")

	for i := 0; i < len(temp_board); i++ {
		
		if i % 8 == 0 {
			fmt.Println("")
			fmt.Print(int(i / 8) + 1, " ")
		}
		
		fmt.Print(temp_board[i], " ")

		if (temp_board[i] >= 0) {
			fmt.Print(" ")
		}
		
		
    }
	fmt.Println("")
}

