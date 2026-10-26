package main

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/board"
	_ "github.com/Tanzor-Disco/GOndon/engine"
)

func main() {
	// engine.Run()
	test := board.NewBoard()
	fmt.Println(test.DrawBoard())
}
