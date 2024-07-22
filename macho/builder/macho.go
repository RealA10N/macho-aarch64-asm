package builder

import (
	"encoding"
	"fmt"
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/header"
)

type CommandHeaderBuilder interface {
	io.WriterTo
}
type CommandDataBuilder interface {
	io.WriterTo
}

type CommandBuilder interface {
	GetHeaderSize() int64
	SetOffset(offest int64) (CommandHeaderBuilder, CommandDataBuilder)
}

type MachoBuilder struct {
	Header   header.MachoHeader
	Commands []CommandBuilder
}

func WriteTo(obj encoding.BinaryMarshaler, writer io.Writer) (int, error) {
	data, err := obj.MarshalBinary()
	if err != nil {
		return 0, err
	}
	return writer.Write(data)
}

func (macho MachoBuilder) allHeadersSize() (n int64) {
	for _, cmd := range macho.Commands {
		n += cmd.GetHeaderSize()
	}
	return
}

func (macho MachoBuilder) WriteTo(writer io.Writer) (n int64, err error) {
	var k int64

	n, err = macho.Header.WriteTo(writer)
	if err != nil {
		return
	}

	dataBuilders := []CommandDataBuilder{}
	dataOffset := n + macho.allHeadersSize()

	for _, cmd := range macho.Commands {
		header, data := cmd.SetOffset(dataOffset)

		k, err = header.WriteTo(writer)
		n += k
		if err != nil {
			return
		}

		dataBuilders = append(dataBuilders, data)
		dataOffset += k
	}

	if n != dataOffset {
		err = fmt.Errorf("expected macho header %d (actual %d)", dataOffset, n)
		return
	}

	for _, data := range dataBuilders {
		k, err = data.WriteTo(writer)
		n += k
		if err != nil {
			return
		}
	}

	return
}
