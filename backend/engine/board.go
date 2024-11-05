package engine

import (
	"chess/utils"
	"errors"
	"fmt"
)

// 1D 12x10 array
var Board = [120]int {

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



func GetAlgebraic (index int) (string, int, error) {
	twoDRep, err := GetRowAndColumn(index)
	if (err != nil) {
		return "", -1, err
	}

	var file = utils.FileMap[twoDRep[1]]
	var rank = twoDRep[0] + 1

	return file, rank, nil
}

func GetRowAndColumn (index int) ([]int, error){
	// square = rank * 8 + file 
	
	index = index - 21

	fmt.Println(index)
	
	if (index < 0  || index > 63) {
		return nil, errors.New("index out of bounds")
	}

	var file int = ((index - index/ 10)  / 10) 

	var rank = (index - (file * 10)) 
	
	var result = []int{rank,file}
	fmt.Println(result)

	return result, nil
}