package main

import (
	"fmt"
	"log"

	"github.com/tanzor-disco/GOndon/chess/board"
	"github.com/tanzor-disco/GOndon/chess/engine/movegen/attacks/pawn"
	"github.com/tanzor-disco/GOndon/testutils/display"
)

func main() {
	var bitTest board.BitBoard
	bitTest.SetPiece(board.C6)
	test, err := pawn.GenerateAttacksForPawn()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(display.DisplayBitBoard(test[board.B3][board.White]))
	fmt.Println(display.DisplayBitBoard(test[board.B3][board.Black]))
}
