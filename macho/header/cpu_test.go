package header_test

import (
	"testing"

	"github.com/RealA10N/macho-aarch64-asm/macho/header"
	"github.com/stretchr/testify/assert"
)

var Arch64Cpus = []header.CpuType{
	header.Arm64CpuType,
	header.X8664CpuType,
	header.PowerPC64CpuType,
}

var Arch32Cpus = []header.CpuType{
	header.VaxCpuType,
	header.RompCpuType,
	header.Ns32032CpuType,
	header.Ns32332CpuType,
	header.Mc680x0CpuType,
	header.X86CpuType,
	header.MipsCpuType,
	header.Ns32352CpuType,
	header.Mc98000CpuType,
	header.HppaCpuType,
	header.ArmCpuType,
	header.Mc88000CpuType,
	header.SparcCpuType,
	header.I860BigCpuType,
	header.I860LittleCpuType,
	header.Rs6000CpuType,
	header.PowerPCCpuType,
}

func TestIsArch64(t *testing.T) {
	for _, cpu := range Arch32Cpus {
		assert.False(t, cpu.IsArch64())
	}

	for _, cpu := range Arch64Cpus {
		assert.True(t, cpu.IsArch64())
	}
}

func TestMagic(t *testing.T) {
	for _, cpu := range Arch32Cpus {
		assert.EqualValues(t, 0xfeedface, cpu.ToMagic())
	}

	for _, cpu := range Arch64Cpus {
		assert.EqualValues(t, 0xfeedfacf, cpu.ToMagic())
	}
}
