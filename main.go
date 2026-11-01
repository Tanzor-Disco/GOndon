package main

import (
	"github.com/Tanzor-Disco/GOndon/app"
	_ "github.com/Tanzor-Disco/GOndon/board"
	_ "github.com/Tanzor-Disco/GOndon/engine"
	"log"
)

func main() {
	log.Fatal(app.Run())
}
