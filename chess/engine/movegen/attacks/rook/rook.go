package rook

import (
	"fmt"

	"github.com/tanzor-disco/GOndon/chess/board"
)

func getAttacksLeft(occupied board.BitBoard, rank, file int) (board.BitBoard, error) {
	var leftAttacks board.BitBoard
	for file--; file >= 0; file-- {
		square, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getAttacksLeft: %w", err)
		}

		if err := leftAttacks.SetPiece(square); err != nil {
			return 0, fmt.Errorf("getAttacksLeft: %w", err)
		}

		if occupied.HasPieceAtSquare(square) {
			break
		}
	}
	return leftAttacks, nil
}

func getAttacksRight(occupied board.BitBoard, rank, file int) (board.BitBoard, error) {
	var rightAttacks board.BitBoard
	for file++; file < 8; file++ {
		square, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getAttacksRight: %w", err)
		}

		if err := rightAttacks.SetPiece(square); err != nil {
			return 0, fmt.Errorf("getAttacksRight: %w", err)
		}

		if occupied.HasPieceAtSquare(square) {
			break
		}
	}
	return rightAttacks, nil
}

func getAttacksDown(occupied board.BitBoard, rank, file int) (board.BitBoard, error) {
	var downAttacks board.BitBoard
	for rank--; rank >= 0; rank-- {
		square, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getAttacksDown: %w", err)
		}

		if err := downAttacks.SetPiece(square); err != nil {
			return 0, fmt.Errorf("getAttacksDown: %w", err)
		}

		if occupied.HasPieceAtSquare(square) {
			break
		}
	}
	return downAttacks, nil
}

func getAttacksUp(occupied board.BitBoard, rank, file int) (board.BitBoard, error) {
	var upAttacks board.BitBoard
	for rank++; rank < 8; rank++ {
		square, err := board.GetSquareFromRankAndFile(rank, file)
		if err != nil {
			return 0, fmt.Errorf("getAttacksUp: %w", err)
		}

		if err := upAttacks.SetPiece(square); err != nil {
			return 0, fmt.Errorf("getAttacksUp: %w", err)
		}

		if occupied.HasPieceAtSquare(square) {
			break
		}
	}
	return upAttacks, nil
}

func GenerateAttacksForRookFromSquare(occupied board.BitBoard, square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	var attacks board.BitBoard
	left, err := getAttacksLeft(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	right, err := getAttacksRight(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	up, err := getAttacksUp(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	down, err := getAttacksDown(occupied, rank, file)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	attacks |= left | right | up | down
	return attacks, nil
}
