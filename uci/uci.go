package uci

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EngineName    = "GOndon"
	EngineVersion = "0.0.1"
	EngineAuthor  = "Tanzor-Disco (Pomoyniy Guru)"
)

type CommandType int

const (
	CmdUCI CommandType = iota
	CmdIsReady
	CmdNewGame
	CmdPosition
	CmdGo
	CmdStop
	CmdQuit
)

type UCIHandler struct {
	Scanner *bufio.Scanner
	Writer  *bufio.Writer
}

func NewUCIHandler() *UCIHandler {
	return &UCIHandler{
		Scanner: bufio.NewScanner(os.Stdin),
		Writer:  bufio.NewWriter(os.Stdout),
	}
}

type GoParams struct {
	Depth     *int
	WTime     *int
	BTime     *int
	WInc      *int
	BInc      *int
	MoveTime  *int
	MovesToGo *int
	Infinite  bool
}

type Command struct {
	Type     CommandType
	GoParams *GoParams
}

func parseUCI() Command {
	return Command{
		Type: CmdUCI,
	}
}

func parseIsReady() Command {
	return Command{
		Type: CmdIsReady,
	}
}

func parseQuit() Command {
	return Command{
		Type: CmdQuit,
	}
}

func parseGo(args []string) (Command, error) {
	var currIdx int
	var goParams GoParams
	for currIdx < len(args) {
		switch args[currIdx] {
		case "wtime":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.WTime = &value
			currIdx += 2
		case "btime":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.BTime = &value
			currIdx += 2
		case "winc":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.WInc = &value
			currIdx += 2
		case "binc":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.BInc = &value
			currIdx += 2
		case "movestogo":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.MovesToGo = &value
			currIdx += 2
		case "movetime":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.MoveTime = &value
			currIdx += 2
		case "depth":
			value, err := strconv.Atoi(args[currIdx+1])
			if err != nil {
				return Command{}, fmt.Errorf("parseGo: strconv error: %w", err)
			}
			goParams.Depth = &value
			currIdx += 2
		case "infinite":
			goParams.Infinite = true
			currIdx += 1
		default: // Bad practice: done only to ignore the commands that are not yet implemented
			currIdx += 1
		}

	}
	return Command{
		Type:     CmdGo,
		GoParams: &goParams,
	}, nil
}

func parseCommand(commandString string) (Command, error) {
	if commandString == "" {
		return Command{}, fmt.Errorf("parseCommand: no command found")
	}
	args := strings.Fields(commandString)
	commandName := args[0]
	var commandOutput Command
	switch commandName {
	case "uci":
		commandOutput = parseUCI()
	case "isready":
		commandOutput = parseIsReady()
	case "quit":
		commandOutput = parseQuit()
	case "go":
		var err error
		commandOutput, err = parseGo(args[1:])
		if err != nil {
			return Command{}, fmt.Errorf("parseCommand: parseGo: %w", err)
		}
	default:
		return Command{}, fmt.Errorf("parseCommand: unfamiliar command")
	}
	return commandOutput, nil
}

func (r *UCIHandler) ReadCommand() (Command, error) {
	r.Scanner.Scan()
	commandString := r.Scanner.Text()
	command, err := parseCommand(commandString)
	if err != nil {
		return Command{}, fmt.Errorf("ReadCommand: %w", err)
	}
	return command, nil
}

func (r *UCIHandler) writeUCI() error {
	defer r.Writer.Flush()
	engineLine := fmt.Sprintf("id name %s %s\n", EngineName, EngineVersion)
	authorLine := fmt.Sprintf("id author %s\n", EngineAuthor)
	if _, err := r.Writer.WriteString(engineLine); err != nil {
		return fmt.Errorf("writeUCI: error while writing authorLine: %w", err)
	}
	if _, err := r.Writer.WriteString(authorLine); err != nil {
		return fmt.Errorf("writeUCI: error while writing engineLine: %w", err)
	}

	if _, err := r.Writer.WriteString("uciok\n"); err != nil {
		return fmt.Errorf("writeUCI: error while writing uciok: %w", err)
	}
	return nil
}

func (r *UCIHandler) writeIsReady() error {
	defer r.Writer.Flush()
	if _, err := r.Writer.WriteString("readyok"); err != nil {
		return fmt.Errorf("writeIsReady: error while writing readyok: %w", err)
	}
	return nil
}

func (r *UCIHandler) WriteOutput(command Command) error {
	switch command.Type {
	case CmdUCI:
		if err := r.writeUCI(); err != nil {
			return fmt.Errorf("WriteOutput: %w", err)
		}
	case CmdIsReady:
		if err := r.writeIsReady(); err != nil {
			return fmt.Errorf("WriteOutput: %w", err)
		}
	}
	// TODO: Add more writers to uci
	return nil
}
