package display

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/board"
	"strings"
)

func DrawBoard(b board.Board) (string, error) {
	var builder strings.Builder

	black, err := DrawSide(&b.Black)
	if err != nil {
		return "", fmt.Errorf("DrawBoard: %w", err)
	}
	white, err := DrawSide(&b.White)
	if err != nil {
		return "", fmt.Errorf("DrawBoard: %w", err)
	}

	for squareIdx, square := range white {
		squareString := string(square)
		if squareString != "\n" && squareString != board.EmptySquareString {
			builder.WriteString(strings.ToUpper(squareString))
			continue
		}

		squareString = string(black[squareIdx])
		if squareString != "\n" && squareString != board.EmptySquareString {
			builder.WriteString(squareString)
			continue
		}
		builder.WriteString(squareString)
	}
	return builder.String(), nil
}
