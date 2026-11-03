package board

import ()

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
	if !IsValidSquare(move.To) || !IsValidSquare(move.From) || !IsValidPieceType(*move.Promotion) {
		return false
	}
	return true
}
