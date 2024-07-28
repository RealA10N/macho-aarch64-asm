package nlist64

import "github.com/RealA10N/macho-aarch64-asm/utils"

type SymbolType uint8

const (
	// The n_type (SymbolType) field really contains four fields:
	//  unsigned char N_STAB:3, (DebugSymbolMask)
	//  N_PEXT:1, (PrivateExternalSymbolBit)
	//  N_TYPE:3, (TypeSymbolMask)
	//  N_EXT:1; (ExternalSymbolBit)
	// which are used via the following masks.

	DebugSymbolMask       SymbolType = 0xe0 // if any of these bits set, a symbolic debugging entry
	PrivateExternalSymbol SymbolType = 0x10 // private external symbol bit
	TypeSymbolMask        SymbolType = 0x0e // mask for the type bits
	ExternalSymbol        SymbolType = 0x01 // external symbol bit, set for external symbols

	// Only symbolic debugging entries have some of the N_STAB bits set and if any
	// of these bits are set then it is a symbolic debugging entry (a stab).  In
	// which case then the values of the n_type field (the entire field) are given
	// in <mach-o/stab.h>

	// Values for N_TYPE bits of the n_type field.
	UndefinedSymbolType         SymbolType = 0x0 // undefined, n_sect == NO_SECT
	AbsoluteSymbolType          SymbolType = 0x2 // absolute, n_sect == NO_SECT
	SectionSymbolType           SymbolType = 0xe // defined in section number n_sect
	PreboundUndefinedSymbolType SymbolType = 0xc // prebound undefined (defined in a dylib)
	IndirectSymbolType          SymbolType = 0xa // indirect
)

type SymbolDescription uint16

const (
	ReferenceTypeMask SymbolDescription = 0x7

	// types of references
	// non lazy: 	data symbol
	// lazy: 		function symbol
	// private:		visible only to this shared library
	// defined:		references data/function in this module
	// undefined:	references data/function in another module, should be addressed by linker or compiler
	ReferenceFlagUndefinedNonLazy        SymbolDescription = 0
	ReferenceFlagUndefinedLazy           SymbolDescription = 1
	ReferenceFlagDefined                 SymbolDescription = 2
	ReferenceFlagPrivateDefined          SymbolDescription = 3
	ReferenceFlagPrivateUndefinedNonLazy SymbolDescription = 4
	ReferenceFlagPrivateUndefinedLazy    SymbolDescription = 5

	// additional flags
	ReferencedDynamically SymbolDescription = 0x10
	DescriptionDiscarded  SymbolDescription = 0x20 // Used by the dynamic linker at runtime. Do not set this bit
	WeakReference         SymbolDescription = 0x40
	WeakDefinition        SymbolDescription = 0x80
)

type Nlist64 struct {
	StringTableOffset uint32
	SymbolType        SymbolType
	Section           uint8 // the number of the section that this symbol can be found in
	Description       SymbolDescription
	Value             uint64 // if (SymbolType == SectionSymbolType), this is an offset in the section
}

const Nlist64Size uint64 = 0x10

func (nlist64 Nlist64) MarshalBinary() ([]byte, error) {
	return utils.GenericMarshalBinary(nlist64)
}
