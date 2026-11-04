package parser

import ()

type ParserFunc func(args []string) (Command, error)

// routes is a map used to determine which parser function to use in parser.go
var routes = map[string]ParserFunc{
	"uci":        parseUCI,
	"isready":    parseIsReady,
	"ucinewgame": parseNewGame,
	"go":         parseGo,
	"position":   parsePosition,
}
