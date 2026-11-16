package board_test

import (
	"testing"

	"github.com/Tanzor-Disco/GOndon/chess/board"
)

func TestIsValidSquare(t *testing.T) {
	for value := 0; value <= 255; value++ {
		if got, want := board.IsValidSquare(board.Square(value)), value < 64; got != want {
			t.Errorf("IsValidSquare(%d) = %v, want %v", value, got, want)
		}
	}
}

func TestIsValidPieceType(t *testing.T) {
	for value := 0; value <= 255; value++ {
		if got, want := board.IsValidPieceType(board.PieceType(value)), value < 6; got != want {
			t.Errorf("IsValidPieceType(%d) = %v, want %v", value, got, want)
		}
	}
}

func TestIsValidColor(t *testing.T) {
	for value := 0; value <= 255; value++ {
		if got, want := board.IsValidColor(board.SideColor(value)), value < 2; got != want {
			t.Errorf("IsValidColor(%d) = %v, want %v", value, got, want)
		}
	}
}

func TestPieceTypeMap(t *testing.T) {
	want := map[rune]board.PieceType{
		'p': board.Pawn, 'n': board.Knight, 'b': board.Bishop,
		'r': board.Rook, 'q': board.Queen, 'k': board.King,
	}
	if len(board.PieceTypeMap) != len(want) {
		t.Fatalf("PieceTypeMap has %d entries, want %d", len(board.PieceTypeMap), len(want))
	}
	for symbol, piece := range want {
		if got, ok := board.PieceTypeMap[symbol]; !ok || got != piece {
			t.Errorf("PieceTypeMap[%q] = (%v, %v), want (%v, true)", symbol, got, ok, piece)
		}
	}
}

func TestNewMove(t *testing.T) {
	queen := board.Queen
	for _, promotion := range []*board.PieceType{nil, &queen} {
		got := board.NewMove(board.E7, board.E8, promotion)
		want := board.Move{From: board.E7, To: board.E8, Promotion: promotion}
		if got != want {
			t.Errorf("NewMove = %+v, want %+v", got, want)
		}
	}
}

func TestIsValidMove(t *testing.T) {
	queen, noPiece, invalidPiece := board.Queen, board.NoPiece, board.PieceType(255)
	tests := []struct {
		name string
		move board.Move
		want bool
	}{
		{"quiet", board.Move{From: board.E2, To: board.E4}, true},
		{"board edges", board.Move{From: board.A1, To: board.H8}, true},
		{"promotion", board.Move{From: board.E7, To: board.E8, Promotion: &queen}, true},
		{"empty source sentinel", board.Move{From: board.EmptySquare, To: board.E4}, false},
		{"invalid source", board.Move{From: 255, To: board.E4}, false},
		{"empty target sentinel", board.Move{From: board.E2, To: board.EmptySquare}, false},
		{"invalid target", board.Move{From: board.E2, To: 255}, false},
		{"no piece promotion", board.Move{From: board.E7, To: board.E8, Promotion: &noPiece}, false},
		{"invalid promotion", board.Move{From: board.E7, To: board.E8, Promotion: &invalidPiece}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := board.IsValidMove(tt.move); got != tt.want {
				t.Errorf("IsValidMove(%+v) = %v, want %v", tt.move, got, tt.want)
			}
		})
	}
}
