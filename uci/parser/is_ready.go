package parser

import ()

type IsReadyCommand struct {
}

func (c IsReadyCommand) isCommand() {

}

func parseIsReady(args []string) (Command, error) {
	return IsReadyCommand{}, nil
}
