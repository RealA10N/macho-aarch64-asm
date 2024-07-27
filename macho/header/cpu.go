package header

type CpuType uint32

const (
	// Sources:
	// https://en.wikipedia.org/wiki/Mach-O
	// https://opensource.apple.com/source/xnu/xnu-4570.41.2/osfmk/mach/machine.h.auto.html

	VaxCpuType        CpuType = 0x00000001
	RompCpuType       CpuType = 0x00000002
	Ns32032CpuType    CpuType = 0x00000004
	Ns32332CpuType    CpuType = 0x00000005
	Mc680x0CpuType    CpuType = 0x00000006
	X86CpuType        CpuType = 0x00000007
	MipsCpuType       CpuType = 0x00000008
	Ns32352CpuType    CpuType = 0x00000009
	Mc98000CpuType    CpuType = 0x0000000A
	HppaCpuType       CpuType = 0x0000000B
	ArmCpuType        CpuType = 0x0000000C
	Mc88000CpuType    CpuType = 0x0000000D
	SparcCpuType      CpuType = 0x0000000E
	I860BigCpuType    CpuType = 0x0000000F
	I860LittleCpuType CpuType = 0x00000010
	Rs6000CpuType     CpuType = 0x00000011
	PowerPCCpuType    CpuType = 0x00000012

	ABI64Arch        CpuType = 0x01000000 // 64 bit ABI
	X8664CpuType     CpuType = X86CpuType | ABI64Arch
	Arm64CpuType     CpuType = ArmCpuType | ABI64Arch
	PowerPC64CpuType CpuType = PowerPCCpuType | ABI64Arch
)

func (cpu CpuType) IsArch64() bool {
	return (cpu & ABI64Arch) != 0
}

func (cpu CpuType) ToMagic() Magic {
	if cpu.IsArch64() {
		return Magic64Bit
	} else {
		return Magic32Bit
	}
}

type CpuSubType uint32

const (
	// Source: https://en.wikipedia.org/wiki/Mach-O

	AllArmProcessors   CpuSubType = 0x00000000
	ArmA500ARCHOrNewer CpuSubType = 0x00000001
	ArmA500OrNewer     CpuSubType = 0x00000002
	ArmA440OrNewer     CpuSubType = 0x00000003
	ArmM4OrNewer       CpuSubType = 0x00000004
	ArmV4TOrNewer      CpuSubType = 0x00000005
	ArmV6OrNewer       CpuSubType = 0x00000006
	ArmV5TEJOrNewer    CpuSubType = 0x00000007
	ArmXSCALEOrNewer   CpuSubType = 0x00000008
	ArmV7OrNewer       CpuSubType = 0x00000009
	ArmV7FOrNewer      CpuSubType = 0x0000000A
	ArmV7SOrNewer      CpuSubType = 0x0000000B
	ArmV7KOrNewer      CpuSubType = 0x0000000C
	ArmV8OrNewer       CpuSubType = 0x0000000D
	ArmV6MOrNewer      CpuSubType = 0x0000000E
	ArmV7MOrNewer      CpuSubType = 0x0000000F
	ArmV7EMOrNewer     CpuSubType = 0x00000010

	Allx86Processors       CpuSubType = 0x00000003
	X86486OrNewer          CpuSubType = 0x00000004
	X86486SXOrNewer        CpuSubType = 0x00000084
	X86PentiumM5OrNewer    CpuSubType = 0x00000056
	X86CeleronOrNewer      CpuSubType = 0x00000067
	X86CeleronMobile       CpuSubType = 0x00000077
	X86Pentium3OrNewer     CpuSubType = 0x00000008
	X86Pentium3MOrNewer    CpuSubType = 0x00000018
	X86Pentium3XEONOrNewer CpuSubType = 0x00000028
	X86Pentium4OrNewer     CpuSubType = 0x0000000A
	X86ItaniumOrNewer      CpuSubType = 0x0000000B
	X86Itanium2OrNewer     CpuSubType = 0x0000001B
	X86XEONOrNewer         CpuSubType = 0x0000000C
	X86XEONMPOrNewer       CpuSubType = 0x0000001C
)
