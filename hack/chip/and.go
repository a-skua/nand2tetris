package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// AND
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 0
// 1 | 0 | 0
// 1 | 1 | 1
func AND(a, b builtin.Bit) (out builtin.Bit) {
	out = NOT(builtin.NAND(a, b))
	return
}

// AND16
//
// a | b | out
// 0 | 0 | 0
// 0 | 1 | 0
// 1 | 0 | 0
// 1 | 1 | 1
func AND16(a, b [16]builtin.Bit) (out [16]builtin.Bit) {
	out[0] = AND(a[0], b[0])
	out[1] = AND(a[1], b[1])
	out[2] = AND(a[2], b[2])
	out[3] = AND(a[3], b[3])
	out[4] = AND(a[4], b[4])
	out[5] = AND(a[5], b[5])
	out[6] = AND(a[6], b[6])
	out[7] = AND(a[7], b[7])
	out[8] = AND(a[8], b[8])
	out[9] = AND(a[9], b[9])
	out[10] = AND(a[10], b[10])
	out[11] = AND(a[11], b[11])
	out[12] = AND(a[12], b[12])
	out[13] = AND(a[13], b[13])
	out[14] = AND(a[14], b[14])
	out[15] = AND(a[15], b[15])
	return
}
