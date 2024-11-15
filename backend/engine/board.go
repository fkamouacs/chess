package engine

import (
	"chess/utils"
	"errors"
	"fmt"
)

// 1D 12x10 array
// var Board = [120]int {

// 	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

// 	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

// 	99, 0, 1, 2, 3, 4, 5, 6, 7, 99,

// 	99, 8, 9, 10, 11, 12, 13, 14, 15, 99,

// 	99, 16, 17, 18, 19, 20, 21, 22, 23, 99,

// 	99, -24, 0, 0, 27, 28, 29, 30, 31, 99,

// 	99, 32, 33, 34, 35, 36, 37, 38, 39, 99,

// 	99, 40, 41, 42, 43, 44, 45, 46, 47, 99,

// 	99, 48, 49, 50, 51, 52, 53, 54, 55, 99,

// 	99, 56, 57, 58, 59, 60, 61, 62, 63, 99,

// 	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

// 	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,

// 	};


func GetAlgebraic (index int) (string, int, error) {
	twoDRep, err := GetRowAndColumn(index)
	if (err != nil) {
		return "", -1, err
	}

	var file = utils.FileMap[twoDRep[0]]
	var rank = twoDRep[1] + 1

	return file, rank, nil
}

func GetRowAndColumn (index int) ([]int, error){
	const INIT_PADDING = 21
	const MAX_INDEX = 98

	if (index < INIT_PADDING || index > MAX_INDEX || index % 10 == 0 || index % 10 == 9) {
		return nil, errors.New("index out of bounds")
	}

	// index = rank * 10 + file + 21
	// rank = (index - 21) / 10 
	// file = (index - 21) - rank * 10

	square := index - INIT_PADDING

	rank := square / 10
	file := square - rank * 10
	
	var result = []int{file,rank}

	return result, nil
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

