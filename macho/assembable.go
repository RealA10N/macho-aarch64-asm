package macho

import "io"

type Assembable interface {
	WriteAssembly(io.Writer) (n int, err error)
}
