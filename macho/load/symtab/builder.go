package symtab

import (
	"io"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/load"
	writertoutils "github.com/RealA10N/writer-to-utils"
)

type SymtabContext struct {
	StringTableOffset uint32
}

type SymbolBuilder interface {
	StringLen() uint32
	WriteToStringTable(io.Writer) (int64, error)
	WriteToEntryList(io.Writer, *SymtabContext) (int64, error)
}

type SymtabBuilder struct {
	Symbols []SymbolBuilder
}

func (builder SymtabBuilder) Build(
	ctx *context.CommandContext,
) SymtabHeader {
	return SymtabHeader{
		CommandType:       load.SymbolTable,
		CommandSize:       uint32(SymTabHeaderSize),
		SymbolTableOffset: uint32(ctx.DataOffset),
		NumOfSymbols:      uint32(len(builder.Symbols)),
		StringTableOffset: uint32(ctx.DataOffset) + uint32(builder.entryListLen()),
		StringTableSize:   uint32(builder.stringTableLen()),
	}
}

// private methods

func (builder SymtabBuilder) entryListLen() uint64 {
	return Nlist64Size * uint64(len(builder.Symbols))
}

func (builder SymtabBuilder) stringTableLen() uint64 {
	var len uint64 = 1 // string table is always prefixed with a nullbyte.
	for _, symbol := range builder.Symbols {
		len += uint64(symbol.StringLen()) + 1 // add null terminator for each string
	}
	return len
}

// data builders

type writerToEntryList struct{ *SymtabBuilder }

func (builder writerToEntryList) WriteTo(writer io.Writer) (n int64, err error) {
	var k int64

	ctx := SymtabContext{
		// the string table is always prefixed with a nullbyte,
		// so the initial offset is 1 and not 0.
		StringTableOffset: 1,
	}

	for _, symbol := range builder.Symbols {
		k, err = symbol.WriteToEntryList(writer, &ctx)
		n += k
		if err != nil {
			return
		}

		// +1 for the null terminating byte
		ctx.StringTableOffset += symbol.StringLen() + 1
	}

	return
}

type symbolWriterToStringTable struct{ SymbolBuilder }

func (symbol symbolWriterToStringTable) WriteTo(writer io.Writer) (int64, error) {
	return symbol.WriteToStringTable(writer)
}

type writerToStringTable struct{ *SymtabBuilder }

func (builder writerToStringTable) WriteTo(writer io.Writer) (int64, error) {
	nullByteWriterTo := writertoutils.BufferWriterTo([]byte{0})
	writerTos := []io.WriterTo{nullByteWriterTo}
	for _, symbol := range builder.Symbols {
		symbolWriterTo := symbolWriterToStringTable{symbol}
		writerTos = append(writerTos, symbolWriterTo, nullByteWriterTo)
	}
	multiWriterTo := writertoutils.MultiWriterTo(writerTos...)
	return multiWriterTo.WriteTo(writer)
}

// CommandBuilder Implementation

func (builder SymtabBuilder) HeaderLen() uint64 {
	return SymTabHeaderSize
}

func (builder SymtabBuilder) DataLen() uint64 {
	return builder.entryListLen() + builder.stringTableLen()
}

func (builder SymtabBuilder) HeaderWriteTo(
	writer io.Writer,
	ctx *context.CommandContext,
) (int64, error) {
	header := builder.Build(ctx)
	writerTo := writertoutils.BinaryMarshalerAdapter(header)
	return writerTo.WriteTo(writer)
}

func (builder SymtabBuilder) DataWriteTo(writer io.Writer) (int64, error) {
	writerTos := []io.WriterTo{
		writerToEntryList{&builder},
		writerToStringTable{&builder},
	}
	multiWriterTo := writertoutils.MultiWriterTo(writerTos...)
	return multiWriterTo.WriteTo(writer)
}
