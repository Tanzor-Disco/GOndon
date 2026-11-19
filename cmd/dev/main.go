package main

import (
	"fmt"
	"log"

	"github.com/Tanzor-Disco/GOndon/chess/board"
	"github.com/Tanzor-Disco/GOndon/chess/engine/movegen/attacks/bishop"
	"github.com/Tanzor-Disco/GOndon/testutils/display"
)

func main() {
	var bitTest board.BitBoard
	bitTest.SetPiece(board.C6)
	test, err := bishop.GenerateAttacksForBishopFromSquare(bitTest, board.D5)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(display.DisplayBitBoard(test))
}
