package builtin

// Bit - 0 or 1
type Bit int8

func NewBit[V ~int8](value V) Bit {
	return Bit(value & 1)
}
