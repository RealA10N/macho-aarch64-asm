package builder

import "io"

type CommandBuilderContext struct {
	DataOffset uint64
}

type CommandBuilder interface {
	GetHeaderLen() uint64
	GetDataLen() uint64
	HeaderWriteTo(writer io.Writer, ctx CommandBuilderContext) (n int64, err error)
	DataWriteTo(writer io.Writer) (n int64, err error)
}
