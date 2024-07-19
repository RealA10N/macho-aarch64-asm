package sections_test

import (
	"strings"
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/sections"

	"github.com/stretchr/testify/assert"
)

func TestNoGlobalSection(t *testing.T) {
	globals := sections.GlobalSection{Names: []string{}}
	buffer := new(strings.Builder)

	n, err := globals.WriteTo(buffer)

	assert.NoError(t, err)
	assert.Empty(t, buffer)
	assert.EqualValues(t, n, 0)
}

func TestSingleGlobalWriteTo(t *testing.T) {
	globals := sections.GlobalSection{Names: []string{"MyGlobalVariable"}}
	buffer := new(strings.Builder)

	n, err := globals.WriteTo(buffer)
	const expected = ".global MyGlobalVariable\n"

	assert.NoError(t, err)
	assert.EqualValues(t, len(expected), n)
	assert.Equal(t, expected, buffer.String())
}

func TestMultipleGlobals(t *testing.T) {
	globals := sections.GlobalSection{Names: []string{"first", "Second", "_third"}}
	buffer := new(strings.Builder)

	n, err := globals.WriteTo(buffer)
	const expected = ".global first, Second, _third\n"

	assert.NoError(t, err)
	assert.EqualValues(t, len(expected), n)
	assert.Equal(t, expected, buffer.String())
}
