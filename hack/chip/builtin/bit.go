package builtin

// Bit - 0 or 1
type Bit int8

const (
	ON  = 1
	OFF = 0
)

func NewBit[V ~int8](value V) Bit {
	return Bit(value & 1)
}
