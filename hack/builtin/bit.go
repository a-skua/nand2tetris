// Package builtin provides the primitive chips: [NAND] and [DFF].
package builtin

// Pin is a single bit, 0 or 1.
type Pin int8

// Pin values.
const (
	ON  = 1
	OFF = 0
)

// NewPin returns the lowest bit of value.
func NewPin[V ~int8](value V) Pin {
	return Pin(value & 1)
}

// Word is a 16-bit value. Index 0 is the least significant bit.
type Word [16]Pin

// NewWord returns value as a Word in two's complement.
func NewWord[V ~int16](value V) (w Word) {
	for i := range w {
		w[i] = NewPin(int8(value >> i))
	}
	return
}
