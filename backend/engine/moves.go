package engine

import (
	"chess/utils"
)

// pos arg =  8x8 position
func GetLegalMoves(boardState utils.BoardState, pos int) (int,[]int) {
	var legalMoves []int
	
	piece := boardState.Board[pos]

	switch piece {
	case 1:
		fallthrough
	case -1:
		legalMoves = PawnLegalMoves(boardState, pos, piece)
	case 2:
		fallthrough
	case -2:
		
	

	}


	return piece, legalMoves
}


// TODO pawn move 2 forward on start square
// returns an array of legal moves 
func PawnLegalMoves(boardState utils.BoardState, startPos int, piece int) []int {
	
	var legalMoves []int
	board := boardState.Board

	// first check in separate function if enpassant


	var moveForwardIndex, moveLeftIndex, moveRightIndex int
	var moveTwoIndex int = -1

	var canCaptureLeft, canCaptureRight bool



	if piece > 0 {
		moveForwardIndex = startPos + 10
		moveLeftIndex = startPos + 11
		moveRightIndex = startPos + 9

		canCaptureLeft = board[moveLeftIndex] < 0
		canCaptureRight = board[moveRightIndex] < 0

		if startPos / 10 == 3 {
			moveTwoIndex = startPos + 20
		}
	} else {
		moveForwardIndex = startPos - 10
		moveLeftIndex = startPos - 11
		moveRightIndex = startPos - 9
		
		canCaptureLeft = board[moveLeftIndex] > 0
		canCaptureRight = board[moveRightIndex] > 0

		if startPos / 10 == 8 {
			moveTwoIndex = startPos - 20
		} 
	}

	if (moveTwoIndex != -1 && board[moveTwoIndex] == 0) {
		legalMoves = append(legalMoves, moveTwoIndex)
	}


	if board[moveForwardIndex] == 0 {
		legalMoves = append(legalMoves, moveForwardIndex)
	}

	if board[moveLeftIndex] != 0 && board[moveLeftIndex] != utils.BOARD_BOUNDRY && canCaptureLeft  {
		legalMoves = append(legalMoves, moveLeftIndex)
	}

	if board[moveRightIndex] != 0 && board[moveRightIndex] != utils.BOARD_BOUNDRY && canCaptureRight {
		legalMoves = append(legalMoves, moveRightIndex)
	}

	return legalMoves
}


func KnightLegalMoves (startPos int, color int) []int {
	var legalMoves []int


	return legalMoves
}

func IsInCheck(board []int) {

}

func GetWhiteAtackSquares(board []int) {
	//var attackedSquares []int
	
	// get position of all white pieces 

}

func GetBlackAttackSquares(board []int) {

}


func GetPieceAttackSquares(boardState utils.BoardState, pos int) ([]int, int) {
	var attackedSquares []int

	board := boardState.Board

	piece := board[pos]

	switch piece {
	case 1:
		moveLeftIndex := pos + 11
		moveRightIndex := pos + 9
		if (board[moveLeftIndex] != utils.BOARD_BOUNDRY ) {
			attackedSquares = append(attackedSquares, moveLeftIndex)
		} 
		if (board[moveRightIndex] != utils.BOARD_BOUNDRY) {
			attackedSquares = append(attackedSquares, moveRightIndex)
		}
	case -1:
		moveLeftIndex := pos - 11
		moveRightIndex := pos - 9
		if (board[moveLeftIndex] != utils.BOARD_BOUNDRY ) {
			attackedSquares = append(attackedSquares, moveLeftIndex)
		} 
		if (board[moveRightIndex] != utils.BOARD_BOUNDRY) {
			attackedSquares = append(attackedSquares, moveRightIndex)
		}
	}




	return attackedSquares, piece
}