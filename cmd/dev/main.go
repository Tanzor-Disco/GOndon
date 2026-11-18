package main

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/chess/board"
	"github.com/Tanzor-Disco/GOndon/chess/engine/movegen/attacks/rook"
	"github.com/Tanzor-Disco/GOndon/testutils/display"
	"log"
)

func main() {
	var bitTest board.BitBoard
	bitTest.SetPiece(board.E5)
	test, err := rook.GenerateAttacksForRookFromSquare(bitTest, board.E1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(display.DisplayBitBoard(test))
}
