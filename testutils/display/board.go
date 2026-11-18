package display

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/chess/board"
	"strings"
)

func DisplayBoard(b board.Board) (string, error) {
	var builder strings.Builder

	black, err := DisplaySide(&b.Black)
	if err != nil {
		return "", fmt.Errorf("DrawBoard: %w", err)
	}
	white, err := DisplaySide(&b.White)
	if err != nil {
		return "", fmt.Errorf("DrawBoard: %w", err)
	}

	for squareIdx, square := range white {
		squareString := string(square)
		if squareString != "\n" && squareString != EmptySquareString {
			builder.WriteString(strings.ToUpper(squareString))
			continue
		}

		squareString = string(black[squareIdx])
		if squareString != "\n" && squareString != EmptySquareString {
			builder.WriteString(squareString)
			continue
		}
		builder.WriteString(squareString)
	}
	return builder.String(), nil
}
