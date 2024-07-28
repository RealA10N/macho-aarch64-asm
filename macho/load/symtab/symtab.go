package symtab

import (
	"github.com/RealA10N/macho-aarch64-asm/macho/load"
	"github.com/RealA10N/macho-aarch64-asm/utils"
)

type SymtabHeader struct {
	CommandType       load.CommandType
	CommandSize       uint32
	SymbolTableOffset uint32
	NumOfSymbols      uint32
	StringTableOffset uint32
	StringTableSize   uint32
}

const SymTabHeaderSize uint64 = 0x18

func (symtab SymtabHeader) MarshalBinary() ([]byte, error) {
	return utils.GenericMarshalBinary(symtab)
}
