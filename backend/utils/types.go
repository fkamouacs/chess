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
	CurrentMoveNum      int
}

type Move struct {
	MoveNumber int    // The move number (e.g., 1, 2, 3)
	WhiteMove  string // The move made by the white player (e.g., "e4", "Nf3")
	BlackMove  string // The move made by the black player (e.g., "c5", "e6")
}
