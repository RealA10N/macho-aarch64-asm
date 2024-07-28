package header_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/header"
	"github.com/stretchr/testify/assert"
)

func TestMachoHeaderBuilderExpectedBinary(t *testing.T) {
	builder := header.MachoHeaderBuilder{
		Magic:      header.Magic64Bit,
		CpuType:    header.Arm64CpuType,
		CpuSubType: header.AllArmProcessors,
		FileType:   header.Object,
		Flags:      header.SubsectionsViaSymbols,
	}

	ctx := context.CommandContext{
		NumOfLoadCommands:  4,
		SizeOfLoadCommands: 360,
	}

	header := builder.Build(&ctx)

	expectedHeader := []byte{
		0xCF, 0xFA, 0xED, 0xFE, 0x0C, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00, 0x68, 0x01, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	got, err := header.MarshalBinary()
	assert.NoError(t, err)
	assert.EqualValues(t, expectedHeader, got)
}
