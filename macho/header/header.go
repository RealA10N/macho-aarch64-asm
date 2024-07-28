// References:
// https://en.wikipedia.org/wiki/Mach-O
// https://github.com/aidansteele/osx-abi-macho-file-format-reference
// https://github.com/opensource-apple/cctools/
// https://opensource.apple.com/source/xnu/xnu-4570.41.2/EXTERNAL_HEADERS/mach-o/loader.h.auto.html

package header

import (
	"bytes"
	"encoding/binary"
)

type Magic uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	Magic32Bit Magic = 0xfeedface
	Magic64Bit Magic = 0xfeedfacf
)

type FileType uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	Object             FileType = 0x00000001 // Relocatable object file.
	Executable         FileType = 0x00000002 // Demand paged executable file.
	FixedVMLib         FileType = 0x00000003 // Fixed VM shared library file.
	Core               FileType = 0x00000004 // Core file.
	Preload            FileType = 0x00000005 // Preloaded executable file.
	DynamicLibrary     FileType = 0x00000006 // Dynamically bound shared library file.
	DynamicLinker      FileType = 0x00000007 // Dynamic link editor.
	Bundle             FileType = 0x00000008 // Dynamically bound bundle file.
	DynamicLibraryStub FileType = 0x00000009 // Shared library stub for static linking only, no section contents.
	DynamicSymbols     FileType = 0x0000000A // Companion file with only debug sections.
	KextsBundle        FileType = 0x0000000B // x86_64 kexts.
	Composed           FileType = 0x0000000C // a file composed of other Mach-Os to be run in the same userspace sharing a single linkedit.
)

type Flags uint32

const (
	// Sources:
	// https://en.wikipedia.org/wiki/Mach-O
	// https://github.com/opensource-apple/cctools/blob/fdb4825f303fd5c0751be524babd32958181b3ed/include/mach-o/loader.h#L125
	// There might exist additional flags stated below

	NoUndefs              Flags = 1 << 0  // The object file has no undefined references.
	IncLink               Flags = 1 << 1  // The object file is the output of an incremental link against a base file and can't be link edited again.
	DyLdLink              Flags = 1 << 2  // The object file is input for the dynamic linker and can't be statically link edited again.
	BinDatLoad            Flags = 1 << 3  // The object file's undefined references are bound by the dynamic linker when loaded.
	Prebound              Flags = 1 << 4  // The file has its dynamic undefined references prebound.
	SplitSegs             Flags = 1 << 5  // The file has its read-only and read-write segments split.
	LazyInit              Flags = 1 << 6  // The shared library init routine is to be run lazily via catching memory faults to its writeable segments (obsolete).
	TwoLevel              Flags = 1 << 7  // The image is using two-level name space bindings.
	ForceFlat             Flags = 1 << 8  // The executable is forcing all images to use flat name space bindings.
	NoMultiDefs           Flags = 1 << 9  // This umbrella guarantees no multiple definitions of symbols in its sub-images so the two-level namespace hints can always be used.
	NoFixPreBinding       Flags = 1 << 10 // Do not have dyld notify the prebinding agent about this executable.
	PreBindable           Flags = 1 << 11 // The binary is not prebound but can have its prebinding redone. only used when MH_PREBOUND is not set.
	AllModsBound          Flags = 1 << 12 // Indicates that this binary binds to all two-level namespace modules of its dependent libraries.
	SubsectionsViaSymbols Flags = 1 << 13 // Safe to divide up the sections into sub-sections via symbols for dead code stripping.
	Canonical             Flags = 1 << 14 // The binary has been canonicalized via the un-prebind operation.
	WeakDefines           Flags = 1 << 15 // The final linked image contains external weak symbols.
	BindsToWeak           Flags = 1 << 16 // The final linked image uses weak symbols.
	AllowStackExecution   Flags = 1 << 17 // When this bit is set, all stacks in the task will be given stack execution privilege.
	RootSafe              Flags = 1 << 18 // When this bit is set, the binary declares it is safe for use in processes with uid zero.
	SetUidSafe            Flags = 1 << 19 // When this bit is set, the binary declares it is safe for use in processes when UGID is true.
	NoReExportedDyLibs    Flags = 1 << 20 // When this bit is set on a dylib, the static linker does not need to examine dependent dylibs to see if any are re-exported.
	Pie                   Flags = 1 << 21 // When this bit is set, the OS will load the main executable at a random address.
	DeadStrippableDyLib   Flags = 1 << 22 // Only for use on dylibs. When linking against a dylib that has this bit set, the static linker will automatically not create a load command to the dylib if no symbols are being referenced from the dylib.
	HasTlvDescriptors     Flags = 1 << 23 // Contains a section of type S_THREAD_LOCAL_VARIABLES.
	NoHeapExecution       Flags = 1 << 24 // When this bit is set, the OS will run the main executable with a non-executable heap even on platforms (e.g. i386) that don't require it.
	AppExtensionSafe      Flags = 1 << 25 // The code was linked for use in an application.
)

type MachoHeader struct {
	Magic              Magic
	CpuType            CpuType
	CpuSubType         CpuSubType
	FileType           FileType
	NumOfLoadCommands  uint32
	SizeOfLoadCommands uint32
	Flags              Flags
	Reserved           uint32 // additional 32 padding bits are required in 64bit architectures
}

const MachoHeaderSize uint64 = 0x20

func (header MachoHeader) MarshalBinary() ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, header)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
