package engine

import (
	"fmt"
	"github.com/Tanzor-Disco/GOndon/uci"
)

func Run() {
	UCIHandler := uci.NewUCIHandler()

	for {
		command, err := UCIHandler.ReadCommand()
		if err != nil {
			fmt.Printf("Run: Invalid command entered: %v", err)
			return
		}

		if command.Type == uci.CmdQuit {
			return
		}

		if err := UCIHandler.WriteOutput(command); err != nil {
			fmt.Printf("Run: error processing the output: %v", err)
			return
		}
	}
}
