package cpu

import (
	"log"
	"neutronstar-gb/hardware/cpu/registers"
	"neutronstar-gb/hardware/memory"
	"neutronstar-gb/hardware/types"
)

type Instruction struct {
	cycles uint8
	exec func()
}

type CPU struct {
	mmu memory.MMU
	registers registers.Registers
}

func (c *CPU) Init() {
	c.initRegisters()
}

func (c *CPU) initRegisters() {
	c.registers = registers.Registers{
		A:  types.ByteRegister{},
		B:  types.ByteRegister{},
		D:  types.ByteRegister{},
		H:  types.ByteRegister{},
		F:  types.ByteRegister{},
		C:  types.ByteRegister{},
		E:  types.ByteRegister{},
		L:  types.ByteRegister{},
		SP: registers.SingleWordRegister{},
		PC: registers.SingleWordRegister{},
	}
	c.registers.AF = registers.PairedWordRegister{
		H: &c.registers.A,
		L: &c.registers.F,
	}
	c.registers.BC = registers.PairedWordRegister{
		H: &c.registers.B,
		L: &c.registers.C,
	}
	c.registers.DE = registers.PairedWordRegister{
		H: &c.registers.D,
		L: &c.registers.E,
	}
	c.registers.HL = registers.PairedWordRegister{
		H: &c.registers.H,
		L: &c.registers.L,
	}
	c.registers.Flags = registers.Flags{R: &c.registers.F}
}

func (c *CPU) Decode(opcode types.Byte) Instruction {
	switch opcode {
	case 0x06:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.FetchNextByte())
		}, cycles: 8}
	case 0x0E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.FetchNextByte())
		}, cycles: 8}
	case 0x16:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.FetchNextByte())
		}, cycles: 8}
	case 0x1E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.FetchNextByte())
		}, cycles: 8}
	case 0x26:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.FetchNextByte())
		}, cycles: 8}
	case 0x2E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.FetchNextByte())
		}, cycles: 8}
	case 0x7F:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.A.Get())
		}, cycles: 4}
	case 0x78:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.B.Get())
		}, cycles: 4}
	case 0x79:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.C.Get())
		}, cycles: 4}
	case 0x7A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.D.Get())
		}, cycles: 4}
	case 0x7B:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.E.Get())
		}, cycles: 4}
	case 0x7C:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.H.Get())
		}, cycles: 4}
	case 0x7D:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.registers.L.Get())
		}, cycles: 4}
	case 0x7E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x40:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.B.Get())
		}, cycles: 4}
	case 0x41:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.C.Get())
		}, cycles: 4}
	case 0x42:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.D.Get())
		}, cycles: 4}
	case 0x43:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.E.Get())
		}, cycles: 4}
	case 0x44:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.H.Get())
		}, cycles: 4}
	case 0x45:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.L.Get())
		}, cycles: 4}
	case 0x46:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x48:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.B.Get())
		}, cycles: 4}
	case 0x49:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.C.Get())
		}, cycles: 4}
	case 0x4A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.D.Get())
		}, cycles: 4}
	case 0x4B:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.E.Get())
		}, cycles: 4}
	case 0x4C:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.H.Get())
		}, cycles: 4}
	case 0x4D:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.L.Get())
		}, cycles: 4}
	case 0x4E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x50:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.B.Get())
		}, cycles: 4}
	case 0x51:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.C.Get())
		}, cycles: 4}
	case 0x52:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.D.Get())
		}, cycles: 4}
	case 0x53:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.E.Get())
		}, cycles: 4}
	case 0x54:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.H.Get())
		}, cycles: 4}
	case 0x55:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.L.Get())
		}, cycles: 4}
	case 0x56:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x58:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.B.Get())
		}, cycles: 4}
	case 0x59:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.C.Get())
		}, cycles: 4}
	case 0x5A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.D.Get())
		}, cycles: 4}
	case 0x5B:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.E.Get())
		}, cycles: 4}
	case 0x5C:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.H.Get())
		}, cycles: 4}
	case 0x5D:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.L.Get())
		}, cycles: 4}
	case 0x5E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x60:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.registers.B.Get())
		}, cycles: 4}
	case 0x61:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.registers.C.Get())
		}, cycles: 4}
	case 0x62:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.D.Get())
		}, cycles: 4}
	case 0x63:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.registers.E.Get())
		}, cycles: 4}
	case 0x64:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.registers.H.Get())
		}, cycles: 4}
	case 0x65:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.registers.L.Get())
		}, cycles: 4}
	case 0x66:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x68:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.B.Get())
		}, cycles: 4}
	case 0x69:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.C.Get())
		}, cycles: 4}
	case 0x6A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.D.Get())
		}, cycles: 4}
	case 0x6B:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.E.Get())
		}, cycles: 4}
	case 0x6C:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.H.Get())
		}, cycles: 4}
	case 0x6D:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.L.Get())
		}, cycles: 4}
	case 0x6E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0x70:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.B.Get())
		}, cycles: 8}
	case 0x71:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.C.Get())
		}, cycles: 8}
	case 0x72:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.D.Get())
		}, cycles: 8}
	case 0x73:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.E.Get())
		}, cycles: 8}
	case 0x74:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.H.Get())
		}, cycles: 8}
	case 0x75:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.L.Get())
		}, cycles: 8}
	case 0x36:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.FetchNextByte())
		}, cycles: 12}
	case 0x0A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(c.registers.BC.Get()))
		}, cycles: 8}
	case 0x1A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(c.registers.DE.Get()))
		}, cycles: 8}
	case 0xFA:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(c.FetchNextWord()))
		}, cycles: 16}
	case 0x3E:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.FetchNextByte())
		}, cycles: 8}
	case 0x47:
		return Instruction{exec: func() {
			c.LDr8(c.registers.B, c.registers.A.Get())
		}, cycles: 4}
	case 0x4F:
		return Instruction{exec: func() {
			c.LDr8(c.registers.C, c.registers.A.Get())
		}, cycles: 4}
	case 0x57:
		return Instruction{exec: func() {
			c.LDr8(c.registers.D, c.registers.A.Get())
		}, cycles: 4}
	case 0x5F:
		return Instruction{exec: func() {
			c.LDr8(c.registers.E, c.registers.A.Get())
		}, cycles: 4}
	case 0x67:
		return Instruction{exec: func() {
			c.LDr8(c.registers.H, c.registers.A.Get())
		}, cycles: 4}
	case 0x6F:
		return Instruction{exec: func() {
			c.LDr8(c.registers.L, c.registers.A.Get())
		}, cycles: 4}
	case 0x02:
		return Instruction{exec: func() {
			c.LDm8(c.registers.BC.Get(), c.registers.A.Get())
		}, cycles: 8}
	case 0x12:
		return Instruction{exec: func() {
			c.LDm8(c.registers.DE.Get(), c.registers.A.Get())
		}, cycles: 8}
	case 0x77:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.A.Get())
		}, cycles: 8}
	case 0xEA:
		return Instruction{exec: func() {
			c.LDm8(c.FetchNextWord(), c.registers.A.Get())
		}, cycles: 16}
	case 0xF2:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(0xFF00 + types.Word(c.registers.C.Get())))
		}, cycles: 8}
	case 0xE2:
		return Instruction{exec: func() {
			c.LDm8(0xFF00 + types.Word(c.registers.C.Get()), c.registers.A.Get())
		}, cycles: 8}
	case 0x3A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(c.registers.HL.Get()))
			c.DEC16(&c.registers.HL)
		}, cycles: 8}
	case 0x32:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.A.Get())
			c.DEC16(&c.registers.HL)
		}, cycles: 8}
	case 0x2A:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(c.registers.HL.Get()))
			c.INC16(&c.registers.HL)
		}, cycles: 8}
	case 0x22:
		return Instruction{exec: func() {
			c.LDm8(c.registers.HL.Get(), c.registers.A.Get())
			c.INC16(&c.registers.HL)
		}, cycles: 8}
	case 0xE0:
		return Instruction{exec: func() {
			c.LDm8(0xFF00 + types.Word(c.FetchNextByte()), c.registers.A.Get())
		}, cycles: 12}
	case 0xF0:
		return Instruction{exec: func() {
			c.LDr8(c.registers.A, c.mmu.Get(0xFF00 + types.Word(c.FetchNextByte())))
		}, cycles: 12}
	case 0x01:
		return Instruction{exec: func() {
			c.LDr16(&c.registers.BC, c.FetchNextWord())
		}, cycles: 12}
	case 0x11:
		return Instruction{exec: func() {
			c.LDr16(&c.registers.DE, c.FetchNextWord())
		}, cycles: 12}
	case 0x21:
		return Instruction{exec: func() {
			c.LDr16(&c.registers.HL, c.FetchNextWord())
		}, cycles: 12}
	case 0x31:
		return Instruction{exec: func() {
			c.LDr16(&c.registers.SP, c.FetchNextWord())
		}, cycles: 12}
	case 0xF9:
		return Instruction{exec: func() {
			c.LDr16(&c.registers.SP, c.registers.HL.Get())
		}, cycles: 8}
	case 0xF8:
		return Instruction{exec: func() {
			var n = c.FetchNextByte()

			// @todo what is this?
			if n > 127 {
				c.LDr16(&c.registers.HL, c.registers.SP.Get() - types.Word(-n))
			} else {
				c.LDr16(&c.registers.HL, c.registers.SP.Get() + types.Word(n))
			}

			var check = c.registers.SP.Get() ^ types.Word(n) ^ ((c.registers.SP.Get() + types.Word(n)) & 0xffff)

			// @todo is this the carry alg?
			if (check & 0x100) == 0x100 {
				c.registers.Flags.Set(registers.C)
			} else {
				c.registers.Flags.Reset(registers.C)
			}

			// @todo is this the half carry alg?
			if (check & 0x10) == 0x10 {
				c.registers.Flags.Set(registers.H)
			} else {
				c.registers.Flags.Reset(registers.H)
			}

			c.registers.Flags.Reset(registers.Z)
			c.registers.Flags.Reset(registers.N)
		}, cycles: 12}
	case 0x08:
		return Instruction{exec: func() {
			h, l := types.WordToBytes(c.registers.SP.Get())
			addr := c.FetchNextWord()

			c.mmu.Set(addr+1, h)
			c.mmu.Set(addr, l)
		}, cycles: 20}
	case 0xF5:
		return Instruction{exec: func() {
			c.PushWord(c.registers.AF.Get())
		}, cycles: 16}
	case 0xC5:
		return Instruction{exec: func() {
			c.PushWord(c.registers.BC.Get())
		}, cycles: 16}
	case 0xD5:
		return Instruction{exec: func() {
			c.PushWord(c.registers.DE.Get())
		}, cycles: 16}
	case 0xE5:
		return Instruction{exec: func() {
			c.PushWord(c.registers.HL.Get())
		}, cycles: 16}
	case 0xF1:
		return Instruction{exec: func() {
			c.PopWord(&c.registers.AF)
		}, cycles: 12}
	case 0xC1:
		return Instruction{exec: func() {
			c.PopWord(&c.registers.BC)
		}, cycles: 12}
	case 0xD1:
		return Instruction{exec: func() {
			c.PopWord(&c.registers.DE)
		}, cycles: 12}
	case 0xE1:
		return Instruction{exec: func() {
			c.PopWord(&c.registers.HL)
		}, cycles: 12}
	// ALU start
	case 0x87:
		return Instruction{exec: func() {
			c.ADD8(c.registers.A.Get())
		}, cycles: 4}
	case 0x80:
		return Instruction{exec: func() {
			c.ADD8(c.registers.B.Get())
		}, cycles: 4}
	case 0x81:
		return Instruction{exec: func() {
			c.ADD8(c.registers.C.Get())
		}, cycles: 4}
	case 0x82:
		return Instruction{exec: func() {
			c.ADD8(c.registers.D.Get())
		}, cycles: 4}
	case 0x83:
		return Instruction{exec: func() {
			c.ADD8(c.registers.E.Get())
		}, cycles: 4}
	case 0x84:
		return Instruction{exec: func() {
			c.ADD8(c.registers.H.Get())
		}, cycles: 4}
	case 0x85:
		return Instruction{exec: func() {
			c.ADD8(c.registers.L.Get())
		}, cycles: 4}
	case 0x86:
		return Instruction{exec: func() {
			c.ADD8(c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0xC6:
		return Instruction{exec: func() {
			c.ADD8(c.FetchNextByte())
		}, cycles: 8}
	case 0x8F:
		return Instruction{exec: func() {
			c.ADD8(c.registers.A.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x88:
		return Instruction{exec: func() {
			c.ADD8(c.registers.B.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x89:
		return Instruction{exec: func() {
			c.ADD8(c.registers.C.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x8A:
		return Instruction{exec: func() {
			c.ADD8(c.registers.D.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x8B:
		return Instruction{exec: func() {
			c.ADD8(c.registers.E.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x8C:
		return Instruction{exec: func() {
			c.ADD8(c.registers.H.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x8D:
		return Instruction{exec: func() {
			c.ADD8(c.registers.L.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x8E:
		return Instruction{exec: func() {
			c.ADD8(c.mmu.Get(c.registers.HL.Get()) + c.registers.Flags.Get(registers.C))
		}, cycles: 8}
	case 0xCE:
		return Instruction{exec: func() {
			c.ADD8(c.FetchNextByte() + c.registers.Flags.Get(registers.C))
		}, cycles: 8}
	case 0x97:
		return Instruction{exec: func() {
			c.SUB8(c.registers.A.Get())
		}, cycles: 4}
	case 0x90:
		return Instruction{exec: func() {
			c.SUB8(c.registers.B.Get())
		}, cycles: 4}
	case 0x91:
		return Instruction{exec: func() {
			c.SUB8(c.registers.C.Get())
		}, cycles: 4}
	case 0x92:
		return Instruction{exec: func() {
			c.SUB8(c.registers.D.Get())
		}, cycles: 4}
	case 0x93:
		return Instruction{exec: func() {
			c.SUB8(c.registers.E.Get())
		}, cycles: 4}
	case 0x94:
		return Instruction{exec: func() {
			c.SUB8(c.registers.H.Get())
		}, cycles: 4}
	case 0x95:
		return Instruction{exec: func() {
			c.SUB8(c.registers.L.Get())
		}, cycles: 4}
	case 0x96:
		return Instruction{exec: func() {
			c.SUB8(c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0xD6:
		return Instruction{exec: func() {
			c.SUB8(c.FetchNextByte())
		}, cycles: 8}
	case 0x9F:
		return Instruction{exec: func() {
			c.SUB8(c.registers.A.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x98:
		return Instruction{exec: func() {
			c.SUB8(c.registers.B.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x99:
		return Instruction{exec: func() {
			c.SUB8(c.registers.C.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x9A:
		return Instruction{exec: func() {
			c.SUB8(c.registers.D.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x9B:
		return Instruction{exec: func() {
			c.SUB8(c.registers.E.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x9C:
		return Instruction{exec: func() {
			c.SUB8(c.registers.H.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x9D:
		return Instruction{exec: func() {
			c.SUB8(c.registers.L.Get() + c.registers.Flags.Get(registers.C))
		}, cycles: 4}
	case 0x9E:
		return Instruction{exec: func() {
			c.SUB8(c.mmu.Get(c.registers.HL.Get()) + c.registers.Flags.Get(registers.C))
		}, cycles: 8}
	case 0xDE:
		return Instruction{exec: func() {
			c.SUB8(c.FetchNextByte() + c.registers.Flags.Get(registers.C))
		}, cycles: 8}
	case 0xA7:
		return Instruction{exec: func() {
			c.AND8(c.registers.A.Get())
		}, cycles: 4}
	case 0xA0:
		return Instruction{exec: func() {
			c.AND8(c.registers.B.Get())
		}, cycles: 4}
	case 0xA1:
		return Instruction{exec: func() {
			c.AND8(c.registers.C.Get())
		}, cycles: 4}
	case 0xA2:
		return Instruction{exec: func() {
			c.AND8(c.registers.D.Get())
		}, cycles: 4}
	case 0xA3:
		return Instruction{exec: func() {
			c.AND8(c.registers.E.Get())
		}, cycles: 4}
	case 0xA4:
		return Instruction{exec: func() {
			c.AND8(c.registers.H.Get())
		}, cycles: 4}
	case 0xA5:
		return Instruction{exec: func() {
			c.AND8(c.registers.L.Get())
		}, cycles: 4}
	case 0xA6:
		return Instruction{exec: func() {
			c.AND8(c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0xE6:
		return Instruction{exec: func() {
			c.AND8(c.FetchNextByte())
		}, cycles: 8}
	case 0xB7:
		return Instruction{exec: func() {
			c.OR8(c.registers.A.Get())
		}, cycles: 4}
	case 0xB0:
		return Instruction{exec: func() {
			c.OR8(c.registers.B.Get())
		}, cycles: 4}
	case 0xB1:
		return Instruction{exec: func() {
			c.OR8(c.registers.C.Get())
		}, cycles: 4}
	case 0xB2:
		return Instruction{exec: func() {
			c.OR8(c.registers.D.Get())
		}, cycles: 4}
	case 0xB3:
		return Instruction{exec: func() {
			c.OR8(c.registers.E.Get())
		}, cycles: 4}
	case 0xB4:
		return Instruction{exec: func() {
			c.OR8(c.registers.H.Get())
		}, cycles: 4}
	case 0xB5:
		return Instruction{exec: func() {
			c.OR8(c.registers.L.Get())
		}, cycles: 4}
	case 0xB6:
		return Instruction{exec: func() {
			c.OR8(c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0xF6:
		return Instruction{exec: func() {
			c.OR8(c.FetchNextByte())
		}, cycles: 8}
	case 0xAF:
		return Instruction{exec: func() {
			c.XOR8(c.registers.A.Get())
		}, cycles: 4}
	case 0xA8:
		return Instruction{exec: func() {
			c.XOR8(c.registers.B.Get())
		}, cycles: 4}
	case 0xA9:
		return Instruction{exec: func() {
			c.XOR8(c.registers.C.Get())
		}, cycles: 4}
	case 0xAA:
		return Instruction{exec: func() {
			c.XOR8(c.registers.D.Get())
		}, cycles: 4}
	case 0xAB:
		return Instruction{exec: func() {
			c.XOR8(c.registers.E.Get())
		}, cycles: 4}
	case 0xAC:
		return Instruction{exec: func() {
			c.XOR8(c.registers.H.Get())
		}, cycles: 4}
	case 0xAD:
		return Instruction{exec: func() {
			c.XOR8(c.registers.L.Get())
		}, cycles: 4}
	case 0xAE:
		return Instruction{exec: func() {
			c.XOR8(c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0xEE:
		return Instruction{exec: func() {
			c.XOR8(c.FetchNextByte())
		}, cycles: 8}
	case 0xBF:
		return Instruction{exec: func() {
			c.CP8(c.registers.A.Get())
		}, cycles: 4}
	case 0xB8:
		return Instruction{exec: func() {
			c.CP8(c.registers.B.Get())
		}, cycles: 4}
	case 0xB9:
		return Instruction{exec: func() {
			c.CP8(c.registers.C.Get())
		}, cycles: 4}
	case 0xBA:
		return Instruction{exec: func() {
			c.CP8(c.registers.D.Get())
		}, cycles: 4}
	case 0xBB:
		return Instruction{exec: func() {
			c.CP8(c.registers.E.Get())
		}, cycles: 4}
	case 0xBC:
		return Instruction{exec: func() {
			c.CP8(c.registers.H.Get())
		}, cycles: 4}
	case 0xBD:
		return Instruction{exec: func() {
			c.CP8(c.registers.L.Get())
		}, cycles: 4}
	case 0xBE:
		return Instruction{exec: func() {
			c.CP8(c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 8}
	case 0xFE:
		return Instruction{exec: func() {
			c.CP8(c.FetchNextByte())
		}, cycles: 8}
	case 0x3C:
		return Instruction{exec: func() {
			c.INCr8(c.registers.A)
		}, cycles: 4}
	case 0x04:
		return Instruction{exec: func() {
			c.INCr8(c.registers.B)
		}, cycles: 4}
	case 0x0C:
		return Instruction{exec: func() {
			c.INCr8(c.registers.C)
		}, cycles: 4}
	case 0x14:
		return Instruction{exec: func() {
			c.INCr8(c.registers.D)
		}, cycles: 4}
	case 0x1C:
		return Instruction{exec: func() {
			c.INCr8(c.registers.E)
		}, cycles: 4}
	case 0x24:
		return Instruction{exec: func() {
			c.INCr8(c.registers.H)
		}, cycles: 4}
	case 0x2C:
		return Instruction{exec: func() {
			c.INCr8(c.registers.L)
		}, cycles: 4}
	case 0x34:
		return Instruction{exec: func() {
			c.INCm8(c.registers.HL.Get())
		}, cycles: 12}
	case 0x3D:
		return Instruction{exec: func() {
			c.DECr8(c.registers.A)
		}, cycles: 4}
	case 0x05:
		return Instruction{exec: func() {
			c.DECr8(c.registers.B)
		}, cycles: 4}
	case 0x0D:
		return Instruction{exec: func() {
			c.DECr8(c.registers.C)
		}, cycles: 4}
	case 0x15:
		return Instruction{exec: func() {
			c.DECr8(c.registers.D)
		}, cycles: 4}
	case 0x1D:
		return Instruction{exec: func() {
			c.DECr8(c.registers.E)
		}, cycles: 4}
	case 0x25:
		return Instruction{exec: func() {
			c.DECr8(c.registers.H)
		}, cycles: 4}
	case 0x2D:
		return Instruction{exec: func() {
			c.DECr8(c.registers.L)
		}, cycles: 4}
	case 0x35:
		return Instruction{exec: func() {
			c.DECm8(c.registers.HL.Get())
		}, cycles: 12}
	case 0x09:
		return Instruction{exec: func() {
			c.ADD16(c.registers.BC.Get())
		}, cycles: 8}
	case 0x19:
		return Instruction{exec: func() {
			c.ADD16(c.registers.DE.Get())
		}, cycles: 8}
	case 0x29:
		return Instruction{exec: func() {
			c.ADD16(c.registers.HL.Get())
		}, cycles: 8}
	case 0x39:
		return Instruction{exec: func() {
			c.ADD16(c.registers.SP.Get())
		}, cycles: 8}
	case 0xE8:
		return Instruction{exec: func() {
			var n = c.FetchNextByte()

			// @todo what is this?
			if n > 127 {
				c.LDr16(&c.registers.SP, c.registers.SP.Get() - types.Word(-n))
			} else {
				c.LDr16(&c.registers.SP, c.registers.SP.Get() + types.Word(n))
			}

			var check = c.registers.SP.Get() ^ types.Word(n) ^ ((c.registers.SP.Get() + types.Word(n)) & 0xffff)

			// @todo is this the carry alg?
			if (check & 0x100) == 0x100 {
				c.registers.Flags.Set(registers.C)
			} else {
				c.registers.Flags.Reset(registers.C)
			}

			// @todo is this the half carry alg?
			if (check & 0x10) == 0x10 {
				c.registers.Flags.Set(registers.H)
			} else {
				c.registers.Flags.Reset(registers.H)
			}

			c.registers.Flags.Reset(registers.Z)
			c.registers.Flags.Reset(registers.N)
		}, cycles: 16}
	case 0x03:
		return Instruction{exec: func() {
			c.INC16(&c.registers.BC)
		}, cycles: 8}
	case 0x13:
		return Instruction{exec: func() {
			c.INC16(&c.registers.DE)
		}, cycles: 8}
	case 0x23:
		return Instruction{exec: func() {
			c.INC16(&c.registers.HL)
		}, cycles: 8}
	case 0x33:
		return Instruction{exec: func() {
			c.INC16(&c.registers.SP)
		}, cycles: 8}
	case 0x0B:
		return Instruction{exec: func() {
			c.DEC16(&c.registers.BC)
		}, cycles: 8}
	case 0x1B:
		return Instruction{exec: func() {
			c.DEC16(&c.registers.DE)
		}, cycles: 8}
	case 0x2B:
		return Instruction{exec: func() {
			c.DEC16(&c.registers.HL)
		}, cycles: 8}
	case 0x3B:
		return Instruction{exec: func() {
			c.DEC16(&c.registers.SP)
		}, cycles: 8}
	// ALU end
	case 0x27:
		return Instruction{exec: c.DAA, cycles: 4}
	case 0x2F:
		return Instruction{exec: c.CPL, cycles: 4}
	case 0x3F:
		return Instruction{exec: c.CCF, cycles: 4}
	case 0x37:
		return Instruction{exec: c.SCF, cycles: 4}
	case 0x00:
		return Instruction{exec: c.NOP, cycles: 4}
	case 0x76:
		return Instruction{exec: c.HALT, cycles: 4}
	case 0x10:
		return Instruction{exec: c.STOP, cycles: 4}
	case 0xF3:
		return Instruction{exec: c.DI, cycles: 4}
	case 0xFB:
		return Instruction{exec: c.EI, cycles: 4}
	case 0x07:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.A)
		}, cycles: 4}
	case 0x17:
		return Instruction{exec: func() {
			c.RLr8(c.registers.A)
		}, cycles: 4}
	case 0x0F:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.A)
		}, cycles: 4}
	case 0x1F:
		return Instruction{exec: func() {
			c.RRr8(c.registers.A)
		}, cycles: 4}
	case 0xC3:
		return Instruction{exec: func() {
			c.JP(c.FetchNextWord())
		}, cycles: 12}
	case 0xC2:
		return Instruction{exec: func() {
			c.JPc(registers.N, 0x0, c.FetchNextWord())
		}, cycles: 12}
	case 0xCA:
		return Instruction{exec: func() {
			c.JPc(registers.Z, 0x1, c.FetchNextWord())
		}, cycles: 12}
	case 0xD2:
		return Instruction{exec: func() {
			c.JPc(registers.C, 0x0, c.FetchNextWord())
		}, cycles: 12}
	case 0xDA:
		return Instruction{exec: func() {
			c.JPc(registers.C, 0x1, c.FetchNextWord())
		}, cycles: 12}
	case 0xE9:
		return Instruction{exec: func() {
			c.JP(c.registers.HL.Get())
		}, cycles: 4}
	case 0x18:
		return Instruction{exec: func() {
			c.JR(c.FetchNextByte())
		}, cycles: 8}
	case 0x20:
		return Instruction{exec: func() {
			c.JRc(registers.N, 0x0, c.FetchNextByte())
		}, cycles: 8}
	case 0x28:
		return Instruction{exec: func() {
			c.JRc(registers.Z, 0x1, c.FetchNextByte())
		}, cycles: 8}
	case 0x30:
		return Instruction{exec: func() {
			c.JRc(registers.C, 0x0, c.FetchNextByte())
		}, cycles: 8}
	case 0x38:
		return Instruction{exec: func() {
			c.JRc(registers.C, 0x1, c.FetchNextByte())
		}, cycles: 8}
	case 0xCD:
		return Instruction{exec: func() {
			c.CALL(c.FetchNextWord())
		}, cycles: 12}
	case 0xC4:
		return Instruction{exec: func() {
			c.CALLc(registers.N, 0x0, c.FetchNextWord())
		}, cycles: 12}
	case 0xCC:
		return Instruction{exec: func() {
			c.CALLc(registers.Z, 0x1, c.FetchNextWord())
		}, cycles: 12}
	case 0xD4:
		return Instruction{exec: func() {
			c.CALLc(registers.C, 0x0, c.FetchNextWord())
		}, cycles: 12}
	case 0xDC:
		return Instruction{exec: func() {
			c.CALLc(registers.C, 0x1, c.FetchNextWord())
		}, cycles: 12}
	case 0xC7:
		return Instruction{exec: func() {
			c.RST(0x0)
		}, cycles: 16}
	case 0xCF:
		return Instruction{exec: func() {
			c.RST(0x8)
		}, cycles: 16}
	case 0xD7:
		return Instruction{exec: func() {
			c.RST(0x10)
		}, cycles: 16}
	case 0xDF:
		return Instruction{exec: func() {
			c.RST(0x18)
		}, cycles: 16}
	case 0xE7:
		return Instruction{exec: func() {
			c.RST(0x20)
		}, cycles: 16}
	case 0xEF:
		return Instruction{exec: func() {
			c.RST(0x28)
		}, cycles: 16}
	case 0xF7:
		return Instruction{exec: func() {
			c.RST(0x30)
		}, cycles: 16}
	case 0xFF:
		return Instruction{exec: func() {
			c.RST(0x38)
		}, cycles: 16}
	case 0xC9:
		return Instruction{exec: func() {
			c.RET()
		}, cycles: 8}
	case 0xC0:
		return Instruction{exec: func() {
			c.RETc(registers.N, 0x0)
		}, cycles: 8}
	case 0xC8:
		return Instruction{exec: func() {
			c.RETc(registers.Z, 0x1)
		}, cycles: 8}
	case 0xD0:
		return Instruction{exec: func() {
			c.RETc(registers.C, 0x0)
		}, cycles: 8}
	case 0xD8:
		return Instruction{exec: func() {
			c.RETc(registers.C, 0x1)
		}, cycles: 8}
	case 0xD9:
		return Instruction{exec: c.RETI, cycles: 8}
	case 0xCB:
		return Instruction{exec: func() {
			opcode := c.FetchNextByte()
			inst := c.DecodeCB(opcode)
			inst.exec()
		}, cycles: 4}
	default:
		log.Fatalf("OpCode 0x%02X not implemented", opcode)
		return Instruction{}
	}
}

func (c *CPU) DecodeCB(opcode types.Byte) Instruction {
	switch opcode {
	case 0x37:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.A)
		}, cycles: 8}
	case 0x30:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.B)
		}, cycles: 8}
	case 0x31:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.C)
		}, cycles: 8}
	case 0x32:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.D)
		}, cycles: 8}
	case 0x33:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.E)
		}, cycles: 8}
	case 0x34:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.H)
		}, cycles: 8}
	case 0x35:
		return Instruction{exec: func() {
			c.SWAPr8(c.registers.L)
		}, cycles: 8}
	case 0x36:
		return Instruction{exec: func() {
			c.SWAPm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x07:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.A)
		}, cycles: 8}
	case 0x00:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.B)
		}, cycles: 8}
	case 0x01:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.C)
		}, cycles: 8}
	case 0x02:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.D)
		}, cycles: 8}
	case 0x03:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.E)
		}, cycles: 8}
	case 0x04:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.H)
		}, cycles: 8}
	case 0x05:
		return Instruction{exec: func() {
			c.RLCr8(c.registers.L)
		}, cycles: 8}
	case 0x06:
		return Instruction{exec: func() {
			c.RLCm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x17:
		return Instruction{exec: func() {
			c.RLr8(c.registers.A)
		}, cycles: 8}
	case 0x10:
		return Instruction{exec: func() {
			c.RLr8(c.registers.B)
		}, cycles: 8}
	case 0x11:
		return Instruction{exec: func() {
			c.RLr8(c.registers.C)
		}, cycles: 8}
	case 0x12:
		return Instruction{exec: func() {
			c.RLr8(c.registers.D)
		}, cycles: 8}
	case 0x13:
		return Instruction{exec: func() {
			c.RLr8(c.registers.E)
		}, cycles: 8}
	case 0x14:
		return Instruction{exec: func() {
			c.RLr8(c.registers.H)
		}, cycles: 8}
	case 0x15:
		return Instruction{exec: func() {
			c.RLr8(c.registers.L)
		}, cycles: 8}
	case 0x16:
		return Instruction{exec: func() {
			c.RLm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x0F:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.A)
		}, cycles: 8}
	case 0x08:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.B)
		}, cycles: 8}
	case 0x09:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.C)
		}, cycles: 8}
	case 0x0A:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.D)
		}, cycles: 8}
	case 0x0B:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.E)
		}, cycles: 8}
	case 0x0C:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.H)
		}, cycles: 8}
	case 0x0D:
		return Instruction{exec: func() {
			c.RRCr8(c.registers.L)
		}, cycles: 8}
	case 0x0E:
		return Instruction{exec: func() {
			c.RRCm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x1F:
		return Instruction{exec: func() {
			c.RRr8(c.registers.A)
		}, cycles: 8}
	case 0x18:
		return Instruction{exec: func() {
			c.RRr8(c.registers.B)
		}, cycles: 8}
	case 0x19:
		return Instruction{exec: func() {
			c.RRr8(c.registers.C)
		}, cycles: 8}
	case 0x1A:
		return Instruction{exec: func() {
			c.RRr8(c.registers.D)
		}, cycles: 8}
	case 0x1B:
		return Instruction{exec: func() {
			c.RRr8(c.registers.E)
		}, cycles: 8}
	case 0x1C:
		return Instruction{exec: func() {
			c.RRr8(c.registers.H)
		}, cycles: 8}
	case 0x1D:
		return Instruction{exec: func() {
			c.RRr8(c.registers.L)
		}, cycles: 8}
	case 0x1E:
		return Instruction{exec: func() {
			c.RRm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x27:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.A)
		}, cycles: 8}
	case 0x20:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.B)
		}, cycles: 8}
	case 0x21:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.C)
		}, cycles: 8}
	case 0x22:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.D)
		}, cycles: 8}
	case 0x23:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.E)
		}, cycles: 8}
	case 0x24:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.H)
		}, cycles: 8}
	case 0x25:
		return Instruction{exec: func() {
			c.SLAr8(c.registers.L)
		}, cycles: 8}
	case 0x26:
		return Instruction{exec: func() {
			c.SLAm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x2F:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.A)
		}, cycles: 8}
	case 0x28:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.B)
		}, cycles: 8}
	case 0x29:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.C)
		}, cycles: 8}
	case 0x2A:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.D)
		}, cycles: 8}
	case 0x2B:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.E)
		}, cycles: 8}
	case 0x2C:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.H)
		}, cycles: 8}
	case 0x2D:
		return Instruction{exec: func() {
			c.SRAr8(c.registers.L)
		}, cycles: 8}
	case 0x2E:
		return Instruction{exec: func() {
			c.SRAm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x3F:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.A)
		}, cycles: 8}
	case 0x38:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.B)
		}, cycles: 8}
	case 0x39:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.C)
		}, cycles: 8}
	case 0x3A:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.D)
		}, cycles: 8}
	case 0x3B:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.E)
		}, cycles: 8}
	case 0x3C:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.H)
		}, cycles: 8}
	case 0x3D:
		return Instruction{exec: func() {
			c.SRLr8(c.registers.L)
		}, cycles: 8}
	case 0x3E:
		return Instruction{exec: func() {
			c.SRLm8(c.registers.HL.Get())
		}, cycles: 16}
	case 0x47:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.A.Get())
		}, cycles: 8}
	case 0x40:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.B.Get())
		}, cycles: 8}
	case 0x41:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.C.Get())
		}, cycles: 8}
	case 0x42:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.D.Get())
		}, cycles: 8}
	case 0x43:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.E.Get())
		}, cycles: 8}
	case 0x44:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.H.Get())
		}, cycles: 8}
	case 0x45:
		return Instruction{exec: func() {
			c.BIT(0, c.registers.L.Get())
		}, cycles: 8}
	case 0x46:
		return Instruction{exec: func() {
			c.BIT(0, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x48:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.A.Get())
		}, cycles: 8}
	case 0x49:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.B.Get())
		}, cycles: 8}
	case 0x4A:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.C.Get())
		}, cycles: 8}
	case 0x4B:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.D.Get())
		}, cycles: 8}
	case 0x4C:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.E.Get())
		}, cycles: 8}
	case 0x4D:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.H.Get())
		}, cycles: 8}
	case 0x4E:
		return Instruction{exec: func() {
			c.BIT(1, c.registers.L.Get())
		}, cycles: 8}
	case 0x4F:
		return Instruction{exec: func() {
			c.BIT(1, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x57:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.A.Get())
		}, cycles: 8}
	case 0x50:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.B.Get())
		}, cycles: 8}
	case 0x51:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.C.Get())
		}, cycles: 8}
	case 0x52:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.D.Get())
		}, cycles: 8}
	case 0x53:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.E.Get())
		}, cycles: 8}
	case 0x54:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.H.Get())
		}, cycles: 8}
	case 0x55:
		return Instruction{exec: func() {
			c.BIT(2, c.registers.L.Get())
		}, cycles: 8}
	case 0x56:
		return Instruction{exec: func() {
			c.BIT(2, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x58:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.A.Get())
		}, cycles: 8}
	case 0x59:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.B.Get())
		}, cycles: 8}
	case 0x5A:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.C.Get())
		}, cycles: 8}
	case 0x5B:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.D.Get())
		}, cycles: 8}
	case 0x5C:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.E.Get())
		}, cycles: 8}
	case 0x5D:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.H.Get())
		}, cycles: 8}
	case 0x5E:
		return Instruction{exec: func() {
			c.BIT(3, c.registers.L.Get())
		}, cycles: 8}
	case 0x5F:
		return Instruction{exec: func() {
			c.BIT(3, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x67:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.A.Get())
		}, cycles: 8}
	case 0x60:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.B.Get())
		}, cycles: 8}
	case 0x61:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.C.Get())
		}, cycles: 8}
	case 0x62:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.D.Get())
		}, cycles: 8}
	case 0x63:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.E.Get())
		}, cycles: 8}
	case 0x64:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.H.Get())
		}, cycles: 8}
	case 0x65:
		return Instruction{exec: func() {
			c.BIT(4, c.registers.L.Get())
		}, cycles: 8}
	case 0x66:
		return Instruction{exec: func() {
			c.BIT(4, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x68:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.A.Get())
		}, cycles: 8}
	case 0x69:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.B.Get())
		}, cycles: 8}
	case 0x6A:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.C.Get())
		}, cycles: 8}
	case 0x6B:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.D.Get())
		}, cycles: 8}
	case 0x6C:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.E.Get())
		}, cycles: 8}
	case 0x6D:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.H.Get())
		}, cycles: 8}
	case 0x6E:
		return Instruction{exec: func() {
			c.BIT(5, c.registers.L.Get())
		}, cycles: 8}
	case 0x6F:
		return Instruction{exec: func() {
			c.BIT(5, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x77:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.A.Get())
		}, cycles: 8}
	case 0x70:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.B.Get())
		}, cycles: 8}
	case 0x71:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.C.Get())
		}, cycles: 8}
	case 0x72:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.D.Get())
		}, cycles: 8}
	case 0x73:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.E.Get())
		}, cycles: 8}
	case 0x74:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.H.Get())
		}, cycles: 8}
	case 0x75:
		return Instruction{exec: func() {
			c.BIT(6, c.registers.L.Get())
		}, cycles: 8}
	case 0x76:
		return Instruction{exec: func() {
			c.BIT(6, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0x78:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.A.Get())
		}, cycles: 8}
	case 0x79:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.B.Get())
		}, cycles: 8}
	case 0x7A:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.C.Get())
		}, cycles: 8}
	case 0x7B:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.D.Get())
		}, cycles: 8}
	case 0x7C:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.E.Get())
		}, cycles: 8}
	case 0x7D:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.H.Get())
		}, cycles: 8}
	case 0x7E:
		return Instruction{exec: func() {
			c.BIT(7, c.registers.L.Get())
		}, cycles: 8}
	case 0x7F:
		return Instruction{exec: func() {
			c.BIT(7, c.mmu.Get(c.registers.HL.Get()))
		}, cycles: 16}
	case 0xC7:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.A)
		}, cycles: 8}
	case 0xC0:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.B)
		}, cycles: 8}
	case 0xC1:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.C)
		}, cycles: 8}
	case 0xC2:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.D)
		}, cycles: 8}
	case 0xC3:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.E)
		}, cycles: 8}
	case 0xC4:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.H)
		}, cycles: 8}
	case 0xC5:
		return Instruction{exec: func() {
			c.SETr8(0, c.registers.L)
		}, cycles: 8}
	case 0xC6:
		return Instruction{exec: func() {
			c.SETm8(0, c.registers.HL.Get())
		}, cycles: 16}
	case 0xC8:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.A)
		}, cycles: 8}
	case 0xC9:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.B)
		}, cycles: 8}
	case 0xCA:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.C)
		}, cycles: 8}
	case 0xCB:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.D)
		}, cycles: 8}
	case 0xCC:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.E)
		}, cycles: 8}
	case 0xCD:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.H)
		}, cycles: 8}
	case 0xCE:
		return Instruction{exec: func() {
			c.SETr8(1, c.registers.L)
		}, cycles: 8}
	case 0xCF:
		return Instruction{exec: func() {
			c.SETm8(1, c.registers.HL.Get())
		}, cycles: 16}
	case 0xD7:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.A)
		}, cycles: 8}
	case 0xD0:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.B)
		}, cycles: 8}
	case 0xD1:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.C)
		}, cycles: 8}
	case 0xD2:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.D)
		}, cycles: 8}
	case 0xD3:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.E)
		}, cycles: 8}
	case 0xD4:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.H)
		}, cycles: 8}
	case 0xD5:
		return Instruction{exec: func() {
			c.SETr8(2, c.registers.L)
		}, cycles: 8}
	case 0xD6:
		return Instruction{exec: func() {
			c.SETm8(2, c.registers.HL.Get())
		}, cycles: 16}
	case 0xD8:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.A)
		}, cycles: 8}
	case 0xD9:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.B)
		}, cycles: 8}
	case 0xDA:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.C)
		}, cycles: 8}
	case 0xDB:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.D)
		}, cycles: 8}
	case 0xDC:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.E)
		}, cycles: 8}
	case 0xDD:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.H)
		}, cycles: 8}
	case 0xDE:
		return Instruction{exec: func() {
			c.SETr8(3, c.registers.L)
		}, cycles: 8}
	case 0xDF:
		return Instruction{exec: func() {
			c.SETm8(3, c.registers.HL.Get())
		}, cycles: 16}
	case 0xE7:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.A)
		}, cycles: 8}
	case 0xE0:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.B)
		}, cycles: 8}
	case 0xE1:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.C)
		}, cycles: 8}
	case 0xE2:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.D)
		}, cycles: 8}
	case 0xE3:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.E)
		}, cycles: 8}
	case 0xE4:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.H)
		}, cycles: 8}
	case 0xE5:
		return Instruction{exec: func() {
			c.SETr8(4, c.registers.L)
		}, cycles: 8}
	case 0xE6:
		return Instruction{exec: func() {
			c.SETm8(4, c.registers.HL.Get())
		}, cycles: 16}
	case 0xE8:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.A)
		}, cycles: 8}
	case 0xE9:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.B)
		}, cycles: 8}
	case 0xEA:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.C)
		}, cycles: 8}
	case 0xEB:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.D)
		}, cycles: 8}
	case 0xEC:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.E)
		}, cycles: 8}
	case 0xED:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.H)
		}, cycles: 8}
	case 0xEE:
		return Instruction{exec: func() {
			c.SETr8(5, c.registers.L)
		}, cycles: 8}
	case 0xEF:
		return Instruction{exec: func() {
			c.SETm8(5, c.registers.HL.Get())
		}, cycles: 16}
	case 0xF7:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.A)
		}, cycles: 8}
	case 0xF0:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.B)
		}, cycles: 8}
	case 0xF1:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.C)
		}, cycles: 8}
	case 0xF2:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.D)
		}, cycles: 8}
	case 0xF3:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.E)
		}, cycles: 8}
	case 0xF4:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.H)
		}, cycles: 8}
	case 0xF5:
		return Instruction{exec: func() {
			c.SETr8(6, c.registers.L)
		}, cycles: 8}
	case 0xF6:
		return Instruction{exec: func() {
			c.SETm8(6, c.registers.HL.Get())
		}, cycles: 16}
	case 0xF8:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.A)
		}, cycles: 8}
	case 0xF9:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.B)
		}, cycles: 8}
	case 0xFA:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.C)
		}, cycles: 8}
	case 0xFB:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.D)
		}, cycles: 8}
	case 0xFC:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.E)
		}, cycles: 8}
	case 0xFD:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.H)
		}, cycles: 8}
	case 0xFE:
		return Instruction{exec: func() {
			c.SETr8(7, c.registers.L)
		}, cycles: 8}
	case 0xFF:
		return Instruction{exec: func() {
			c.SETm8(7, c.registers.HL.Get())
		}, cycles: 16}
	case 0x87:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.A)
		}, cycles: 8}
	case 0x80:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.B)
		}, cycles: 8}
	case 0x81:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.C)
		}, cycles: 8}
	case 0x82:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.D)
		}, cycles: 8}
	case 0x83:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.E)
		}, cycles: 8}
	case 0x84:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.H)
		}, cycles: 8}
	case 0x85:
		return Instruction{exec: func() {
			c.RESr8(0, c.registers.L)
		}, cycles: 8}
	case 0x86:
		return Instruction{exec: func() {
			c.RESm8(0, c.registers.HL.Get())
		}, cycles: 16}
	case 0x88:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.A)
		}, cycles: 8}
	case 0x89:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.B)
		}, cycles: 8}
	case 0x8A:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.C)
		}, cycles: 8}
	case 0x8B:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.D)
		}, cycles: 8}
	case 0x8C:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.E)
		}, cycles: 8}
	case 0x8D:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.H)
		}, cycles: 8}
	case 0x8E:
		return Instruction{exec: func() {
			c.RESr8(1, c.registers.L)
		}, cycles: 8}
	case 0x8F:
		return Instruction{exec: func() {
			c.RESm8(1, c.registers.HL.Get())
		}, cycles: 16}
	case 0x97:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.A)
		}, cycles: 8}
	case 0x90:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.B)
		}, cycles: 8}
	case 0x91:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.C)
		}, cycles: 8}
	case 0x92:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.D)
		}, cycles: 8}
	case 0x93:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.E)
		}, cycles: 8}
	case 0x94:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.H)
		}, cycles: 8}
	case 0x95:
		return Instruction{exec: func() {
			c.RESr8(2, c.registers.L)
		}, cycles: 8}
	case 0x96:
		return Instruction{exec: func() {
			c.RESm8(2, c.registers.HL.Get())
		}, cycles: 16}
	case 0x98:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.A)
		}, cycles: 8}
	case 0x99:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.B)
		}, cycles: 8}
	case 0x9A:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.C)
		}, cycles: 8}
	case 0x9B:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.D)
		}, cycles: 8}
	case 0x9C:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.E)
		}, cycles: 8}
	case 0x9D:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.H)
		}, cycles: 8}
	case 0x9E:
		return Instruction{exec: func() {
			c.RESr8(3, c.registers.L)
		}, cycles: 8}
	case 0x9F:
		return Instruction{exec: func() {
			c.RESm8(3, c.registers.HL.Get())
		}, cycles: 16}
	case 0xA7:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.A)
		}, cycles: 8}
	case 0xA0:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.B)
		}, cycles: 8}
	case 0xA1:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.C)
		}, cycles: 8}
	case 0xA2:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.D)
		}, cycles: 8}
	case 0xA3:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.E)
		}, cycles: 8}
	case 0xA4:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.H)
		}, cycles: 8}
	case 0xA5:
		return Instruction{exec: func() {
			c.RESr8(4, c.registers.L)
		}, cycles: 8}
	case 0xA6:
		return Instruction{exec: func() {
			c.RESm8(4, c.registers.HL.Get())
		}, cycles: 16}
	case 0xA8:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.A)
		}, cycles: 8}
	case 0xA9:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.B)
		}, cycles: 8}
	case 0xAA:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.C)
		}, cycles: 8}
	case 0xAB:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.D)
		}, cycles: 8}
	case 0xAC:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.E)
		}, cycles: 8}
	case 0xAD:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.H)
		}, cycles: 8}
	case 0xAE:
		return Instruction{exec: func() {
			c.RESr8(5, c.registers.L)
		}, cycles: 8}
	case 0xAF:
		return Instruction{exec: func() {
			c.RESm8(5, c.registers.HL.Get())
		}, cycles: 16}
	case 0xB7:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.A)
		}, cycles: 8}
	case 0xB0:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.B)
		}, cycles: 8}
	case 0xB1:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.C)
		}, cycles: 8}
	case 0xB2:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.D)
		}, cycles: 8}
	case 0xB3:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.E)
		}, cycles: 8}
	case 0xB4:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.H)
		}, cycles: 8}
	case 0xB5:
		return Instruction{exec: func() {
			c.RESr8(6, c.registers.L)
		}, cycles: 8}
	case 0xB6:
		return Instruction{exec: func() {
			c.RESm8(6, c.registers.HL.Get())
		}, cycles: 16}
	case 0xB8:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.A)
		}, cycles: 8}
	case 0xB9:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.B)
		}, cycles: 8}
	case 0xBA:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.C)
		}, cycles: 8}
	case 0xBB:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.D)
		}, cycles: 8}
	case 0xBC:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.E)
		}, cycles: 8}
	case 0xBD:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.H)
		}, cycles: 8}
	case 0xBE:
		return Instruction{exec: func() {
			c.RESr8(7, c.registers.L)
		}, cycles: 8}
	case 0xBF:
		return Instruction{exec: func() {
			c.RESm8(7, c.registers.HL.Get())
		}, cycles: 16}
	default:
		log.Fatalf("OpCode 0xCB%02X not implemented", opcode)
		return Instruction{}
	}
}

// Load a byte into a register
func (c *CPU) LDr16(r registers.WordRegister, w types.Word)  {
	r.Set(w)
}

// Load a byte into a register
func (c *CPU) LDr8(r types.ByteRegister, b types.Byte)  {
	r.Set(b)
}

// Load a byte into a memory address
func (c *CPU) LDm8(address types.Word, b types.Byte)  {
	c.mmu.Set(address, b)
}

func (c *CPU) NOP() {
}

func (c *CPU) Run() {
	for {
		opcode := c.FetchNextByte()
		inst := c.Decode(opcode)
		inst.exec()
	}
}

func (c *CPU) FetchNextByte() types.Byte {
	b := c.mmu.Get(c.registers.PC.Get())
	c.INC16(&c.registers.PC)
	return b
}

func (c *CPU) FetchNextWord() types.Word {
	l := c.FetchNextByte()
	h := c.FetchNextByte()
	return types.WordFromBytes(h, l)
}

func (c *CPU) PushWord(w types.Word) {
	h, l := types.WordToBytes(w)
	c.PushByte(h)
	c.PushByte(l)
}

func (c *CPU) PushByte(b types.Byte) {
	c.DEC16(&c.registers.SP)
	c.mmu.Set(c.registers.SP.Get(), b)
}

func (c *CPU) DEC16(r registers.WordRegister) {
	r.Set(r.Get() - 0x1)
}

func (c *CPU) DECr8(r types.ByteRegister) {
	result := r.Get() - 0x1

	c.registers.Flags.Set(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (result^0x01^r.Get())&0x10 == 0x10 {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	r.Set(result)
}

func (c *CPU) DECm8(address types.Word) {
	result := c.mmu.Get(address) - 0x1

	c.registers.Flags.Set(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (result^0x01^ c.mmu.Get(address))&0x10 == 0x10 {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	c.mmu.Set(address, result)
}

func (c *CPU) INC16(r registers.WordRegister) {
	r.Set(r.Get() + 0x1)
}

func (c *CPU) INCr8(r types.ByteRegister) {
	result := r.Get() + 0x1

	c.registers.Flags.Reset(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (result^0x01^r.Get())&0x10 == 0x10 {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	r.Set(result)
}

func (c *CPU) INCm8(address types.Word) {
	result := c.mmu.Get(address) + 0x1

	c.registers.Flags.Reset(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (result^0x01^c.mmu.Get(address))&0x10 == 0x10 {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	c.mmu.Set(address, result)
}

func (c *CPU) PopWord(r registers.WordRegister) {
	l := c.mmu.Get(c.registers.SP.Get())
	c.INC16(&c.registers.SP)
	h := c.mmu.Get(c.registers.SP.Get())
	c.INC16(&c.registers.SP)

	r.Set(types.WordFromBytes(h, l))
}

func (c *CPU) PopByte(r types.ByteRegister) {
	r.Set(c.mmu.Get(c.registers.SP.Get()))
	c.INC16(&c.registers.SP)
}

func (c *CPU) ADD8(b types.Byte) {
	result := c.registers.A.Get() + b

	c.registers.Flags.Reset(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (result ^ b ^ c.registers.A.Get()) & 0x10 == 0x10 {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	if result < c.registers.A.Get() {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	c.registers.A.Set(result)
}

func (c *CPU) ADD16(b types.Word) {
	result := c.registers.HL.Get() + b

	c.registers.Flags.Reset(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (result ^ b ^ c.registers.HL.Get()) & 0x1000 == 0x1000 {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	if result < c.registers.HL.Get() {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	c.registers.HL.Set(result)
}

func (c *CPU) SUB8(b types.Byte) {
	result := c.registers.A.Get() - b

	c.registers.Flags.Set(registers.N)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	if (int(c.registers.A.Get()) & 0xF) < (int(b) & 0xF) {
		c.registers.Flags.Set(registers.H)
	} else {
		c.registers.Flags.Reset(registers.H)
	}

	if (int(c.registers.A.Get()) & 0xFF) < (int(b) & 0xFF) {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	c.registers.A.Set(result)
}

func (c *CPU) AND8(b types.Byte) {
	result := c.registers.A.Get() & b

	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.C)
	c.registers.Flags.Set(registers.H)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	c.registers.A.Set(result)
}

func (c *CPU) OR8(b types.Byte) {
	result := c.registers.A.Get() | b

	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.C)
	c.registers.Flags.Reset(registers.H)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	c.registers.A.Set(result)
}

func (c *CPU) XOR8(b types.Byte) {
	result := c.registers.A.Get() ^ b

	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.C)
	c.registers.Flags.Reset(registers.H)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}

	c.registers.A.Set(result)
}

func (c *CPU) CP8(b types.Byte) {
	result := c.registers.A.Get() - b

	// @todo

	c.registers.A.Set(result)
}

func (c *CPU) BIT(bit byte, b types.Byte) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Set(registers.H)
	if types.GetBit(bit, b) == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
}

func (c *CPU) SETr8(bit byte, r types.ByteRegister) {
	r.Set(types.SetBit(bit, r.Get()))
}

func (c *CPU) SETm8(bit byte, address types.Word) {
	c.mmu.Set(address, types.SetBit(bit, c.mmu.Get(address)))
}

func (c *CPU) RESr8(bit byte, r types.ByteRegister) {
	r.Set(types.ResetBit(bit, r.Get()))
}

func (c *CPU) RESm8(bit byte, address types.Word) {
	c.mmu.Set(address, types.ResetBit(bit, c.mmu.Get(address)))
}

func (c *CPU) SWAPr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)
	c.registers.Flags.Reset(registers.C)

	result := ((uint16(r.Get()) << 4) | (uint16(r.Get()) >> 4)) & 0xFF
	r.Set(types.Byte(result))

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
}

func (c *CPU) SWAPm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)
	c.registers.Flags.Reset(registers.C)

	result := ((uint16(c.mmu.Get(address)) << 4) | (uint16(c.mmu.Get(address)) >> 4)) & 0xFF
	c.mmu.Set(address, types.Byte(result))

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
}

func (c *CPU) CCF() {
	if c.registers.Flags.Get(registers.C) == 0x1 {
		c.registers.Flags.Reset(registers.C)
	} else {
		c.registers.Flags.Set(registers.C)
	}
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)
}

func (c *CPU) SCF() {
	c.registers.Flags.Set(registers.C)
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)
}

func (c *CPU) DAA() {
	// @todo
}

func (c *CPU) CPL() {
	// @todo
}

func (c *CPU) HALT() {
	// @todo
}

func (c *CPU) STOP() {
	// @todo
	c.HALT()
}

func (c *CPU) DI() {
	// @todo
}

func (c *CPU) EI() {
	// @todo
}

func (c *CPU) RLCr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	if r.Get() & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	}

	result := (r.Get() << 1) | c.registers.Flags.Get(registers.C)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) RLr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := (r.Get() << 1) | c.registers.Flags.Get(registers.C)

	if r.Get() & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) RRCr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	if r.Get() & 0x01 == 0x01 {
		c.registers.Flags.Set(registers.C)
	}

	result := (r.Get() >> 1) | (c.registers.Flags.Get(registers.C) << 7)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) RRr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := (r.Get() >> 1) | (c.registers.Flags.Get(registers.C) << 7)

	if r.Get() & 0x01 == 0x01 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) RLCm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	if c.mmu.Get(address) & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	}

	result := (c.mmu.Get(address) << 1) | c.registers.Flags.Get(registers.C)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) RLm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := (c.mmu.Get(address) << 1) | c.registers.Flags.Get(registers.C)

	if c.mmu.Get(address) & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) RRCm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	if c.mmu.Get(address) & 0x01 == 0x01 {
		c.registers.Flags.Set(registers.C)
	}

	result := (c.mmu.Get(address) >> 1) | (c.registers.Flags.Get(registers.C) << 7)

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) RRm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := (c.mmu.Get(address) >> 1) | (c.registers.Flags.Get(registers.C) << 7)

	if c.mmu.Get(address) & 0x01 == 0x01 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) SLAr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := r.Get() << 1

	if r.Get() & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) SLAm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := c.mmu.Get(address) << 1

	if c.mmu.Get(address) & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) SRAr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := (r.Get() >> 1) | (r.Get() & 0x80)

	if r.Get() & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) SRAm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := (c.mmu.Get(address) >> 1) | (c.mmu.Get(address) & 0x80)

	if c.mmu.Get(address) & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) SRLr8(r types.ByteRegister) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := r.Get() >> 1

	if r.Get() & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	r.Set(result)
}

func (c *CPU) SRLm8(address types.Word) {
	c.registers.Flags.Reset(registers.N)
	c.registers.Flags.Reset(registers.H)

	result := c.mmu.Get(address) >> 1

	if c.mmu.Get(address) & 0x80 == 0x80 {
		c.registers.Flags.Set(registers.C)
	} else {
		c.registers.Flags.Reset(registers.C)
	}

	if result == 0x0 {
		c.registers.Flags.Set(registers.Z)
	} else {
		c.registers.Flags.Reset(registers.Z)
	}
	c.mmu.Set(address, result)
}

func (c *CPU) JP(address types.Word) {
	c.registers.PC.Set(address)
}

func (c *CPU) JPc(flag int, flagValue types.Byte, address types.Word) {
	if c.registers.Flags.Get(flag) == flagValue {
		c.JP(address)
	}
}

func (c *CPU) JR(b types.Byte) {
	c.JP(c.registers.PC.Get() + types.Word(b))
}

func (c *CPU) JRc(flag int, flagValue types.Byte, b types.Byte) {
	if c.registers.Flags.Get(flag) == flagValue {
		c.JR(b)
	}
}

func (c *CPU) CALL(address types.Word) {
	c.PushWord(c.registers.PC.Get() + 0x3)
	c.JP(address)
}

func (c *CPU) CALLc(flag int, flagValue types.Byte, address types.Word) {
	if c.registers.Flags.Get(flag) == flagValue {
		c.CALL(address)
	}
}

func (c *CPU) RST(address types.Word) {
	c.PushWord(c.registers.PC.Get() + 0x1)
	c.JP(address)
}

func (c *CPU) RET() {
	c.PopWord(&c.registers.PC)
}

func (c *CPU) RETc(flag int, flagValue types.Byte) {
	if c.registers.Flags.Get(flag) == flagValue {
		c.RET()
	}
}

func (c *CPU) RETI() {
	c.RET()
	c.EI()
}