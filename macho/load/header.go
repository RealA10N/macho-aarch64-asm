package load

type CommandType uint32

const (
	// Source: https://opensource.apple.com/source/xnu/xnu-4570.41.2/EXTERNAL_HEADERS/mach-o/loader.h.auto.html

	// After MacOS X 10.1 when a new load command is added that is required to be
	// understood by the dynamic linker for the image to execute properly the
	// 'Required' bit will be or'ed into the load command constant.  If the dynamic
	// linker sees such a load command it it does not understand will issue a
	// "unknown load command required for execution" error and refuse to use the
	// image.  Other load commands without this bit that are not understood will
	// simply be ignored.
	Required CommandType = 0x80000000

	Segment                CommandType = 0x1             // segment of this file to be mapped
	SymbolTable            CommandType = 0x2             // link-edit stab symbol table info
	SymbolSegment          CommandType = 0x3             // link-edit gdb symbol table info (obsolete)
	Thread                 CommandType = 0x4             // thread
	UnixThread             CommandType = 0x5             // unix thread (includes a stack)
	FixedVMLib             CommandType = 0x6             // load a specified fixed VM shared library
	IdFixedVMLib           CommandType = 0x7             // fixed VM shared library identification
	Id                     CommandType = 0x8             // object identification info (obsolete)
	FixedVMFile            CommandType = 0x9             // fixed VM file inclusion (internal use)
	Prepage                CommandType = 0xa             // prepage command (internal use)
	DynamicSymTable        CommandType = 0xb             // dynamic link-edit symbol table info
	DynamicLib             CommandType = 0xc             // load a dynamically linked shared library
	IdDynamicLib           CommandType = 0xd             // dynamically linked shared lib ident
	DynamicLinker          CommandType = 0xe             // load a dynamic linker
	IdDynamicLinker        CommandType = 0xf             // dynamic linker identification
	PreboundDynamicLib     CommandType = 0x10            // modules prebound for a dynamically linked shared library
	Routines               CommandType = 0x11            // image routines
	SubFramework           CommandType = 0x12            // sub framework
	SubUmbrella            CommandType = 0x13            // sub umbrella
	SubClient              CommandType = 0x14            // sub client
	SubLibrary             CommandType = 0x15            // sub library
	TwoLevelHints          CommandType = 0x16            // two-level namespace lookup hints
	PrebindChecksum        CommandType = 0x17            // prebind checksum
	LoadWeakDynamicLib     CommandType = 0x18 | Required // load a dynamically linked shared library that is allowed to be missing (all symbols are weak imported).
	Segment64              CommandType = 0x19            // 64-bit segment of this file to be mapped
	Routines64             CommandType = 0x1a            // 64-bit image routines
	Uuid                   CommandType = 0x1b            // the uuid
	RunPath                CommandType = 0x1c | Required // runpath additions
	CodeSig                CommandType = 0x1d            // local of code signature
	SegSplitInfo           CommandType = 0x1e            // local of info to split segments
	ReExportDynamicLib     CommandType = 0x1f | Required // load and re-export dylib
	LazyLoadDynamicLib     CommandType = 0x20            // delay load of dylib until first use
	EncryptionInfo         CommandType = 0x21            // encrypted segment information
	DyLdInfo               CommandType = 0x22            // compressed dyld information
	DyLdInfoOnly           CommandType = 0x22 | Required // compressed dyld information only
	LoadUnwardDynamicLib   CommandType = 0x23 | Required // load upward dylib
	VersionMinMacOS        CommandType = 0x24            // build for MacOSX min OS version
	VersionMinIOS          CommandType = 0x25            // build for iPhoneOS min OS version
	FunctionStarts         CommandType = 0x26            // compressed table of function start addresses
	DyLdEnvironment        CommandType = 0x27            // string for dyld to treat like environment variable
	Main                   CommandType = 0x28 | Required // replacement for LC_UNIXTHREAD
	DataInCode             CommandType = 0x29            // table of non-instructions in __text
	SourceVersion          CommandType = 0x2A            // source version used to build binary
	DynamicLibCodeSignDrs  CommandType = 0x2B            // Code signing DRs copied from linked dylibs
	EncryptionInfo64       CommandType = 0x2C            // 64-bit encrypted segment information
	LinkerOption           CommandType = 0x2D            // linker options in MH_OBJECT files
	LinkerOptimizationHint CommandType = 0x2E            // optimization hints in MH_OBJECT files
	VersionMinTVOS         CommandType = 0x2F            // build for AppleTV min OS version
	VersionMinWatchOS      CommandType = 0x30            // build for Watch min OS version
	BuildVersion           CommandType = 0x32            // build for platform min OS version
	Note                   CommandType = 0x31            // arbitrary data included within a Mach-O file
)
