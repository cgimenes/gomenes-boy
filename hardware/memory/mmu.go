package memory

import (
	"neutronstar-gb/hardware/types"
)

type MMU struct {
	addresses [100000]types.Byte
}

func (r *MMU) Get(address types.Word) types.Byte {
	if address <= 0x100 {
		return BootROM[address]
	} else {
		return r.addresses[address]
	}
}

func (r *MMU) Set(address types.Word, value types.Byte) {
	if address <= 0x100 {
		BootROM[address] = value
	} else {
		r.addresses[address] = value
	}
}