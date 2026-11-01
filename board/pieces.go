package board

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
