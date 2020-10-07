package types

type Word = uint16
type Byte = uint8

type ByteRegister struct {
	r Byte
}

func (r ByteRegister) Get() Byte {
	return r.r
}

func (r *ByteRegister) Set(b Byte) {
	r.r = b
}

func WordFromBytes(h, l Byte) Word {
	return (Word(h) << 8) ^ Word(l)
}

func WordToBytes(w Word) (Byte, Byte) {
	return Byte(w >> 8), Byte(w & 0x00FF)
}

func ResetBit(bit byte, a Byte) Byte {
	return a & ^(1 << uint(bit))
}

func SetBit(bit byte, a Byte) Byte {
	return a | (1 << uint(bit))
}

func GetBit(bit byte, a Byte) Byte {
	return (a & (1 << uint(bit))) >> uint(bit)
}