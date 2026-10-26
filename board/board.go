package board

import (
	"fmt"
	"strings"
)

type Board struct {
	White Side
	Black Side
}

func NewBoard() Board {
	white := NewSide(White)
	black := NewSide(Black)

	return Board{
		White: white,
		Black: black,
	}
}

func (b *Board) DrawBoard() (string, error) {
	var builder strings.Builder

	black, err := b.Black.DrawSide()
	if err != nil {
		return "", fmt.Errorf("DrawBoard: %w", err)
	}
	white, err := b.White.DrawSide()
	if err != nil {
		return "", fmt.Errorf("DrawBoard: %w", err)
	}

	for squareIdx, square := range white {
		squareString := string(square)
		if squareString != "\n" && squareString != EmptySquare {
			builder.WriteString(strings.ToUpper(squareString))
			continue
		}

		squareString = string(black[squareIdx])
		if squareString != "\n" && squareString != EmptySquare {
			builder.WriteString(squareString)
			continue
		}
		builder.WriteString(squareString)
	}
	return builder.String(), nil
}
