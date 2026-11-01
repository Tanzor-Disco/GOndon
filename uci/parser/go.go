package parser

import (
	"fmt"
	"strconv"
)

type GoCommand struct {
	Wtime     *int
	Btime     *int
	Winc      *int
	Binc      *int
	MovesToGo *int
	MoveTime  *int
}

func (c GoCommand) isCommand() {}

func parseArg(args []string, currIdx int, name string) (int, error) {
	if currIdx+1 == len(args) {
		return 0, fmt.Errorf("ParseArg: no argument passed to %v", name)
	}

	amount, err := strconv.Atoi(args[currIdx+1])
	if err != nil {
		return 0, fmt.Errorf("ParseArg: invalid amount passed as an argument to %v; err: %w", name, err)
	}

	if amount < 0 {
		return 0, fmt.Errorf("ParseArg: invalid amount value passed as an argument to %v", name)
	}
	return amount, nil
}

func parseGo(args []string) (Command, error) {
	var command GoCommand
	var currIdx int
	for currIdx < len(args) {
		arg := args[currIdx]
		switch arg {
		case "wtime":
			amount, err := parseArg(args, currIdx, arg)
			if err != nil {
				return GoCommand{}, fmt.Errorf("ParseGo: %w", err)
			}
			command.Wtime = &amount
			currIdx += 2

		case "btime":
			amount, err := parseArg(args, currIdx, arg)
			if err != nil {
				return GoCommand{}, fmt.Errorf("ParseGo: %w", err)
			}
			command.Btime = &amount
			currIdx += 2

		case "winc":
			amount, err := parseArg(args, currIdx, arg)
			if err != nil {
				return GoCommand{}, fmt.Errorf("ParseGo: %w", err)
			}
			command.Winc = &amount
			currIdx += 2

		case "binc":
			amount, err := parseArg(args, currIdx, arg)
			if err != nil {
				return GoCommand{}, fmt.Errorf("ParseGo: %w", err)
			}
			command.Binc = &amount
			currIdx += 2

		case "movestogo":
			amount, err := parseArg(args, currIdx, arg)
			if err != nil {
				return GoCommand{}, fmt.Errorf("ParseGo: %w", err)
			}
			command.MovesToGo = &amount
			currIdx += 2

		case "movetime":
			amount, err := parseArg(args, currIdx, arg)
			if err != nil {
				return GoCommand{}, fmt.Errorf("ParseGo: %w", err)
			}
			command.MoveTime = &amount
			currIdx += 2

		default:
			return GoCommand{}, fmt.Errorf("ParseGo: invalid argument passed")
		}
	}
	return command, nil
}
