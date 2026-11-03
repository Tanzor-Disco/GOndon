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

func (b *Board) PieceAt(square Square) (PieceType, SideColor, bool) {
	if !IsValidSquare(square) {
		return NoPiece, NoColor, false
	}

	pieceType, ok := b.Black.PieceAt(square)
	if ok {
		return pieceType, Black, true
	}

	pieceType, ok = b.White.PieceAt(square)
	if ok {
		return pieceType, White, true
	}

	return NoPiece, NoColor, false
}

func (b *Board) AddPiece(square Square, piece PieceType, color SideColor) error {
	if !IsValidSquare(square) {
		return fmt.Errorf("AddPiece: invalid square")
	}
	if !IsValidPieceType(piece) {
		return fmt.Errorf("AddPiece: invalid piece type")
	}
	if !IsValidColor(color) {
		return fmt.Errorf("AddPiece: invalid color")
	}

	if color == White {
		b.White.AddPiece(piece, square)
	} else {
		b.Black.AddPiece(piece, square)
	}
	return nil
}

func (b *Board) RemovePiece(square Square) error {
	if !IsValidSquare(square) {
		return fmt.Errorf("RemovePiece: invalid square")
	}

	if _, ok := b.White.PieceAt(square); ok {
		return b.White.RemovePiece(square)
	}
	if _, ok := b.Black.PieceAt(square); ok {
		return b.Black.RemovePiece(square)
	}

	return fmt.Errorf("RemovePiece: no pieces at the square")
}

func (b *Board) GetOccupied() BitBoard {
	return b.White.GetOccupied() | b.Black.GetOccupied()
}

func (b *Board) GetSideOccupied(side SideColor) BitBoard {
	switch side {
	case White:
		return b.White.GetOccupied()
	case Black:
		return b.Black.GetOccupied()
	default:
		return 0
	}
}

func (b *Board) MakeMove(move Move) error {
	fromPieceType, fromColor, ok := b.PieceAt(move.From)
	if !ok {
		return fmt.Errorf("MakeMove: PieceAt: cannot make move: no piece found at the square")
	}
	if _, toColor, ok := b.PieceAt(move.To); ok {
		if toColor == fromColor {
			return fmt.Errorf("MakeMove: cannot move to the square already occupied by the same side")
		}
		if err := b.RemovePiece(move.To); err != nil {
			return fmt.Errorf("MakeMove: %w", err)
		}
	}
	if err := b.RemovePiece(move.From); err != nil {
		return fmt.Errorf("MakeMove: %w", err)
	}

	if move.Promotion != nil {
		if err := b.AddPiece(move.To, *move.Promotion, fromColor); err != nil {
			return fmt.Errorf("MakeMove: %w", err)
		}
	} else {
		if err := b.AddPiece(move.To, fromPieceType, fromColor); err != nil {
			return fmt.Errorf("MakeMove: %w", err)
		}
	}

	return nil
}

// drawing the board
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
