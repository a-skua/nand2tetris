package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// NOT
//
// in | out
// 0  | 1
// 1  | 0
func NOT(in builtin.Bit) (out builtin.Bit) {
	out = builtin.NAND(in, in)
	return
}

// NOT16
//
// in | out
// 0  | 1
// 1  | 0
func NOT16(in [16]builtin.Bit) (out [16]builtin.Bit) {
	out[0] = NOT(in[0])
	out[1] = NOT(in[1])
	out[2] = NOT(in[2])
	out[3] = NOT(in[3])
	out[4] = NOT(in[4])
	out[5] = NOT(in[5])
	out[6] = NOT(in[6])
	out[7] = NOT(in[7])
	out[8] = NOT(in[8])
	out[9] = NOT(in[9])
	out[10] = NOT(in[10])
	out[11] = NOT(in[11])
	out[12] = NOT(in[12])
	out[13] = NOT(in[13])
	out[14] = NOT(in[14])
	out[15] = NOT(in[15])
	return
}
