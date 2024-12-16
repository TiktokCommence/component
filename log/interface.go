package log

import "io"

type WriterBuilder interface {
	Build() (io.Writer, error)
}
