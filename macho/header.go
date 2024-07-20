// References:
// https://en.wikipedia.org/wiki/Mach-O
// https://github.com/aidansteele/osx-abi-macho-file-format-reference
// https://github.com/opensource-apple/cctools/

package macho

type HeaderMagic uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	Magic32Bit HeaderMagic = 0xfeedface
	Magic64Bit HeaderMagic = 0xfeedfacf
)

type HeaderCpuType uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	VAX        HeaderCpuType = 0x00000001
	ROMP       HeaderCpuType = 0x00000002
	NS32032    HeaderCpuType = 0x00000004
	NS32332    HeaderCpuType = 0x00000005
	MC680x0    HeaderCpuType = 0x00000006
	x86        HeaderCpuType = 0x00000007
	MIPS       HeaderCpuType = 0x00000008
	NS32352    HeaderCpuType = 0x00000009
	MC98000    HeaderCpuType = 0x0000000A
	HPPA       HeaderCpuType = 0x0000000B
	ARM        HeaderCpuType = 0x0000000C
	MC88000    HeaderCpuType = 0x0000000D
	SPARC      HeaderCpuType = 0x0000000E
	i860Big    HeaderCpuType = 0x0000000F
	i860Little HeaderCpuType = 0x00000010
	RS6000     HeaderCpuType = 0x00000011
	PowerPC    HeaderCpuType = 0x00000012
)

type HeaderCpuSubType uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	AllArmProcessors   HeaderCpuSubType = 0x00000000
	ArmA500ARCHOrNewer HeaderCpuSubType = 0x00000001
	ArmA500OrNewer     HeaderCpuSubType = 0x00000002
	ArmA440OrNewer     HeaderCpuSubType = 0x00000003
	ArmM4OrNewer       HeaderCpuSubType = 0x00000004
	ArmV4TOrNewer      HeaderCpuSubType = 0x00000005
	ArmV6OrNewer       HeaderCpuSubType = 0x00000006
	ArmV5TEJOrNewer    HeaderCpuSubType = 0x00000007
	ArmXSCALEOrNewer   HeaderCpuSubType = 0x00000008
	ArmV7OrNewer       HeaderCpuSubType = 0x00000009
	ArmV7FOrNewer      HeaderCpuSubType = 0x0000000A
	ArmV7SOrNewer      HeaderCpuSubType = 0x0000000B
	ArmV7KOrNewer      HeaderCpuSubType = 0x0000000C
	ArmV8OrNewer       HeaderCpuSubType = 0x0000000D
	ArmV6MOrNewer      HeaderCpuSubType = 0x0000000E
	ArmV7MOrNewer      HeaderCpuSubType = 0x0000000F
	ArmV7EMOrNewer     HeaderCpuSubType = 0x00000010

	Allx86Processors       HeaderCpuSubType = 0x00000003
	x86486OrNewer          HeaderCpuSubType = 0x00000004
	x86486SXOrNewer        HeaderCpuSubType = 0x00000084
	x86PentiumM5OrNewer    HeaderCpuSubType = 0x00000056
	x86CeleronOrNewer      HeaderCpuSubType = 0x00000067
	x86CeleronMobile       HeaderCpuSubType = 0x00000077
	x86Pentium3OrNewer     HeaderCpuSubType = 0x00000008
	x86Pentium3MOrNewer    HeaderCpuSubType = 0x00000018
	x86Pentium3XEONOrNewer HeaderCpuSubType = 0x00000028
	x86Pentium4OrNewer     HeaderCpuSubType = 0x0000000A
	x86ItaniumOrNewer      HeaderCpuSubType = 0x0000000B
	x86Itanium2OrNewer     HeaderCpuSubType = 0x0000001B
	x86XEONOrNewer         HeaderCpuSubType = 0x0000000C
	x86XEONMPOrNewer       HeaderCpuSubType = 0x0000001C
)

type HeaderFileType uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	ObjectFile             HeaderFileType = 0x00000001 // Relocatable object file.
	ExecutableFile         HeaderFileType = 0x00000002 // Demand paged executable file.
	FixedVMLibFile         HeaderFileType = 0x00000003 // Fixed VM shared library file.
	CoreFile               HeaderFileType = 0x00000004 // Core file.
	PreloadFile            HeaderFileType = 0x00000005 // Preloaded executable file.
	DynamicLibraryFile     HeaderFileType = 0x00000006 // Dynamically bound shared library file.
	DynamicLinkerFile      HeaderFileType = 0x00000007 // Dynamic link editor.
	BundleFile             HeaderFileType = 0x00000008 // Dynamically bound bundle file.
	DynamicLibraryStubFile HeaderFileType = 0x00000009 // Shared library stub for static linking only, no section contents.
	DynamicSymbolsFile     HeaderFileType = 0x0000000A // Companion file with only debug sections.
	KextsBundleFile        HeaderFileType = 0x0000000B // x86_64 kexts.
	ComposedFile           HeaderFileType = 0x0000000C // a file composed of other Mach-Os to be run in the same userspace sharing a single linkedit.
)

type HeaderFlags uint32

const (
	// Sources:
	// https://en.wikipedia.org/wiki/Mach-O
	// https://github.com/opensource-apple/cctools/blob/fdb4825f303fd5c0751be524babd32958181b3ed/include/mach-o/loader.h#L125
	// There might exist additional flags stated below

	NoUndefsFlag              HeaderFlags = 1 << 0  // The object file has no undefined references.
	IncLinkFlag               HeaderFlags = 1 << 1  // The object file is the output of an incremental link against a base file and can't be link edited again.
	DyLdLinkFlag              HeaderFlags = 1 << 2  // The object file is input for the dynamic linker and can't be statically link edited again.
	BinDatLoadFlag            HeaderFlags = 1 << 3  // The object file's undefined references are bound by the dynamic linker when loaded.
	PreboundFlag              HeaderFlags = 1 << 4  // The file has its dynamic undefined references prebound.
	SplitSegsFlag             HeaderFlags = 1 << 5  // The file has its read-only and read-write segments split.
	LazyInitFlag              HeaderFlags = 1 << 6  // The shared library init routine is to be run lazily via catching memory faults to its writeable segments (obsolete).
	TwoLevelFlag              HeaderFlags = 1 << 7  // The image is using two-level name space bindings.
	ForceFlatFlag             HeaderFlags = 1 << 8  // The executable is forcing all images to use flat name space bindings.
	NoMultiDefsFlag           HeaderFlags = 1 << 9  // This umbrella guarantees no multiple definitions of symbols in its sub-images so the two-level namespace hints can always be used.
	NoFixPreBindingFlag       HeaderFlags = 1 << 10 // Do not have dyld notify the prebinding agent about this executable.
	PreBindableFlag           HeaderFlags = 1 << 11 // The binary is not prebound but can have its prebinding redone. only used when MH_PREBOUND is not set.
	AllModsBoundFlag          HeaderFlags = 1 << 12 // Indicates that this binary binds to all two-level namespace modules of its dependent libraries.
	SubsectionsViaSymbolsFlag HeaderFlags = 1 << 13 // Safe to divide up the sections into sub-sections via symbols for dead code stripping.
	CanonicalFlag             HeaderFlags = 1 << 14 // The binary has been canonicalized via the un-prebind operation.
	WeakDefinesFlag           HeaderFlags = 1 << 15 // The final linked image contains external weak symbols.
	BindsToWeakFlag           HeaderFlags = 1 << 16 // The final linked image uses weak symbols.
	AllowStackExecutionFlag   HeaderFlags = 1 << 17 // When this bit is set, all stacks in the task will be given stack execution privilege.
	RootSafeFlag              HeaderFlags = 1 << 18 // When this bit is set, the binary declares it is safe for use in processes with uid zero.
	SetUidSafeFlag            HeaderFlags = 1 << 19 // When this bit is set, the binary declares it is safe for use in processes when UGID is true.
	NoReExportedDyLibsFlag    HeaderFlags = 1 << 20 // When this bit is set on a dylib, the static linker does not need to examine dependent dylibs to see if any are re-exported.
	PieFlag                   HeaderFlags = 1 << 21 // When this bit is set, the OS will load the main executable at a random address.
	DeadStrippableDyLibFlag   HeaderFlags = 1 << 22 // Only for use on dylibs. When linking against a dylib that has this bit set, the static linker will automatically not create a load command to the dylib if no symbols are being referenced from the dylib.
	HasTlvDescriptorsFlag     HeaderFlags = 1 << 23 // Contains a section of type S_THREAD_LOCAL_VARIABLES.
	NoHeapExecutionFlag       HeaderFlags = 1 << 24 // When this bit is set, the OS will run the main executable with a non-executable heap even on platforms (e.g. i386) that don't require it.
	AppExtensionSafeFlag      HeaderFlags = 1 << 25 // The code was linked for use in an application.
)

type MachoHeader struct {
	magic              HeaderMagic
	cpuType            HeaderCpuType
	cpuSubType         HeaderCpuSubType
	fileType           HeaderFileType
	numOfLoadCommands  uint32
	sizeOfLoadCommands uint32
	flags              HeaderFlags
	reserved           uint32 // additional 32 padding bits are required in 64bit architectures
}
