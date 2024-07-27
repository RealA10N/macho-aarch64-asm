package header

import "github.com/RealA10N/macho-aarch64-asm/macho/builder/context"

type MachoHeaderBuilder struct {
	Magic      Magic
	CpuType    CpuType
	CpuSubType CpuSubType
	FileType   FileType
	Flags      Flags
}

func (header MachoHeaderBuilder) Build(
	ctx *context.CommandContext,
) MachoHeader {
	return MachoHeader{
		Magic:              header.Magic,
		CpuType:            header.CpuType,
		CpuSubType:         header.CpuSubType,
		FileType:           header.FileType,
		NumOfLoadCommands:  ctx.NumOfLoadCommands,
		SizeOfLoadCommands: ctx.SizeOfLoadCommands,
		Flags:              header.Flags,
	}
}
