package display

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/board"
	"strings"
)

func fillBoard(board [][]string, tag string, bitBoard board.BitBoard) error {
	for idx := range 64 {
		filled, err := bitBoard.HasPieceAtIdx(idx)
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

func DrawSide(side *board.Side) (string, error) {
	var emptyRow []string
	for range 8 {
		emptyRow = append(emptyRow, board.EmptySquareString)
	}

	currBoard := make([][]string, 8)
	for rowIdx := range currBoard {
		curr := make([]string, 8)
		copy(curr, emptyRow)
		currBoard[rowIdx] = curr
	}

	if err := fillBoard(currBoard, "p", side.Pieces[board.Pawn]); err != nil {
		return "", err
	}
	if err := fillBoard(currBoard, "n", side.Pieces[board.Knight]); err != nil {
		return "", err
	}
	if err := fillBoard(currBoard, "b", side.Pieces[board.Bishop]); err != nil {
		return "", err
	}
	if err := fillBoard(currBoard, "r", side.Pieces[board.Rook]); err != nil {
		return "", err
	}
	if err := fillBoard(currBoard, "q", side.Pieces[board.Queen]); err != nil {
		return "", err
	}
	if err := fillBoard(currBoard, "k", side.Pieces[board.King]); err != nil {
		return "", err
	}

	parts := make([]string, 0, len(currBoard))
	for _, row := range currBoard {
		parts = append(parts, strings.Join(row, ""))
	}
	return strings.Join(parts, "\n"), nil
}
