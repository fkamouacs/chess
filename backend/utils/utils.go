package utils

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
	"white bishop": 2,
	"black bishop": -2,
	"white knight": 3,
	"black knight": -3,
	"white rook":   4,
	"black rook":   -4,
	"white queen":  5,
	"black queen":  -5,
	"white king":   6,
	"black king":   -6,
}
