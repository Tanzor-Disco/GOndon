package knight

import (
	"fmt"

	"github.com/Tanzor-Disco/GOndon/chess/board"
)

type candidateMove struct {
	rank int
	file int
}

func getAttacksFromSquare(square board.Square) (board.BitBoard, error) {
	rank, file, err := board.GetRankAndFileFromSquare(square)
	if err != nil {
		return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
	}

	candidates := [8]candidateMove{
		{rank - 2, file - 1},
		{rank - 2, file + 1},
		{rank - 1, file - 2},
		{rank + 1, file - 2},
		{rank + 2, file - 1},
		{rank + 2, file + 1},
		{rank - 1, file + 2},
		{rank + 1, file + 2},
	}

	var attacks board.BitBoard

	for _, candidate := range candidates {
		if candidate.file < 8 && candidate.file >= 0 &&
			candidate.rank < 8 && candidate.rank >= 0 {
			square := board.Square(candidate.rank*8 + candidate.file)
			if err := attacks.SetPiece(square); err != nil {
				return 0, fmt.Errorf("getAttacksFromSquare: %w", err)
			}
		}
	}
	return attacks, nil
}

func GenerateAttacksForKnight() ([64]board.BitBoard, error) {
	var attacks [64]board.BitBoard
	for square := range 64 {
		attacksFromSquare, err := getAttacksFromSquare(board.Square(square))
		if err != nil {
			return [64]board.BitBoard{}, fmt.Errorf("GenerateAttacksForKnight: %w", err)
		}
		attacks[square] = attacksFromSquare
	}
	return attacks, nil
}
