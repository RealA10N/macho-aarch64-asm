package builder_test

import (
	"bytes"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder"
	"github.com/RealA10N/macho-aarch64-asm/macho/header"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/section64"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/segment64"
	"github.com/stretchr/testify/assert"
)

func TestObjectBuildObjectFile(t *testing.T) {
	headerBuilder := header.MachoHeaderBuilder{
		Magic:      header.Magic64Bit,
		CpuType:    header.Arm64CpuType,
		CpuSubType: header.AllArmProcessors,
		FileType:   header.Object,
	}

	data := []byte{0x00, 0x00, 0x01, 0x8B, 0xC0, 0x03, 0x5F, 0xD6}

	sectionBuilder := section64.Section64Builder{
		SectionName: [16]byte{'_', '_', 't', 'e', 'x', 't', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		SegmentName: [16]byte{'_', '_', 'T', 'E', 'X', 'T', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Data:        data,
		Flags:       section64.AttrPureInstructions | section64.AttrSomeInstructions,
	}

	segmentBuilder := segment64.Segment64Builder{
		SegmentName:        [16]byte{'_', '_', 'T', 'E', 'X', 'T', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Sections:           []section64.Section64Builder{sectionBuilder},
		VirtualMemorySize:  8,
		MaxProtections:     segment64.AllowAllProtection,
		InitialProtections: segment64.AllowAllProtection,
	}

	machoBuilder := builder.MachoBuilder{
		Header:   headerBuilder,
		Commands: []builder.CommandBuilder{segmentBuilder},
	}

	expected := []byte{
		0xCF, 0xFA, 0xED, 0xFE, 0x0C, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x98, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x19, 0x00, 0x00, 0x00, 0x98, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xB8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x74, 0x65, 0x78, 0x74, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xB8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x01, 0x8B, 0xC0, 0x03, 0x5F, 0xD6,
	}

	buffer := new(bytes.Buffer)
	n, err := machoBuilder.WriteTo(buffer)

	assert.NoError(t, err)
	assert.EqualValues(t, len(expected), n)
	assert.Equal(t, expected, buffer.Bytes())
}
