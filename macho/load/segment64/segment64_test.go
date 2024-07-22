package segment64_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/load"
	"github.com/RealA10N/macho-aarch64-asm/macho/load/segment64"
	"github.com/stretchr/testify/assert"
)

func TestSegment64HeaderExpectedMarshalBinary(t *testing.T) {
	expected := []byte{
		0x19, 0x00, 0x00, 0x00, 0x98, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x38, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	segment := segment64.Segment64Header{
		CommandType:          load.Segment64,
		CommandSize:          152,
		SegmentName:          [16]byte{},
		VirtualMemoryAddress: 0,
		VirtualMemorySize:    8,
		FileOffset:           312,
		FileSize:             8,
		MaxProtections:       segment64.AllowAllProtection,
		InitialProtections:   segment64.AllowAllProtection,
		NumberOfSections:     1,
		Flags:                0,
	}

	got, err := segment.MarshalBinary()

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}
