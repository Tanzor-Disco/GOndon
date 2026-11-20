package pawn

import (
	"fmt"

	"github.com/tanzor-disco/GOndon/chess/board"
	"github.com/tanzor-disco/GOndon/utils/models/candidatesquare"
)

func getWhiteAttacksFromSquare(square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getWhiteAttacksFromSquare: %w", err)
	}

	var attacks board.BitBoard

	candidates := [2]candidatesquare.CandidateSquare{
		{rank + 1, file - 1},
		{rank + 1, file + 1},
	}

	for _, candidate := range candidates {
		if !candidate.IsValidSquare() {
			continue
		}

		square, err := board.GetSquareFromRankAndFile(candidate.Rank, candidate.File)
		if err != nil {
			return 0, fmt.Errorf("getWhiteAttacksFromSquare: %w", err)
		}
		attacks.SetPiece(square)
	}
	return attacks, nil
}

func getBlackAttacksFromSquare(square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getWhiteAttacksFromSquare: %w", err)
	}

	var attacks board.BitBoard

	candidates := [2]candidatesquare.CandidateSquare{
		{rank - 1, file - 1},
		{rank - 1, file + 1},
	}

	for _, candidate := range candidates {
		if !candidate.IsValidSquare() {
			continue
		}

		square, err := board.GetSquareFromRankAndFile(candidate.Rank, candidate.File)
		if err != nil {
			return 0, fmt.Errorf("getWhiteAttacksFromSquare: %w", err)
		}
		attacks.SetPiece(square)
	}
	return attacks, nil
}

func getAttacksFromSquare(square board.Square) ([2]board.BitBoard, error) {
	whiteAttacks, err := getWhiteAttacksFromSquare(square)
	if err != nil {
		return [2]board.BitBoard{}, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	blackAttacks, err := getBlackAttacksFromSquare(square)
	if err != nil {
		return [2]board.BitBoard{}, fmt.Errorf("getBlackAttacksFromSquare: %w", err)
	}
	var attacks [2]board.BitBoard
	attacks[board.White] = whiteAttacks
	attacks[board.Black] = blackAttacks
	return attacks, nil
}

func GenerateAttacksForPawn() ([64][2]board.BitBoard, error) {
	var attacks [64][2]board.BitBoard
	for square := range 64 {
		squareAttacks, err := getAttacksFromSquare(board.Square(square))
		if err != nil {
			return [64][2]board.BitBoard{}, fmt.Errorf("GenerateAttacksForPawn: %w", err)
		}
		attacks[square] = squareAttacks
	}
	return attacks, nil
}
