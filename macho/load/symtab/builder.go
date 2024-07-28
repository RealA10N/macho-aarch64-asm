package symtab

import (
	"fmt"
	"io"
	"unicode"

	"github.com/RealA10N/macho-aarch64-asm/macho/builder/context"
	"github.com/RealA10N/macho-aarch64-asm/macho/load"
	writertoutils "github.com/RealA10N/writer-to-utils"
)

type SymtabContext struct {
	StringTableOffset uint32
}

func (ctx *SymtabContext) updateAfterProcessing(sym SymbolBuilder) {
	ctx.StringTableOffset += uint32(len(sym.GenString())) + 1
}

type SymbolBuilder interface {
	GenString() string
	GenEntryList(*SymtabContext) (Nlist64, error)
}

type SymtabBuilder struct {
	Symbols []SymbolBuilder
}

// private methods

func (builder SymtabBuilder) buildHeaderFromCtx(
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

func (builder SymtabBuilder) entryListLen() uint64 {
	return Nlist64Size * uint64(len(builder.Symbols))
}

func (builder SymtabBuilder) stringTableLen() uint64 {
	var n uint64 = 1 // string table is always prefixed with a nullbyte.
	for _, symbol := range builder.Symbols {
		n += uint64(len(symbol.GenString())) + 1 // add null terminator for each string
	}
	return n
}

// data builders

type writerToEntryList struct{ *SymtabBuilder }

func (builder writerToEntryList) WriteTo(writer io.Writer) (n int64, err error) {
	ctx := SymtabContext{
		// the string table is always prefixed with a nullbyte,
		// so the initial offset is 1 and not 0.
		StringTableOffset: 1,
	}

	writerTos := []io.WriterTo{}

	for _, symbol := range builder.Symbols {
		var nlist Nlist64
		nlist, err = symbol.GenEntryList(&ctx)
		if err != nil {
			return
		}

		nlistWriterTo := writertoutils.BinaryMarshalerAdapter(nlist)
		writerTos = append(writerTos, nlistWriterTo)
		ctx.updateAfterProcessing(symbol)
	}

	multiWriterTo := writertoutils.MultiWriterTo(writerTos...)
	return multiWriterTo.WriteTo(writer)
}

type symbolWriterToStringTable struct{ SymbolBuilder }

func (symbol symbolWriterToStringTable) WriteTo(writer io.Writer) (int64, error) {
	str := symbol.GenString()
	for i, rune := range str {
		if rune > unicode.MaxASCII {
			err := fmt.Errorf("symbol %s is not an ascii string (at rune index %d)", str, i)
			return 0, err
		}
	}

	n, err := io.WriteString(writer, str)
	return int64(n), err
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
	header := builder.buildHeaderFromCtx(ctx)
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
