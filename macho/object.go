package macho

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/sections"
	multiwriterto "github.com/RealA10N/multi-writer-to"
)

type Object struct {
	GlobalSection sections.GlobalSection
	TextSection   sections.TextSection
}

func (object Object) WriteTo(writer io.Writer) (n int64, err error) {
	writerTo := multiwriterto.MultiWriterTo(object.GlobalSection, object.TextSection)
	return writerTo.WriteTo(writer)
}
