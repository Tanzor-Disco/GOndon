package bishop

import (
	"fmt"

	"github.com/tanzor-disco/GOndon/chess/board"
	"github.com/tanzor-disco/GOndon/chess/engine/movegen/attacks/squares/diagonal"
)

func GenerateAttacksForBishopFromSquare(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	upperLeft, err := diagonal.GetUpperLeftDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	upperRight, err := diagonal.GetUpperRightDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	lowerLeft, err := diagonal.GetLowerLeftDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	lowerRight, err := diagonal.GetLowerRightDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	return upperLeft | upperRight | lowerLeft | lowerRight, nil
}
