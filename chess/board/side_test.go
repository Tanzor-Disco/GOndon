package board_test

import (
	"fmt"
	"testing"

	"github.com/Tanzor-Disco/GOndon/chess/board"
)

func TestNewSide(t *testing.T) {
	for _, color := range []board.SideColor{board.White, board.Black} {
		got, err := board.NewSide(color)
		if err != nil {
			t.Fatalf("NewSide(%d): %v", color, err)
		}
		if want := (board.Side{Color: color}); got != want {
			t.Errorf("NewSide(%d) = %+v, want %+v", color, got, want)
		}
	}
	for _, color := range []board.SideColor{board.NoColor, 3, 255} {
		if _, err := board.NewSide(color); err == nil {
			t.Errorf("NewSide(%d) succeeded", color)
		}
	}
}

func TestSidePieceLifecycle(t *testing.T) {
	for _, color := range []board.SideColor{board.White, board.Black} {
		for piece := board.Pawn; piece <= board.King; piece++ {
			t.Run(fmt.Sprintf("color_%d/piece_%d", color, piece), func(t *testing.T) {
				s := board.Side{Color: color}
				for _, square := range []board.Square{board.A1, board.D4, board.H8} {
					if !s.IsEmpty(square) {
						t.Fatalf("IsEmpty(%d) = false before adding", square)
					}
					if got, ok := s.PieceAt(square); got != board.NoPiece || ok {
						t.Fatalf("PieceAt(%d) = (%d, %v) before adding", square, got, ok)
					}
					for range 2 {
						if err := s.AddPiece(piece, square); err != nil {
							t.Fatalf("AddPiece(%d, %d): %v", piece, square, err)
						}
					}
					if got, ok := s.PieceAt(square); got != piece || !ok {
						t.Errorf("PieceAt(%d) = (%d, %v), want (%d, true)", square, got, ok, piece)
					}
					if s.IsEmpty(square) {
						t.Errorf("IsEmpty(%d) = true after adding", square)
					}
				}
				if want := board.BitBoard(0x8000000008000001); s.GetOccupied() != want {
					t.Errorf("GetOccupied() = %#x, want %#x", s.GetOccupied(), want)
				}
				if err := s.RemovePiece(board.D4); err != nil {
					t.Fatalf("RemovePiece(D4): %v", err)
				}
				if got, ok := s.PieceAt(board.D4); got != board.NoPiece || ok || !s.IsEmpty(board.D4) {
					t.Error("D4 is not empty after removal")
				}
				if want := board.BitBoard(0x8000000000000001); s.GetOccupied() != want {
					t.Errorf("GetOccupied() after removal = %#x, want %#x", s.GetOccupied(), want)
				}
				for _, square := range []board.Square{board.A1, board.H8} {
					if err := s.RemovePiece(square); err != nil {
						t.Fatalf("RemovePiece(%d): %v", square, err)
					}
				}
				if want := (board.Side{Color: color}); s != want {
					t.Errorf("side after removing all pieces = %+v, want %+v", s, want)
				}
			})
		}
	}
}

func TestSideMixedPieces(t *testing.T) {
	s := board.Side{Color: board.Black}
	for piece := board.Pawn; piece <= board.King; piece++ {
		if err := s.AddPiece(piece, board.Square(piece)); err != nil {
			t.Fatal(err)
		}
	}
	if got := s.GetOccupied(); got != 0x3f {
		t.Errorf("GetOccupied() = %#x, want 0x3f", got)
	}
	for piece := board.Pawn; piece <= board.King; piece++ {
		if got, ok := s.PieceAt(board.Square(piece)); got != piece || !ok {
			t.Errorf("PieceAt(%d) = (%d, %v), want (%d, true)", piece, got, ok, piece)
		}
	}
}

func TestSideInvalidOperations(t *testing.T) {
	s := board.Side{Color: board.White, Pieces: [6]board.BitBoard{board.Pawn: 1}}
	before := s
	for _, square := range []board.Square{board.EmptySquare, 65, 255} {
		if err := s.AddPiece(board.Queen, square); err == nil {
			t.Errorf("AddPiece(Queen, %d) succeeded", square)
		}
		if err := s.RemovePiece(square); err == nil {
			t.Errorf("RemovePiece(%d) succeeded", square)
		}
		if got, ok := s.PieceAt(square); got != board.NoPiece || ok {
			t.Errorf("PieceAt(%d) = (%d, %v), want (NoPiece, false)", square, got, ok)
		}
		if s.IsEmpty(square) {
			t.Errorf("IsEmpty(%d) = true for an invalid square", square)
		}
	}
	for _, piece := range []board.PieceType{board.NoPiece, 7, 255} {
		if err := s.AddPiece(piece, board.H8); err == nil {
			t.Errorf("AddPiece(%d, H8) succeeded", piece)
		}
	}
	if err := s.RemovePiece(board.H8); err == nil {
		t.Error("RemovePiece on an empty square succeeded")
	}
	if s != before {
		t.Errorf("invalid operations changed the side: got %+v, want %+v", s, before)
	}
}
