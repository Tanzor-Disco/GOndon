package board

import (
	"fmt"
	"strings"
)

type SideColor uint8

const (
	White SideColor = iota
	Black

	EmptySquare = "."
)

type Side struct {
	Color   SideColor
	Pawns   BitBoard
	Bishops BitBoard
	Knights BitBoard
	Rooks   BitBoard
	Queens  BitBoard
	King    BitBoard
}

func createWhiteSide() Side {
	var pawns BitBoard
	for shift := range 8 {
		letter := 'a' + shift
		pawns.Set(string(letter) + "2")
	}

	var bishops BitBoard
	bishops.Set("c1")
	bishops.Set("f1")
	bishops.DrawBoard()

	var knights BitBoard
	knights.Set("b1")
	knights.Set("g1")

	var rooks BitBoard
	rooks.Set("a1")
	rooks.Set("h1")

	var queens BitBoard
	queens.Set("d1")

	var king BitBoard
	king.Set("e1")

	return Side{
		Color:   White,
		Pawns:   pawns,
		Bishops: bishops,
		Knights: knights,
		Rooks:   rooks,
		Queens:  queens,
		King:    king,
	}
}

func createBlackSide() Side {
	var pawns BitBoard
	for shift := range 8 {
		letter := 'a' + shift
		pawns.Set(string(letter) + "7")
	}

	var bishops BitBoard
	bishops.Set("c8")
	bishops.Set("f8")

	var knights BitBoard
	knights.Set("b8")
	knights.Set("g8")

	var rooks BitBoard
	rooks.Set("a8")
	rooks.Set("h8")

	var queens BitBoard
	queens.Set("d8")

	var king BitBoard
	king.Set("e8")

	return Side{
		Color:   Black,
		Pawns:   pawns,
		Bishops: bishops,
		Knights: knights,
		Rooks:   rooks,
		Queens:  queens,
		King:    king,
	}
}

func NewSide(color SideColor) Side {
	if color == White {
		return createWhiteSide()
	}
	return createBlackSide()
}

func fillBoard(board [][]string, tag string, bitBoard BitBoard) error {
	for idx := range 64 {
		filled, err := bitBoard.checkIdx(idx)
		if err != nil {
			return fmt.Errorf("fillBoard: %w", err)
		}
		if filled {
			// inverting to achieve the view from the white perspective
			board[7-idx/8][idx%8] = tag
		}
	}
	return nil
}

func (s *Side) DrawSide() (string, error) {
	var emptyRow []string
	for range 8 {
		emptyRow = append(emptyRow, EmptySquare)
	}

	board := make([][]string, 8)
	for rowIdx := range board {
		curr := make([]string, 8)
		copy(curr, emptyRow)
		board[rowIdx] = curr
	}

	if err := fillBoard(board, "p", s.Pawns); err != nil {
		return "", err
	}
	if err := fillBoard(board, "n", s.Knights); err != nil {
		return "", err
	}
	if err := fillBoard(board, "b", s.Bishops); err != nil {
		return "", err
	}
	if err := fillBoard(board, "r", s.Rooks); err != nil {
		return "", err
	}
	if err := fillBoard(board, "q", s.Queens); err != nil {
		return "", err
	}
	if err := fillBoard(board, "k", s.King); err != nil {
		return "", err
	}

	parts := make([]string, 0, len(board))
	for _, row := range board {
		parts = append(parts, strings.Join(row, ""))
	}
	return strings.Join(parts, "\n"), nil
}
