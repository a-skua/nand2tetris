package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// OR returns a OR b.
//
//	a | b | out
//	0 | 0 | 0
//	0 | 1 | 1
//	1 | 0 | 1
//	1 | 1 | 1
func OR(a, b builtin.Bit) (out builtin.Bit) {
	out = NOT(AND(NOT(a), NOT(b)))
	return
}

// OR16 applies [OR] to each of the 16 bits.
func OR16(a, b [16]builtin.Bit) (out [16]builtin.Bit) {
	out[0] = OR(a[0], b[0])
	out[1] = OR(a[1], b[1])
	out[2] = OR(a[2], b[2])
	out[3] = OR(a[3], b[3])
	out[4] = OR(a[4], b[4])
	out[5] = OR(a[5], b[5])
	out[6] = OR(a[6], b[6])
	out[7] = OR(a[7], b[7])
	out[8] = OR(a[8], b[8])
	out[9] = OR(a[9], b[9])
	out[10] = OR(a[10], b[10])
	out[11] = OR(a[11], b[11])
	out[12] = OR(a[12], b[12])
	out[13] = OR(a[13], b[13])
	out[14] = OR(a[14], b[14])
	out[15] = OR(a[15], b[15])
	return
}

// OR8WAY returns 1 if any of the 8 bits is 1.
func OR8WAY(a [8]builtin.Bit) (out builtin.Bit) {
	out = OR(
		OR(
			OR(a[0], a[1]),
			OR(a[2], a[3]),
		),
		OR(
			OR(a[4], a[5]),
			OR(a[6], a[7]),
		),
	)
	return
}
