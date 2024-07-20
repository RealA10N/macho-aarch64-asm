package commands

import (
	"github.com/RealA10N/macho-aarch64-asm/macho/load"
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

type Section64 struct {
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
	// not sure if there are 2 or 3 reserved values?
	// My tests show that there are 3 but the reference says there are only 2.
}

type VirtualMemoryProtection uint32

const (
	// Source: https://opensource.apple.com/source/xnu/xnu-1456.1.26/osfmk/mach/vm_prot.h.auto.html

	AllowNothingProtection VirtualMemoryProtection = 0x00
	AllowReadProtection    VirtualMemoryProtection = 0x01 // read permission
	AllowWriteProtection   VirtualMemoryProtection = 0x02 // write permission
	AllowExecuteProtection VirtualMemoryProtection = 0x04 // execute permission

	// The default protection for newly-created virtual memory
	DefaultProtection VirtualMemoryProtection = (AllowReadProtection | AllowWriteProtection)

	// The maximum privileges possible, for parameter checking.
	AllowAllProtection VirtualMemoryProtection = (AllowReadProtection | AllowWriteProtection | AllowExecuteProtection)

	// An invalid protection value.
	// Used only by memory_object_lock_request to indicate no change
	// to page locks.  Using -1 here is a bad idea because it
	// looks like VM_PROT_ALL and then some.
	NoCacheProtection VirtualMemoryProtection = 0x08

	// When a caller finds that he cannot obtain write permission on a
	// mapped entry, the following flag can be used.  The entry will
	// be made "needs copy" effectively copying the object (using COW),
	// and write permission will be added to the maximum protections
	// for the associated entry.
	CopyOnWriteProtection VirtualMemoryProtection = 0x10

	// Another invalid protection value.
	// Used only by memory_object_data_request upon an object
	// which has specified a copy_call copy strategy. It is used
	// when the kernel wants a page belonging to a copy of the
	// object, and is only asking the object as a result of
	// following a shadow chain. This solves the race between pages
	// being pushed up by the memory manager and the kernel
	// walking down the shadow chain.
	WantsCopyProtection VirtualMemoryProtection = 0x10
)

type Segment64Flag uint32

const (
	// the file contents for this segment is for the high part of the VM space, the low part is zero filled (for stacks in core files)
	HighVirtualMemory Segment64Flag = 0x1

	// this segment is the VM that is allocated by a fixed VM library, for overlap checking in the link editor
	FixedVirtualMemoryLibrary Segment64Flag = 0x2

	// this segment has nothing that was relocated in it and nothing relocated to it, that is it maybe safely replaced without relocation
	NoRelocation Segment64Flag = 0x4

	// This segment is protected.  If the segment starts at file offset 0, the first page of the segment is not protected. All other pages of the segment are protected.
	ProtectedVersion1 Segment64Flag = 0x8
)

type Segment64 struct {
	CommandType          load.CommandType
	CommandSize          uint32
	SegmentName          [16]byte
	VirtualMemoryAddress uint64
	VirtualMemorySize    uint64
	FileOffset           uint64
	FileSize             uint64
	MaxProtections       VirtualMemoryProtection
	InitialProtections   VirtualMemoryProtection
	NumberOfSections     uint32
	Flags                Segment64Flag
}
