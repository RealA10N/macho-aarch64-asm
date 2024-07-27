package section64

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	writertoutils "github.com/RealA10N/writer-to-utils"
)

type Section64Builder struct {
	SectionName [16]byte
	SegmentName [16]byte
	Data        []byte
	Address     uint64
	Align       uint32
	Flags       Section64Flags
	// TODO: support custom section size (for BSS, etc.)
	// TODO: support Relocations
}

func (builder Section64Builder) HeaderLen() uint64 {
	return Section64HeaderSize
}

func (builder Section64Builder) DataLen() uint64 {
	return uint64(len(builder.Data))
}

func (builder Section64Builder) Build(
	ctx *context.CommandContext,
) Section64Header {
	return Section64Header{
		SectionName: builder.SectionName,
		SegmentName: builder.SegmentName,
		Address:     builder.Address,
		Size:        uint64(len(builder.Data)),
		Offset:      uint32(ctx.DataOffset),
		Align:       builder.Align,
		Flags:       builder.Flags,
	}
}

func (builder Section64Builder) HeaderWriteTo(
	writer io.Writer,
	ctx *context.CommandContext,
) (int64, error) {
	section := builder.Build(ctx)
	writerTo := writertoutils.BinaryMarshalerAdapter(section)
	return writerTo.WriteTo(writer)
}

func (builder Section64Builder) DataWriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write(builder.Data)
	return int64(n), err
}
