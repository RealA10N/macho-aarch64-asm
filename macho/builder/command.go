package builder

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
)

type CommandBuilder interface {
	HeaderLen() uint64
	DataLen() uint64
	HeaderWriteTo(writer io.Writer, ctx *context.CommandContext) (int64, error)
	DataWriteTo(writer io.Writer) (int64, error)
}

// HeaderWriterTo

type headerWriterTo struct {
	cmd *CommandBuilder
	ctx *context.CommandContext
}

func (hdr headerWriterTo) WriteTo(writer io.Writer) (int64, error) {
	return (*hdr.cmd).HeaderWriteTo(writer, hdr.ctx)
}

func HeaderWriterTo(cmd CommandBuilder, ctx *context.CommandContext) io.WriterTo {
	return headerWriterTo{cmd: &cmd, ctx: ctx}
}

// DataWriterTo

type dataWriterTo struct {
	cmd *CommandBuilder
}

func (data dataWriterTo) WriteTo(writer io.Writer) (int64, error) {
	return (*data.cmd).DataWriteTo(writer)
}

func DataWriterTo(cmd CommandBuilder) io.WriterTo {
	return dataWriterTo{cmd: &cmd}
}
