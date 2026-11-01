package parser

type UCICommand struct {
}

func (c UCICommand) isCommand() {}

func parseUCI(args []string) (Command, error) {
	return UCICommand{}, nil
}
