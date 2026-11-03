package board

import (
	"fmt"
	"strings"
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

type Side struct {
	Color  SideColor
	Pieces [6]BitBoard
}

func createWhiteSide() Side {
	pawns := BitBoard(255) << BitBoard(A2)

	var bishops BitBoard
	bishops.SetPiece(C1)
	bishops.SetPiece(F1)

	var knights BitBoard
	knights.SetPiece(B1)
	knights.SetPiece(G1)

	var rooks BitBoard
	rooks.SetPiece(A1)
	rooks.SetPiece(H1)

	var queens BitBoard
	queens.SetPiece(D1)

	var king BitBoard
	king.SetPiece(E1)

	var pieces [6]BitBoard
	pieces[Pawn] = pawns
	pieces[Bishop] = bishops
	pieces[Knight] = knights
	pieces[Rook] = rooks
	pieces[Queen] = queens
	pieces[King] = king

	return Side{
		Color:  White,
		Pieces: pieces,
	}
}

func createBlackSide() Side {
	pawns := BitBoard(255) << BitBoard(A7)

	var bishops BitBoard
	bishops.SetPiece(C8)
	bishops.SetPiece(F8)

	var knights BitBoard
	knights.SetPiece(B8)
	knights.SetPiece(G8)

	var rooks BitBoard
	rooks.SetPiece(A8)
	rooks.SetPiece(H8)

	var queens BitBoard
	queens.SetPiece(D8)

	var king BitBoard
	king.SetPiece(E8)

	var pieces [6]BitBoard
	pieces[Pawn] = pawns
	pieces[Bishop] = bishops
	pieces[Knight] = knights
	pieces[Rook] = rooks
	pieces[Queen] = queens
	pieces[King] = king

	return Side{
		Color:  Black,
		Pieces: pieces,
	}
}

func NewSide(color SideColor) Side {
	if color == White {
		return createWhiteSide()
	}
	return createBlackSide()
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

// drawing the board
func fillBoard(board [][]string, tag string, bitBoard BitBoard) error {
	for idx := range 64 {
		filled, err := bitBoard.hasPieceAtIdx(idx)
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
		emptyRow = append(emptyRow, EmptySquareString)
	}

	board := make([][]string, 8)
	for rowIdx := range board {
		curr := make([]string, 8)
		copy(curr, emptyRow)
		board[rowIdx] = curr
	}

	if err := fillBoard(board, "p", s.Pieces[Pawn]); err != nil {
		return "", err
	}
	if err := fillBoard(board, "n", s.Pieces[Knight]); err != nil {
		return "", err
	}
	if err := fillBoard(board, "b", s.Pieces[Bishop]); err != nil {
		return "", err
	}
	if err := fillBoard(board, "r", s.Pieces[Rook]); err != nil {
		return "", err
	}
	if err := fillBoard(board, "q", s.Pieces[Queen]); err != nil {
		return "", err
	}
	if err := fillBoard(board, "k", s.Pieces[King]); err != nil {
		return "", err
	}

	parts := make([]string, 0, len(board))
	for _, row := range board {
		parts = append(parts, strings.Join(row, ""))
	}
	return strings.Join(parts, "\n"), nil
}
