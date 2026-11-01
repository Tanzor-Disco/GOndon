package main

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/board"
	_ "github.com/Tanzor-Disco/GOndon/engine"
)

func main() {
	// engine.Run()
	test := board.NewBoard()
	err := test.RemovePiece(board.A1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test.DrawBoard())
}
