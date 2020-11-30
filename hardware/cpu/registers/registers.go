package registers

import (
	"log"
	"github.com/cgimenes/gomenes-boy/hardware/types"
)

type WordRegister interface {
	Get() types.Word
	Set(w types.Word)
}

type SingleWordRegister struct {
	r types.Word
}

func (r SingleWordRegister) Get() types.Word {
	return r.r
}

func (r *SingleWordRegister) Set(w types.Word) {
	r.r = w
}

type PairedWordRegister struct {
	H *types.ByteRegister
	L *types.ByteRegister
}

func (r PairedWordRegister) Get() types.Word {
	return types.WordFromBytes(r.H.Get(), r.L.Get())
}

func (r *PairedWordRegister) Set(w types.Word) {
	h, l := types.WordToBytes(w)
	r.H.Set(h)
	r.L.Set(l)
}

type Flags struct {
	R *types.ByteRegister
}

//flags
const (
	_ = iota
	C
	H
	N
	Z
)

func (f *Flags) Reset(flag int) {
	switch flag {
	case Z:
		f.R.Set(types.ResetBit(7, f.R.Get()))
	case N:
		f.R.Set(types.ResetBit(6, f.R.Get()))
	case H:
		f.R.Set(types.ResetBit(5, f.R.Get()))
	case C:
		f.R.Set(types.ResetBit(4, f.R.Get()))
	default:
		log.Fatalf("Unknown flag %c", flag)
	}
	f.R.Set(0xF0 & f.R.Get())
}

func (f *Flags) Set(flag int) {
	switch flag {
	case Z:
		f.R.Set(types.SetBit(7, f.R.Get()))
	case N:
		f.R.Set(types.SetBit(6, f.R.Get()))
	case H:
		f.R.Set(types.SetBit(5, f.R.Get()))
	case C:
		f.R.Set(types.SetBit(4, f.R.Get()))
	default:
		log.Fatalf("Unknown flag %c", flag)
	}
	f.R.Set(0xF0 & f.R.Get())
}

func (f Flags) Get(flag int) types.Byte {
	switch flag {
	case Z:
		return types.GetBit(7, f.R.Get())
	case N:
		return types.GetBit(6, f.R.Get())
	case H:
		return types.GetBit(5, f.R.Get())
	case C:
		return types.GetBit(4, f.R.Get())
	default:
		log.Fatalf("Unknown flag %c", flag)
		return 0
	}
}

type Registers struct {
	// 8-bit high
	A types.ByteRegister // accumulator
	B types.ByteRegister
	D types.ByteRegister
	H types.ByteRegister

	// 8-bit low
	F types.ByteRegister // flags
	C types.ByteRegister
	E types.ByteRegister
	L types.ByteRegister

	// 16-bit
	SP SingleWordRegister // stack pointer
	PC SingleWordRegister // program counter

	// 16-bit paired
	AF PairedWordRegister
	BC PairedWordRegister
	DE PairedWordRegister
	HL PairedWordRegister

	// flags
	Flags Flags
}