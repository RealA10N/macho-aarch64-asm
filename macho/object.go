package macho

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/sections"
	 "github.com/RealA10N/writer-to-utils"
)

type Object struct {
	GlobalSection sections.GlobalSection
	TextSection   sections.TextSection
}

func (object Object) WriteTo(writer io.Writer) (n int64, err error) {
	writerTo := writertoutils.MultiWriterTo(object.GlobalSection, object.TextSection)
	return writerTo.WriteTo(writer)
}
