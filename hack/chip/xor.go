package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// XOR returns a XOR b.
//
//	a | b | out
//	0 | 0 | 0
//	0 | 1 | 1
//	1 | 0 | 1
//	1 | 1 | 0
func XOR(a, b builtin.Pin) (out builtin.Pin) {
	out = AND(OR(a, b), NOT(AND(a, b)))
	return
}

// XOR16 applies [XOR] to each of the 16 bits.
func XOR16(a, b [16]builtin.Pin) (out [16]builtin.Pin) {
	out[0] = XOR(a[0], b[0])
	out[1] = XOR(a[1], b[1])
	out[2] = XOR(a[2], b[2])
	out[3] = XOR(a[3], b[3])
	out[4] = XOR(a[4], b[4])
	out[5] = XOR(a[5], b[5])
	out[6] = XOR(a[6], b[6])
	out[7] = XOR(a[7], b[7])
	out[8] = XOR(a[8], b[8])
	out[9] = XOR(a[9], b[9])
	out[10] = XOR(a[10], b[10])
	out[11] = XOR(a[11], b[11])
	out[12] = XOR(a[12], b[12])
	out[13] = XOR(a[13], b[13])
	out[14] = XOR(a[14], b[14])
	out[15] = XOR(a[15], b[15])
	return
}
