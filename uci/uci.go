package uci

import (
	"fmt"

	"github.com/Tanzor-Disco/GOndon/uci/parser"
	"github.com/Tanzor-Disco/GOndon/uci/writer"
)

type UCIHandler struct {
	UCIScanner UCIScanner
	UCIWriter  writer.UCIWriter
}

func NewUCIHandler() UCIHandler {
	return UCIHandler{
		UCIScanner: NewUCIScanner(),
		UCIWriter:  writer.NewUCIWriter(),
	}
}

func (h *UCIHandler) GetCommand() (parser.Command, error) {
	input, err := h.UCIScanner.ScanCommand()
	if err != nil {
		return nil, fmt.Errorf("GetCommand: %w", err)
	}
	command, err := parser.ParseCommand(input)
	if err != nil {
		return nil, fmt.Errorf("GetCommand: %w", err)
	}
	return command, nil
}
