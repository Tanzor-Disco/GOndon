package app

import (
	"github.com/Tanzor-Disco/GOndon/chess/builder"
	"github.com/Tanzor-Disco/GOndon/uci"
	"github.com/Tanzor-Disco/GOndon/uci/output"
	"github.com/Tanzor-Disco/GOndon/uci/parser"

	"fmt"
)

func handleCommand(command parser.Command) (output.Output, error) {
	switch cmd := command.(type) {
	case parser.UCICommand:
		return output.UCIOutput{}, nil
	case parser.IsReadyCommand:
		return output.ReadyOutput{}, nil
	case parser.PositionCommand:
		_, err := builder.BuildPosition(cmd.Position)
		if err != nil {
			return nil, fmt.Errorf("handleCommand: %w", err)
		}

	}

	return nil, fmt.Errorf("handleCommand: fell out of switch")
}

func Run() error {
	UCIHandler := uci.NewUCIHandler()

	for {
		command, err := UCIHandler.GetCommand()
		if err != nil {
			return fmt.Errorf("Run: %w", err)
		}

		output, err := handleCommand(command)
		if err != nil {
			return fmt.Errorf("handleCommand: %w", err)
		}

		if err := UCIHandler.UCIWriter.WriteOutput(output); err != nil {
			return fmt.Errorf("Run: %w", err)
		}
	}
}
