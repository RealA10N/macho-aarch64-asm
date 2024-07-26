package section64

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Section64Flags uint32

const (
	// Source: https://opensource.apple.com/source/xnu/xnu-4570.41.2/EXTERNAL_HEADERS/mach-o/loader.h.auto.html

	// Constants for the type of a section

	RegularSection            Section64Flags = 0x0 // regular section
	ZeroFillSection           Section64Flags = 0x1 // zero fill on demand section
	CStringsSection           Section64Flags = 0x2 // section with only literal C strings
	Only4ByteLiteralSection   Section64Flags = 0x3 // section with only 4 byte literals
	Only8ByteLiteralSection   Section64Flags = 0x4 // section with only 8 byte literals
	OnlyPointerLiteralSection Section64Flags = 0x5 // section with only pointers to literals

	// For the two types of symbol pointers sections and the symbol stubs section
	// they have indirect symbol table entries.  For each of the entries in the
	// section the indirect symbol table entries, in corresponding order in the
	// indirect symbol table, start at the index stored in the reserved1 field
	// of the section structure.  Since the indirect symbol table entries
	// correspond to the entries in the section the number of indirect symbol table
	// entries is inferred from the size of the section divided by the size of the
	// entries in the section.  For symbol pointers sections the size of the entries
	// in the section is 4 bytes and for symbol stubs sections the byte size of the
	// stubs is stored in the reserved2 field of the section structure.

	NonLazySymbolPointersSection        Section64Flags = 0x6  // section with only non-lazy symbol pointers
	LazySymbolPointersSection           Section64Flags = 0x7  // section with only lazy symbol pointers
	SymbolStubsSection                  Section64Flags = 0x8  // section with only symbol stubs, byte size of stub in the reserved2 field
	ModInitFuncPointersSection          Section64Flags = 0x9  // section with only function pointers for initialization
	ModTermFuncPointersSection          Section64Flags = 0xa  // section with only function pointers for termination
	CoalescedSection                    Section64Flags = 0xb  // section contains symbols that are to be coalesced
	LargeZeroFillSection                Section64Flags = 0xc  // zero fill on demand section (that can be larger than 4 gigabytes)
	InterposingSection                  Section64Flags = 0xd  // section with only pairs of function pointers for interposing
	Only16ByteLiteralSection            Section64Flags = 0xe  // section with only 16 byte literals
	DTraceDofSection                    Section64Flags = 0xf  // section contains DTrace Object Format
	LazyDynamicLibSymbolPointersSection Section64Flags = 0x10 // section with only lazy symbol pointers to lazy loaded dylibs

	// Section types to support thread local variables

	ThreadLocalRegularSection              Section64Flags = 0x11 // template of initial values for TLVs
	ThreadLocalZeroFillSection             Section64Flags = 0x12 // template of initial values for TLVs
	ThreadLocalVariablesSection            Section64Flags = 0x13 // TLV descriptors
	ThreadLocalVariablePointersSection     Section64Flags = 0x14 // pointers to TLV descriptors
	ThreadLocalInitFunctionPointersSection Section64Flags = 0x15 // functions to call to initialize TLV values

	// Constants for the section attributes part of the flags field of a section
	// structure.

	AttrPureInstructions   Section64Flags = 0x80000000 // section contains only true machine instructions
	AttrNoToc              Section64Flags = 0x40000000 // section contains coalesced symbols that are not to be in a ranlib table of contents
	AttrStripStaticSyms    Section64Flags = 0x20000000 // ok to strip static symbols in this section in files with the MH_DYLDLINK flag
	AttrNoDeadStrip        Section64Flags = 0x10000000 // no dead stripping
	AttrLiveSupport        Section64Flags = 0x08000000 // blocks are live if they reference live blocks
	AttrSelfModifyingCode  Section64Flags = 0x04000000 // Used with i386 code stubs written on by dyld
	AttrDebug              Section64Flags = 0x02000000 // a debug section
	AttrSomeInstructions   Section64Flags = 0x00000400 // section contains some machine instructions
	AttrExternalRelocation Section64Flags = 0x00000200 // section has external relocation entries
	AttrLocalRelocation    Section64Flags = 0x00000100 // section has local relocation entries
)

type Section64Header struct {
	SectionName         [16]byte
	SegmentName         [16]byte
	Address             uint64
	Size                uint64
	Offset              uint32
	Align               uint32
	RelocationOffset    uint32
	NumberOfRelocations uint32
	Flags               Section64Flags
	Reserved1           uint32
	Reserved2           uint32
	Reserved3           uint32
	// TODO: not sure if there are 2 or 3 reserved values?
	// My tests show that there are 3 but the reference says there are only 2.
}

const Section64HeaderSize uint64 = 0x50

func (section Section64Header) MarshalBinary() ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, section)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (section Section64Header) WriteTo(writer io.Writer) (int64, error) {
	data, err := section.MarshalBinary()
	if err != nil {
		return 0, err
	}

	n, err := writer.Write(data)
	return int64(n), err
}
