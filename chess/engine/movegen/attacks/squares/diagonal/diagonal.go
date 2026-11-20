package diagonal

import (
	"fmt"
	"github.com/tanzor-disco/GOndon/chess/board"
)

func GetUpperLeftDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
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

func GetUpperRightDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
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

func GetLowerLeftDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
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

func GetLowerRightDiagonal(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
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
