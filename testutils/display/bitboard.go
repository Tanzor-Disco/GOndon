package display

import (
	"fmt"
	"strings"

	"github.com/Tanzor-Disco/GOndon/chess/board"
)

func DisplayBitBoard(bitboard board.BitBoard) (string, error) {
	var builder strings.Builder
	for rank := 7; rank >= 0; rank-- {
		builder.WriteString("\n")
		for file := 0; file < 8; file++ {
			square, err := board.GetSquareFromRankAndFile(rank, file)
			if err != nil {
				return "", fmt.Errorf("DisplayBitBoard: %w", err)
			}
			if bitboard.HasPieceAtSquare(square) {
				builder.WriteString(TakenSquareString)
			} else {
				builder.WriteString(EmptySquareString)
			}
		}
	}
	return builder.String(), nil
}
