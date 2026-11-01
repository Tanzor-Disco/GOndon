package parser

import (
	"fmt"
	"strings"
)

func ParseCommand(input string) (Command, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("ParseCommand: empty command")
	}
	command := parts[0]
	var args []string
	if len(parts) > 1 {
		args = parts[1:]
	}
	parseFunc, ok := routes[command]
	if !ok {
		return nil, fmt.Errorf("ParseCommand: invalid command passed")
	}
	output, err := parseFunc(args)
	if err != nil {
		return nil, fmt.Errorf("ParseCommand: %w", err)
	}
	return output, nil
}
