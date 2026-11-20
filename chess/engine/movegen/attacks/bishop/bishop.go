package bishop

import (
	"fmt"

	"github.com/tanzor-disco/GOndon/chess/board"
)

func getUpperLeftDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getUpperLeftDiagonal: %w", err)
	}
	var attacks board.BitBoard
	for rank, file = rank+1, file-1; rank < 8 && file >= 0; rank, file = rank+1, file-1 {
		attackSquare, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getUpperLeftDiagonal: %w", err)
		}
		if err := attacks.SetPiece(attackSquare); err != nil {
			return 0, fmt.Errorf("getUpperLeftDiagonal: %w", err)
		}
		if occupied.HasPieceAtSquare(attackSquare) {
			break
		}
	}
	return attacks, nil
}

func getUpperRightDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getUpperRightDiagonal: %w", err)
	}
	var attacks board.BitBoard
	for rank, file = rank+1, file+1; rank < 8 && file < 8; rank, file = rank+1, file+1 {
		attackSquare, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getUpperRightDiagonal: %w", err)
		}
		if err := attacks.SetPiece(attackSquare); err != nil {
			return 0, fmt.Errorf("getUpperRightDiagonal: %w", err)
		}
		if occupied.HasPieceAtSquare(attackSquare) {
			break
		}
	}
	return attacks, nil
}

func getLowerLeftDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getLowerLeftDiagonal: %w", err)
	}
	var attacks board.BitBoard
	for rank, file = rank-1, file-1; rank >= 0 && file >= 0; rank, file = rank-1, file-1 {
		attackSquare, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getLowerLeftDiagonal: %w", err)
		}
		if err := attacks.SetPiece(attackSquare); err != nil {
			return 0, fmt.Errorf("getLowerLeftDiagonal: %w", err)
		}
		if occupied.HasPieceAtSquare(attackSquare) {
			break
		}
	}
	return attacks, nil
}

func getLowerRightDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getLowerRightDiagonal: %w", err)
	}
	var attacks board.BitBoard
	for rank, file = rank-1, file+1; rank >= 0 && file < 8; rank, file = rank-1, file+1 {
		attackSquare, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getLowerRightDiagonal: %w", err)
		}
		if err := attacks.SetPiece(attackSquare); err != nil {
			return 0, fmt.Errorf("getLowerRightDiagonal: %w", err)
		}
		if occupied.HasPieceAtSquare(attackSquare) {
			break
		}
	}
	return attacks, nil
}

func GenerateAttacksForBishopFromSquare(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	upperLeft, err := getUpperLeftDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	upperRight, err := getUpperRightDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	lowerLeft, err := getLowerLeftDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	lowerRight, err := getLowerRightDiagonal(occupied, square)
	if err != nil {
		return 0, fmt.Errorf("GenerateAttacksForBishopFromSquare: %w", err)
	}
	return upperLeft | upperRight | lowerLeft | lowerRight, nil
}
