package utils

type BoardState struct {
	Board               []int
	ToMove              string
	MovesToDraw         int
	CanWhiteCastleKing  bool
	CanWhiteCastleQueen bool
	CanBlackCastleKing  bool
	CanBlackCastleQueen bool
	IsEnPassant         bool
}
