package board_test

import (
	"fmt"
	"testing"

	"github.com/tanzor-disco/GOndon/chess/board"
)

func newTestBoard(t *testing.T) board.Board {
	t.Helper()
	b, err := board.NewBoard()
	if err != nil {
		t.Fatalf("NewBoard: %v", err)
	}
	return b
}

func addTestPiece(t *testing.T, b *board.Board, square board.Square, piece board.PieceType, color board.SideColor) {
	t.Helper()
	if err := b.AddPiece(square, piece, color); err != nil {
		t.Fatalf("AddPiece(%d, %d, %d): %v", square, piece, color, err)
	}
}

func assertPieceAt(t *testing.T, b *board.Board, square board.Square, piece board.PieceType, color board.SideColor) {
	t.Helper()
	gotPiece, gotColor, ok := b.PieceAt(square)
	if gotPiece != piece || gotColor != color || ok != (piece != board.NoPiece) {
		t.Errorf("PieceAt(%d) = (%d, %d, %v), want (%d, %d, %v)",
			square, gotPiece, gotColor, ok, piece, color, piece != board.NoPiece)
	}
}

func TestNewBoard(t *testing.T) {
	b := newTestBoard(t)
	want := board.Board{White: board.Side{Color: board.White}, Black: board.Side{Color: board.Black}}
	if b != want {
		t.Errorf("NewBoard() = %+v, want %+v", b, want)
	}
	if got := b.GetOccupied(); got != 0 {
		t.Errorf("GetOccupied() = %#x on a new board", got)
	}
	for square := board.A1; square <= board.H8; square++ {
		assertPieceAt(t, &b, square, board.NoPiece, board.NoColor)
	}
}

func TestBoardPieceLifecycle(t *testing.T) {
	for _, color := range []board.SideColor{board.White, board.Black} {
		for piece := board.Pawn; piece <= board.King; piece++ {
			t.Run(fmt.Sprintf("color_%d/piece_%d", color, piece), func(t *testing.T) {
				b := newTestBoard(t)
				before := b
				for _, square := range []board.Square{board.A1, board.D4, board.H8} {
					addTestPiece(t, &b, square, piece, color)
					assertPieceAt(t, &b, square, piece, color)
				}
				if want := board.BitBoard(0x8000000008000001); b.GetOccupied() != want {
					t.Errorf("GetOccupied() = %#x, want %#x", b.GetOccupied(), want)
				}
				for _, square := range []board.Square{board.A1, board.D4, board.H8} {
					if err := b.RemovePiece(square); err != nil {
						t.Fatalf("RemovePiece(%d): %v", square, err)
					}
					assertPieceAt(t, &b, square, board.NoPiece, board.NoColor)
				}
				if b != before {
					t.Errorf("board after removing all pieces = %+v, want %+v", b, before)
				}
			})
		}
	}
}

func TestBoardOccupied(t *testing.T) {
	b := newTestBoard(t)
	addTestPiece(t, &b, board.A1, board.Rook, board.White)
	addTestPiece(t, &b, board.E2, board.Pawn, board.White)
	addTestPiece(t, &b, board.E7, board.Pawn, board.Black)
	addTestPiece(t, &b, board.H8, board.Rook, board.Black)
	for _, tt := range []struct {
		color board.SideColor
		want  board.BitBoard
	}{
		{board.White, 0x0000000000001001},
		{board.Black, 0x8010000000000000},
		{board.NoColor, 0},
		{255, 0},
	} {
		if got := b.GetSideOccupied(tt.color); got != tt.want {
			t.Errorf("GetSideOccupied(%d) = %#x, want %#x", tt.color, got, tt.want)
		}
	}
	if got, want := b.GetOccupied(), board.BitBoard(0x8010000000001001); got != want {
		t.Errorf("GetOccupied() = %#x, want %#x", got, want)
	}
	if err := b.RemovePiece(board.E2); err != nil {
		t.Fatal(err)
	}
	if got, want := b.GetOccupied(), board.BitBoard(0x8010000000000001); got != want {
		t.Errorf("GetOccupied() after removal = %#x, want %#x", got, want)
	}
	assertPieceAt(t, &b, board.A1, board.Rook, board.White)
	assertPieceAt(t, &b, board.E7, board.Pawn, board.Black)
	assertPieceAt(t, &b, board.H8, board.Rook, board.Black)
}

func TestBoardInvalidOperations(t *testing.T) {
	b := newTestBoard(t)
	addTestPiece(t, &b, board.A1, board.King, board.White)
	addTestPiece(t, &b, board.H8, board.King, board.Black)
	before := b
	for _, square := range []board.Square{board.EmptySquare, 65, 255} {
		if err := b.AddPiece(square, board.Pawn, board.White); err == nil {
			t.Errorf("AddPiece at invalid square %d succeeded", square)
		}
		if err := b.RemovePiece(square); err == nil {
			t.Errorf("RemovePiece at invalid square %d succeeded", square)
		}
		assertPieceAt(t, &b, square, board.NoPiece, board.NoColor)
	}
	for _, piece := range []board.PieceType{board.NoPiece, 7, 255} {
		if err := b.AddPiece(board.E4, piece, board.White); err == nil {
			t.Errorf("AddPiece with invalid type %d succeeded", piece)
		}
	}
	for _, color := range []board.SideColor{board.NoColor, 3, 255} {
		if err := b.AddPiece(board.E4, board.Pawn, color); err == nil {
			t.Errorf("AddPiece with invalid color %d succeeded", color)
		}
	}
	if err := b.RemovePiece(board.E4); err == nil {
		t.Error("RemovePiece on an empty square succeeded")
	}
	if b != before {
		t.Errorf("invalid operations changed the board: got %+v, want %+v", b, before)
	}
}

func TestBoardMakeMove(t *testing.T) {
	for _, color := range []board.SideColor{board.White, board.Black} {
		for piece := board.Pawn; piece <= board.King; piece++ {
			t.Run(fmt.Sprintf("color_%d/piece_%d", color, piece), func(t *testing.T) {
				b := newTestBoard(t)
				addTestPiece(t, &b, board.E2, piece, color)
				addTestPiece(t, &b, board.A1, board.Rook, board.White)
				addTestPiece(t, &b, board.H8, board.Rook, board.Black)
				want := newTestBoard(t)
				addTestPiece(t, &want, board.E4, piece, color)
				addTestPiece(t, &want, board.A1, board.Rook, board.White)
				addTestPiece(t, &want, board.H8, board.Rook, board.Black)
				// MakeMove updates placement; movement legality belongs to the caller.
				if err := b.MakeMove(board.NewMove(board.E2, board.E4, nil)); err != nil {
					t.Fatalf("MakeMove: %v", err)
				}
				if b != want {
					t.Errorf("board after move = %+v, want %+v", b, want)
				}
			})
		}
	}
}

func TestBoardMakeMoveCapture(t *testing.T) {
	for _, color := range []board.SideColor{board.White, board.Black} {
		opponent := board.Black
		if color == board.Black {
			opponent = board.White
		}
		for captured := board.Pawn; captured <= board.King; captured++ {
			t.Run(fmt.Sprintf("color_%d/captured_%d", color, captured), func(t *testing.T) {
				b := newTestBoard(t)
				addTestPiece(t, &b, board.D4, board.Queen, color)
				addTestPiece(t, &b, board.D5, captured, opponent)
				want := newTestBoard(t)
				addTestPiece(t, &want, board.D5, board.Queen, color)
				if err := b.MakeMove(board.NewMove(board.D4, board.D5, nil)); err != nil {
					t.Fatalf("MakeMove: %v", err)
				}
				if b != want {
					t.Errorf("board after capture = %+v, want %+v", b, want)
				}
			})
		}
	}
}

func TestBoardMakeMovePromotion(t *testing.T) {
	for _, color := range []board.SideColor{board.White, board.Black} {
		from, to, captureTo, opponent := board.E7, board.E8, board.F8, board.Black
		if color == board.Black {
			from, to, captureTo, opponent = board.E2, board.E1, board.F1, board.White
		}
		for _, promotion := range []board.PieceType{board.Queen, board.Rook, board.Bishop, board.Knight} {
			for _, capture := range []bool{false, true} {
				t.Run(fmt.Sprintf("color_%d/promotion_%d/capture_%v", color, promotion, capture), func(t *testing.T) {
					b := newTestBoard(t)
					addTestPiece(t, &b, from, board.Pawn, color)
					target := to
					if capture {
						target = captureTo
						addTestPiece(t, &b, target, board.Rook, opponent)
					}
					want := newTestBoard(t)
					addTestPiece(t, &want, target, promotion, color)
					if err := b.MakeMove(board.NewMove(from, target, &promotion)); err != nil {
						t.Fatalf("MakeMove: %v", err)
					}
					if b != want {
						t.Errorf("board after promotion = %+v, want %+v", b, want)
					}
				})
			}
		}
	}
}

func TestBoardMakeMoveRejected(t *testing.T) {
	tests := []struct {
		name string
		move board.Move
	}{
		{"empty source", board.Move{From: board.E3, To: board.E4}},
		{"invalid source", board.Move{From: board.EmptySquare, To: board.E4}},
		{"out of bounds source", board.Move{From: 255, To: board.E4}},
		{"friendly target", board.Move{From: board.E2, To: board.D2}},
		{"same square", board.Move{From: board.E2, To: board.E2}},
	}
	for _, color := range []board.SideColor{board.White, board.Black} {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("color_%d/%s", color, tt.name), func(t *testing.T) {
				b := newTestBoard(t)
				addTestPiece(t, &b, board.E2, board.Pawn, color)
				addTestPiece(t, &b, board.D2, board.Knight, color)
				before := b
				if err := b.MakeMove(tt.move); err == nil {
					t.Fatal("MakeMove succeeded, want error")
				}
				if b != before {
					t.Errorf("rejected move changed the board: got %+v, want %+v", b, before)
				}
			})
		}
	}
}

func TestBoardMakeMoveInvalidTargetOrPromotion(t *testing.T) {
	noPiece, invalidPiece := board.NoPiece, board.PieceType(255)
	tests := []struct {
		name string
		move board.Move
	}{
		{"empty target sentinel", board.Move{From: board.E2, To: board.EmptySquare}},
		{"out of bounds target", board.Move{From: board.E2, To: 255}},
		{"no piece promotion", board.Move{From: board.E2, To: board.E4, Promotion: &noPiece}},
		{"invalid promotion", board.Move{From: board.E2, To: board.E4, Promotion: &invalidPiece}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := newTestBoard(t)
			addTestPiece(t, &b, board.E2, board.Pawn, board.White)
			if err := b.MakeMove(tt.move); err == nil {
				t.Error("MakeMove succeeded, want error")
			}
		})
	}
}
