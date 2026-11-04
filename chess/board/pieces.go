package board

// PieceType describes a chess piece type
// using 8-bit unsigned integer
type PieceType uint8

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
	NoPiece
)

func IsValidPieceType(piece PieceType) bool {
	return piece < NoPiece
}
