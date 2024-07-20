package header_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/header"
	"github.com/stretchr/testify/assert"
)

func TestHeaderMarshalBinary(t *testing.T) {
	expected := []byte{
		0xCF, 0xFA, 0xED, 0xFE, 0x0C, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00, 0x18, 0x01, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	header := header.MachoHeader{
		Magic:              header.Magic64Bit,
		CpuType:            header.ARM | header.ArchABI64,
		CpuSubType:         header.AllArmProcessors,
		FileType:           header.Object,
		NumOfLoadCommands:  4,
		SizeOfLoadCommands: 280,
		Flags:              0,
	}

	data, err := header.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, expected, data)
}
