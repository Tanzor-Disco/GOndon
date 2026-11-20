package rook

import (
	"fmt"

	"github.com/tanzor-disco/GOndon/chess/board"
	"github.com/tanzor-disco/GOndon/chess/engine/movegen/attacks/squares/linear"
)

func GenerateAttacksForRookFromSquare(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForRookFromSquare: %w", err)
	}

	left, err := linear.GetAttacksLeft(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForRookFromSquare: %w", err)
	}
	right, err := linear.GetAttacksRight(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForRookFromSquare: %w", err)
	}
	up, err := linear.GetAttacksUp(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForRookFromSquare: %w", err)
	}
	down, err := linear.GetAttacksDown(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForRookFromSquare: %w", err)
	}
	return left | right | up | down, nil
}
