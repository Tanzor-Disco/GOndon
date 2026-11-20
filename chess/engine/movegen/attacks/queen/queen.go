package queen

import (
	"fmt"

	"github.com/tanzor-disco/GOndon/chess/board"
	"github.com/tanzor-disco/GOndon/chess/engine/movegen/attacks/squares/diagonal"
	"github.com/tanzor-disco/GOndon/chess/engine/movegen/attacks/squares/linear"
)

func GenerateAttacksForQueenFromSquare(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}

	left, err := linear.GetAttacksLeft(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}
	right, err := linear.GetAttacksRight(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}
	up, err := linear.GetAttacksUp(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}
	down, err := linear.GetAttacksDown(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}

	upperLeft, err := diagonal.GetUpperLeftDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}
	upperRight, err := diagonal.GetUpperRightDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}
	lowerLeft, err := diagonal.GetLowerLeftDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}
	lowerRight, err := diagonal.GetLowerRightDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForQueenFromSquare: %w", err)
	}

	return right | left | up | down | upperLeft | upperRight | lowerLeft | lowerRight, nil
}
