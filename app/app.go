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
		if err := UCIHandler.UCIWriter.WriteOutput(output); err != nil {
			return fmt.Errorf("Run: %w", err)
		}
	}
}
