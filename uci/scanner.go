package uci

import (
	"bufio"
	"fmt"
	"os"
)

type UCIScanner struct {
	scanner *bufio.Scanner
}

func NewUCIScanner() UCIScanner {
	return UCIScanner{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (s *UCIScanner) ScanCommand() (string, error) {
	if !s.scanner.Scan() {
		return "", fmt.Errorf("ScanCommand: %w", s.scanner.Err())
	}
	return s.scanner.Text(), nil
}
