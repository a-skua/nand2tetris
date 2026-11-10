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
