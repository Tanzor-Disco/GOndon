package board_test

import (
	"fmt"
	"testing"

	"github.com/tanzor-disco/GOndon/chess/board"
)

func TestBitBoardSetAndQuery(t *testing.T) {
	for square := board.A1; square <= board.H8; square++ {
		t.Run(fmt.Sprint(square), func(t *testing.T) {
			var b board.BitBoard
			for range 2 {
				if err := b.SetPiece(square); err != nil {
					t.Fatalf("SetPiece(%d): %v", square, err)
				}
			}
			if want := board.BitBoard(1) << square; b != want {
				t.Fatalf("SetPiece(%d) = %#x, want %#x", square, b, want)
			}
			for idx := 0; idx < 64; idx++ {
				want := idx == int(square)
				if got := b.HasPieceAtSquare(board.Square(idx)); got != want {
					t.Errorf("HasPieceAtSquare(%d) = %v, want %v", idx, got, want)
				}
				got, err := b.HasPieceAtIdx(idx)
				if err != nil || got != want {
					t.Errorf("HasPieceAtIdx(%d) = (%v, %v), want (%v, nil)", idx, got, err, want)
				}
			}
		})
	}
}

func TestBitBoardInvalidInputs(t *testing.T) {
	for _, square := range []board.Square{board.EmptySquare, 65, 255} {
		b := board.BitBoard(0x8000000000000001)
		before := b
		if err := b.SetPiece(square); err == nil {
			t.Errorf("SetPiece(%d) succeeded", square)
		}
		if b != before {
			t.Errorf("SetPiece(%d) changed the bitboard", square)
		}
		if b.HasPieceAtSquare(square) {
			t.Errorf("HasPieceAtSquare(%d) = true", square)
		}
	}
	var b board.BitBoard
	for _, idx := range []int{-1, -100, 64, 255} {
		if got, err := b.HasPieceAtIdx(idx); err == nil || got {
			t.Errorf("HasPieceAtIdx(%d) = (%v, %v), want (false, error)", idx, got, err)
		}
	}
}

func TestBitBoardRemove(t *testing.T) {
	for _, square := range []board.Square{board.A1, board.D4, board.H8} {
		t.Run(fmt.Sprint(square), func(t *testing.T) {
			b := ^board.BitBoard(0)
			before := b
			got, err := b.Remove(square)
			if err != nil {
				t.Fatalf("Remove(%d): %v", square, err)
			}
			if want := before &^ (board.BitBoard(1) << square); got != want {
				t.Errorf("Remove(%d) = %#x, want %#x", square, got, want)
			}
			// Remove returns a new bitboard; the receiver keeps its original value.
			if b != before {
				t.Error("Remove changed the receiver")
			}
		})
	}
	for _, square := range []board.Square{board.A1, board.H8, board.EmptySquare, 255} {
		b := board.BitBoard(1) << board.D4
		before := b
		if got, err := b.Remove(square); err == nil || got != 0 {
			t.Errorf("Remove(%d) = (%#x, %v), want (0, error)", square, got, err)
		}
		if b != before {
			t.Errorf("failed Remove(%d) changed the receiver", square)
		}
	}
}
