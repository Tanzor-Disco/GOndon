package app

import (
	"github.com/Tanzor-Disco/GOndon/uci"
	"github.com/Tanzor-Disco/GOndon/uci/output"
	"github.com/Tanzor-Disco/GOndon/uci/parser"

	"fmt"
)

func handleCommand(command parser.Command) output.Output {
	switch command.(type) {
	case parser.UCICommand:
		return output.UCIOutput{}
	case parser.IsReadyCommand:
		return output.ReadyOutput{}
	case parser.PositionCommand:
	}
	return nil
}

func Run() error {
	UCIHandler := uci.NewUCIHandler()

	for {
		command, err := UCIHandler.GetCommand()
		if err != nil {
			return fmt.Errorf("Run: %w", err)
		}

		output := handleCommand(command)
		if output == nil {
			return fmt.Errorf("handleCommand: nil output")
		}

		if err := UCIHandler.UCIWriter.WriteOutput(output); err != nil {
			return fmt.Errorf("Run: %w", err)
		}
	}
}
