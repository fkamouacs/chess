package engine

import "chess/utils"

func GetLegalMoves(piece int, startPos int) []int {
	var legalMoves []int
	const WHITE = 1
	const BLACK = 0

	switch piece {
	case 1:
		{
			legalMoves = PawnLegalMoves(startPos, WHITE)
		}
	case -1:
		{
			legalMoves = PawnLegalMoves(startPos, BLACK)
		}
	case 2:
		{
			legalMoves = KnightLegalMoves(startPos, WHITE)
		}

	}

	return legalMoves
}

func PawnLegalMoves(startPos int, color int) []int {

	var legalMoves []int
	var moveForwardIndex, moveLeftIndex, moveRightIndex int

	var canCaptureLeft, canCaptureRight bool

	if color > 0 {
		moveForwardIndex = startPos + 10
		moveLeftIndex = startPos + 11
		moveRightIndex = startPos + 9

		canCaptureLeft = Board[moveLeftIndex] < 0
		canCaptureRight = Board[moveRightIndex] < 0
	} else {
		moveForwardIndex = startPos - 10
		moveLeftIndex = startPos - 11
		moveRightIndex = startPos - 9
		
		canCaptureLeft = Board[moveLeftIndex] > 0
		canCaptureRight = Board[moveRightIndex] > 0
	}

	if Board[moveForwardIndex] == 0 {
		legalMoves = append(legalMoves, moveForwardIndex)
	}

	if Board[moveLeftIndex] != 0 && Board[moveLeftIndex] != utils.BOARD_BOUNDRY && canCaptureLeft  {
		legalMoves = append(legalMoves, moveLeftIndex)
	}

	if Board[moveRightIndex] != 0 && Board[moveRightIndex] != utils.BOARD_BOUNDRY && canCaptureRight {
		legalMoves = append(legalMoves, moveRightIndex)
	}

	return legalMoves
}


func KnightLegalMoves (startPos int, color int) []int {
	var legalMoves []int


	return legalMoves
}