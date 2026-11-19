package king

import (
	"fmt"

	"github.com/Tanzor-Disco/GOndon/chess/board"
	"github.com/Tanzor-Disco/GOndon/utils/models/candidatesquare"
)

func getAttacksFromSquare(square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}
	candidates := [8]candidatesquare.CandidateSquare{
		{rank + 1, file},
		{rank + 1, file + 1},
		{rank, file + 1},
		{rank - 1, file + 1},
		{rank - 1, file},
		{rank - 1, file - 1},
		{rank, file - 1},
		{rank + 1, file - 1},
	}
	var attacks board.BitBoard
	for _, candidate := range candidates {
		if !candidate.IsValidSquare() {
			continue
		}

		attackSquare, err := board.GetSquareFromRankAndFile(candidate.Rank, candidate.File)
		if err != nil {
			return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
		}
		if err := attacks.SetPiece(attackSquare); err != nil {
			return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
		}
	}
	return attacks, nil
}

func GenerateAttacksForKing() ([64]board.BitBoard, error) {
	var attacks [64]board.BitBoard
	for square := range 64 {
		attacksFromSquare, err := getAttacksFromSquare(board.Square(square))
		if err != nil {
			return [64]board.BitBoard{}, fmt.Errorf("GenerateAttacksForKing: %w", err)
		}
		attacks[square] = attacksFromSquare
	}
	return attacks, nil
}
