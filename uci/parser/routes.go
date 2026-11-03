package parser

import ()

type ParserFunc func(args []string) (Command, error)

var routes = map[string]ParserFunc{
	"uci":        parseUCI,
	"isready":    parseIsReady,
	"ucinewgame": parseNewGame,
	"go":         parseGo,
	"position":   parsePosition,
}
