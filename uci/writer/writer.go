package writer

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Tanzor-Disco/GOndon/uci/output"
)

type UCIWriter struct {
	writer *bufio.Writer
}

func NewUCIWriter() UCIWriter {
	return UCIWriter{
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (w *UCIWriter) WriteOutput(output output.Output) error {
	defer w.writer.Flush()
	if _, err := w.writer.WriteString(output.GetOutput()); err != nil {
		return fmt.Errorf("Write: %w", err)
	}
	return nil
}
