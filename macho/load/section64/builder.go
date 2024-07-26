package section64

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder"
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

func (builder Section64Builder) HeaderWriteTo(
	writer io.Writer,
	ctx builder.CommandBuilderContext,
) (n int64, err error) {
	section := Section64Header{
		SectionName: builder.SectionName,
		SegmentName: builder.SegmentName,
		Address:     builder.Address,
		Size:        uint64(len(builder.Data)),
		Offset:      uint32(ctx.DataOffset),
		Align:       builder.Align,
		Flags:       builder.Flags,
	}
	return section.WriteTo(writer)
}

func (builder Section64Builder) DataWriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write(builder.Data)
	return int64(n), err
}
