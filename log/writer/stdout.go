package writer

import (
	"io"
	"os"
)

type StdoutBuilder struct{}

func NewStdoutBuilder() *StdoutBuilder {
	return &StdoutBuilder{}
}

func (s *StdoutBuilder) Build() (io.Writer, error) {
	return os.Stdout, nil
}
