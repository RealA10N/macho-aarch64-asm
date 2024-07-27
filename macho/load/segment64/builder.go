package segment64

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder"
	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/load"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/section64"
	writertoutils "github.com/RealA10N/writer-to-utils"
)

type Segment64Builder struct {
	SegmentName          [16]byte
	Sections             []section64.Section64Builder
	VirtualMemoryAddress uint64
	VirtualMemorySize    uint64
	MaxProtections       VirtualMemoryProtection
	InitialProtections   VirtualMemoryProtection
	Flags                Segment64Flag
}

func (segment Segment64Builder) HeaderLen() uint64 {
	var sum uint64 = Segment64HeaderSize
	for _, section := range segment.Sections {
		sum += section.HeaderLen()
	}
	return sum
}

func (segment Segment64Builder) DataLen() (len uint64) {
	for _, section := range segment.Sections {
		len += section.DataLen()
	}
	return
}

func (segment Segment64Builder) Build(
	ctx *context.CommandContext,
) Segment64Header {
	return Segment64Header{
		CommandType:          load.Segment64,
		CommandSize:          uint32(segment.HeaderLen()),
		SegmentName:          segment.SegmentName,
		VirtualMemoryAddress: segment.VirtualMemoryAddress,
		VirtualMemorySize:    segment.VirtualMemorySize,
		FileOffset:           ctx.DataOffset,
		FileSize:             segment.DataLen(),
		MaxProtections:       segment.MaxProtections,
		InitialProtections:   segment.InitialProtections,
		NumberOfSections:     uint32(len(segment.Sections)),
		Flags:                segment.Flags,
	}
}

func (segment Segment64Builder) HeaderWriteTo(
	writer io.Writer,
	ctx *context.CommandContext,
) (int64, error) {
	segmentHeader := segment.Build(ctx)

	writerTos := []io.WriterTo{}
	writerTos = append(writerTos, writertoutils.BinaryMarshalerAdapter(segmentHeader))
	for _, section := range segment.Sections {
		sectionWriterTo := builder.HeaderWriterTo(section, ctx)
		writerTos = append(writerTos, sectionWriterTo)
	}

	multiWriter := writertoutils.MultiWriterTo(writerTos...)
	return multiWriter.WriteTo(writer)
}

func (segment Segment64Builder) DataWriteTo(writer io.Writer) (n int64, err error) {
	writerTos := []io.WriterTo{}
	for _, section := range segment.Sections {
		sectionWriterTo := builder.DataWriterTo(section)
		writerTos = append(writerTos, sectionWriterTo)
	}

	multiWriterTo := writertoutils.MultiWriterTo(writerTos...)
	return multiWriterTo.WriteTo(writer)
}
