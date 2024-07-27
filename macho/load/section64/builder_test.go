package section64_test

import (
	"bytes"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/section64"
	"github.com/stretchr/testify/assert"
)

func TestSegmentBuilderExpectedBinary(t *testing.T) {
	expectedHeader := []byte{
		0x5F, 0x5F, 0x74, 0x65, 0x78, 0x74, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x5F, 0x5F, 0x54, 0x45, 0x58, 0x54, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x38, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x04, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	data := []byte{0x00, 0x00, 0x01, 0x8B, 0xC0, 0x03, 0x5F, 0xD6}

	sectionBuilder := section64.Section64Builder{
		SectionName: [16]byte{'_', '_', 't', 'e', 'x', 't', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		SegmentName: [16]byte{'_', '_', 'T', 'E', 'X', 'T', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Data:        data,
		Flags:       section64.AttrPureInstructions | section64.AttrSomeInstructions,
	}

	assert.EqualValues(t, len(expectedHeader), sectionBuilder.HeaderLen())
	assert.EqualValues(t, len(data), sectionBuilder.DataLen())

	{
		buffer := new(bytes.Buffer)
		ctx := context.CommandContext{DataOffset: 312}
		n, err := sectionBuilder.HeaderWriteTo(buffer, &ctx)

		assert.NoError(t, err)
		assert.EqualValues(t, len(expectedHeader), n)
		assert.Equal(t, expectedHeader, buffer.Bytes())
	}

	{
		buffer := new(bytes.Buffer)
		n, err := sectionBuilder.DataWriteTo(buffer)
		assert.NoError(t, err)
		assert.EqualValues(t, len(data), n)
		assert.Equal(t, data, buffer.Bytes())
	}
}
