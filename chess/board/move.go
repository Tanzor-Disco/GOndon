package board

import ()

// Move is used to describe a chess move in UCI notation
// The start and end coordinates a described
// By a set of uint8 constants
// PieceType is a pointer to differentiate the empty state
type Move struct {
	From      Square
	To        Square
	Promotion *PieceType
}

func NewMove(from, to Square, promotion *PieceType) Move {
	return Move{
		From:      from,
		To:        to,
		Promotion: promotion,
	}
}

func IsValidMove(move Move) bool {
	if !IsValidSquare(move.To) ||
		!IsValidSquare(move.From) ||
		(move.Promotion != nil && !IsValidPieceType(*move.Promotion)) {
		return false
	}
	return true
}
