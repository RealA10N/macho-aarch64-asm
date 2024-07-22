package builder

import (
	"fmt"
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/header"
)

type MachoBuilder struct {
	Header   header.MachoHeader
	Commands []CommandBuilder
}

func (macho MachoBuilder) allHeadersLen() (n uint64) {
	for _, cmd := range macho.Commands {
		n += cmd.GetHeaderLen()
	}
	return
}

func (macho MachoBuilder) WriteTo(writer io.Writer) (n int64, err error) {
	var k int64

	n, err = macho.Header.WriteTo(writer)
	if err != nil {
		return
	}

	ctx := CommandBuilderContext{DataOffset: uint64(n) + macho.allHeadersLen()}

	for _, cmd := range macho.Commands {
		k, err = cmd.HeaderWriteTo(writer, ctx)
		n += k
		if err != nil {
			return
		}

		ctx.DataOffset += cmd.GetDataLen()
	}

	// TODO: we SHOULD  check that the header lengths that the commands have
	// 'committed' to (via GetHeaderSize, GetDataSize) actually equal to the
	// size they write.

	if ctx.DataOffset != uint64(n) {
		err = fmt.Errorf("expected headers size %d (actually %d)", ctx.DataOffset, n)
		return
	}

	for _, cmd := range macho.Commands {
		k, err = cmd.DataWriteTo(writer)
		n += k
		if err != nil {
			return
		}
	}

	return
}
