package parser

import ()

type newGameCommand struct{}

func (c newGameCommand) isCommand() {}

func parseNewGame(args []string) (Command, error) {
	return newGameCommand{}, nil
}
