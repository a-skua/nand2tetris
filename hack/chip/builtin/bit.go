// Package builtin provides the primitive chips: [NAND] and [DFF].
package builtin

// Bit is a single bit, 0 or 1.
type Bit int8

// Bit values.
const (
	ON  = 1
	OFF = 0
)

// NewBit returns the lowest bit of value.
func NewBit[V ~int8](value V) Bit {
	return Bit(value & 1)
}
