package segment64

import (
	"bytes"
	"encoding/binary"

	"github.com/RealA10N/macho-aarch64-asm/macho/load"
)

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

type Segment64Header struct {
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

const Segment64HeaderSize uint64 = 0x48

func (segment Segment64Header) MarshalBinary() ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, segment)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
