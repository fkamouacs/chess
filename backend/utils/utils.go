package utils

const BOARD_BOUNDRY int = 99

var FileMap = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
	6: "g",
	7: "h",
}

var PieceEncoding = map[string]int{
	"empty":        0,
	"white pawn":   1,
	"black pawn":   -1,
	"white knight": 2,
	"black knight": -2,
	"white bishop": 3,
	"black bishop": -3,
	"white rook":   4,
	"black rook":   -4,
	"white queen":  5,
	"black queen":  -5,
	"white king":   6,
	"black king":   -6,
}

func RemoveElement(slice []int, value int) []int {
	result := []int{}
	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}