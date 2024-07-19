package sections_test

import (
	"bytes"
	"macho-aarch64-asm/macho/sections"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoGlobalSection(t *testing.T) {
	globals := sections.GlobalSection{Names: []string{}}
	writer := bytes.Buffer{}
	globals.WriteAssembly(&writer)

	assert.Empty(t, writer)
}

func TestSingleGlobal(t *testing.T) {
	globals := sections.GlobalSection{Names: []string{"MyGlobalVariable"}}
	writer := bytes.Buffer{}
	globals.WriteAssembly(&writer)

	assert.NotEmpty(t, writer)
	assert.Equal(t, ".global MyGlobalVariable\n", writer.String())
}

func TestMultipleGlobals(t *testing.T) {
	globals := sections.GlobalSection{Names: []string{"first", "Second", "_third"}}
	writer := bytes.Buffer{}
	globals.WriteAssembly(&writer)

	assert.NotEmpty(t, writer)
	assert.Equal(t, ".global first, Second, _third\n", writer.String())
}
