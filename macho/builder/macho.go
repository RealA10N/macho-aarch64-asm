package builder

import (
	"fmt"
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/header"
	writertoutils "github.com/RealA10N/writer-to-utils"
)

type MachoBuilder struct {
	Header   header.MachoHeaderBuilder
	Commands []CommandBuilder
}

func (macho MachoBuilder) allHeadersLen() (n uint64) {
	for _, cmd := range macho.Commands {
		n += cmd.HeaderLen()
	}
	return
}

func (macho MachoBuilder) WriteTo(writer io.Writer) (n int64, err error) {
	var k int64

	headersLen := macho.allHeadersLen()
	ctx := context.CommandContext{
		DataOffset:         header.MachoHeaderSize + headersLen,
		NumOfLoadCommands:  uint32(len(macho.Commands)),
		SizeOfLoadCommands: uint32(headersLen),
	}

	machoHeader := macho.Header.Build(&ctx)
	n, err = writertoutils.BinaryMarshalerAdapter(machoHeader).WriteTo(writer)
	if err != nil {
		return
	}

	for _, cmd := range macho.Commands {
		k, err = cmd.HeaderWriteTo(writer, &ctx)
		n += k
		if err != nil {
			return
		}

		ctx.DataOffset += cmd.DataLen()
	}

	// TODO: we SHOULD  check that the header lengths that the commands have
	// 'committed' to (via GetHeaderSize, GetDataSize) actually equal to the
	// size they write.

	if header.MachoHeaderSize+headersLen != uint64(n) {
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
