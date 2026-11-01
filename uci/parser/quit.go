package parser

type QuitCommand struct{}

func (c QuitCommand) isCommand() {}

func parseQuit(args []string) (Command, error) {
	return QuitCommand{}, nil
}
