package macho

import "github.com/RealA10N/macho-aarch64-asm/macho/header"

type machoBuilder struct {
	header header.MachoHeader
}

func MachoBuilder(cpu header.CpuType, subtype header.CpuSubType, filetype header.FileType, flags header.Flags) machoBuilder {
	return machoBuilder{
		header: header.MachoHeader{
			Magic:      cpu.ToMagic(),
			CpuType:    cpu,
			CpuSubType: subtype,
			FileType:   filetype,
			Flags:      flags,
		},
	}
}
