package board

import (
	"fmt"
)

type SideColor uint8

const (
	White SideColor = iota
	Black
	NoColor

	EmptySquareString = "."
)

func IsValidColor(color SideColor) bool {
	return color < NoColor
}

// Side represents one player's pieces using bitboards.
//
// Pieces stores one bitboard per piece type. The array must be accessed
// using PieceType constants as indices:
//
//	Pieces[Pawn]
//	Pieces[Knight]
//	Pieces[Bishop]
//	Pieces[Rook]
//	Pieces[Queen]
//	Pieces[King]
//
// NoPiece is not a valid array index and must not be used to access Pieces.
type Side struct {
	Color  SideColor
	Pieces [6]BitBoard
}

func NewSide(color SideColor) (Side, error) {
	if color != Black && color != White {
		return Side{}, fmt.Errorf("NewSide: invalid color: %v", color)
	}

	return Side{
		Color: color,
	}, nil
}

func (s *Side) PieceAt(square Square) (PieceType, bool) {
	if !IsValidSquare(square) {
		return NoPiece, false
	}

	// square is an uint8; its value corresponds to the bit shift
	mask := BitBoard(1) << BitBoard(square)
	for pieceNum, bitboard := range s.Pieces {
		if bitboard&mask != 0 {
			return PieceType(pieceNum), true
		}
	}
	return NoPiece, false
}

func (s *Side) AddPiece(piece PieceType, square Square) error {
	if !IsValidSquare(square) {
		return fmt.Errorf("AddPiece: adding to invalid square")
	}
	if !IsValidPieceType(piece) {
		return fmt.Errorf("AddPiece: invalid piece type")
	}

	// square is an uint8; its value corresponds to the bit shift
	mask := BitBoard(1) << BitBoard(square)
	s.Pieces[piece] |= mask
	return nil
}

func (s *Side) RemovePiece(square Square) error {
	if !IsValidSquare(square) {
		return fmt.Errorf("RemovePiece: removing from invalid square")
	}

	piece, ok := s.PieceAt(square)
	if !ok {
		return fmt.Errorf("RemovePiece: PieceAt: no piece at the square")
	}

	// square is an uint8; its value corresponds to the bit shift
	mask := BitBoard(1) << BitBoard(square)
	s.Pieces[piece] &^= mask
	return nil
}

func (s *Side) GetOccupied() BitBoard {
	var bitboard BitBoard
	for _, pieceBoard := range s.Pieces {
		bitboard |= pieceBoard
	}
	return bitboard
}

func (s *Side) IsEmpty(square Square) bool {
	if !IsValidSquare(square) {
		return false
	}

	// square is an uint8; its value corresponds to the bit shift
	mask := BitBoard(1) << BitBoard(square)
	for _, bitboard := range s.Pieces {
		if bitboard&mask != 0 {
			return false
		}
	}
	return true
}
