package board

import (
	"fmt"
)

// BitBoard uses a 64-bit unsigned integer to symbolize a chess board
// For a selected type of pieces
type BitBoard uint64

func (b *BitBoard) SetPiece(square Square) error {
	if !IsValidSquare(square) {
		return fmt.Errorf("SetPiece: square out of bounds")
	}

	// square is an uint8; its value corresponds to the bit shift
	*b |= BitBoard(1) << square
	return nil
}

func (b *BitBoard) HasPieceAtSquare(square Square) bool {
	// square is an uint8; its value corresponds to the bit shift
	pattern := BitBoard(1) << square
	return *b&pattern != 0
}

func (b *BitBoard) HasPieceAtIdx(idx int) (bool, error) {
	if idx > 63 || idx < 0 {
		return false, fmt.Errorf("checkIdx: index out of bounds")
	}

	return *b&(BitBoard(1)<<idx) != 0, nil
}

func (b *BitBoard) Remove(square Square) (BitBoard, error) {
	// square is an uint8; its value corresponds to the bit shift
	pattern := BitBoard(1) << square
	squareTaken := b.HasPieceAtSquare(square)
	if !squareTaken {
		return BitBoard(0), fmt.Errorf("Remove: the square isn't taken")
	}
	return *b &^ pattern, nil
}
