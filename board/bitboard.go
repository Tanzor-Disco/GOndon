package board

import (
	"fmt"
	"strconv"
	"strings"
)

const charBase = 97
const period = 8

type BitBoard uint64

func getBitShift(square string) (int, error) {
	if len(square) != 2 {
		return 0, fmt.Errorf("getBitShift: invalid length supplied: wanted 2, got %d", len(square))
	}

	char := square[0]
	if !(char >= 'a' && char <= 'h') {
		return 0, fmt.Errorf("getBitShift: invalid character supplied: %b", char)
	}

	number, err := strconv.Atoi(string(square[1]))
	if err != nil {
		return 0, fmt.Errorf("getBitShift: strconv.Atoi: %w", err)
	}
	shift := int(char) - charBase + period*(number-1)
	return shift, nil
}

func (b *BitBoard) getBits() string {
	return fmt.Sprintf("%064b", *b)
}

func (b *BitBoard) DrawBoard() string {
	bits := b.getBits()
	parts := make([]string, 0, len(bits)+len(bits)/8)
	var counter int
	for _, bit := range bits {
		counter++
		parts = append(parts, string(bit))
		if counter == 8 {
			counter = 0
			parts = append(parts, "\n")
		}
	}
	return strings.Join(parts, "")
}

func (b *BitBoard) Set(square string) error {
	bitShift, err := getBitShift(square)
	if err != nil {
		return fmt.Errorf("Set: %w", err)
	}
	*b |= BitBoard(1) << bitShift
	return nil
}

func (b *BitBoard) CheckPosition(square string) (bool, error) {
	bitShift, err := getBitShift(square)
	if err != nil {
		return false, fmt.Errorf("CheckPosition: %w", err)
	}
	pattern := BitBoard(1) << bitShift
	return *b&pattern != 0, nil
}

func (b *BitBoard) checkIdx(idx int) (bool, error) {
	if idx > 63 || idx < 0 {
		return false, fmt.Errorf("checkIdx: index out of bounds")
	}

	return *b&(BitBoard(1)<<idx) != 0, nil
}

func (b *BitBoard) Remove(square string) (BitBoard, error) {
	bitShift, err := getBitShift(square)
	if err != nil {
		return BitBoard(0), fmt.Errorf("Remove: %w", err)
	}
	pattern := BitBoard(1) << bitShift
	squareTaken, err := b.CheckPosition(square)
	if err != nil {
		return BitBoard(0), fmt.Errorf("Remove: %w", err)
	}
	if !squareTaken {
		return BitBoard(0), fmt.Errorf("Remove: the square isn't taken")
	}
	return *b &^ pattern, nil
}
