package ppu

import (
	"neutronstar-gb/hardware/types"
)

const defaultBGP types.Byte = 0x1B

type Tile struct {
	Tile [16]types.Byte
}

type Tiles struct {
	Tiles [256]Tile
}

type Registers struct {
	LCDC types.ByteRegister
	STAT types.ByteRegister
	SCY types.ByteRegister
	SCX types.ByteRegister
	LY types.ByteRegister
	LYC types.ByteRegister
	DMA types.ByteRegister
	BGP types.ByteRegister
	OBP0 types.ByteRegister
	OBP1 types.ByteRegister
	WY types.ByteRegister
	WX types.ByteRegister
}