package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// MUX
//
// a | b | sel | out
// 0 | 0 |  0  | 0
// 0 | 1 |  0  | 0
// 1 | 0 |  0  | 1
// 1 | 1 |  0  | 1
// 0 | 0 |  1  | 0
// 0 | 1 |  1  | 1
// 1 | 0 |  1  | 0
// 1 | 1 |  1  | 1
func MUX(a, b, sel builtin.Bit) (out builtin.Bit) {
	out = OR(AND(NOT(sel), a), AND(sel, b))
	return
}

// MUX16
//
// a | b | sel | out
// 0 | 0 |  0  | 0
// 0 | 1 |  0  | 0
// 1 | 0 |  0  | 1
// 1 | 1 |  0  | 1
// 0 | 0 |  1  | 0
// 0 | 1 |  1  | 1
// 1 | 0 |  1  | 0
// 1 | 1 |  1  | 1
func MUX16(a, b [16]builtin.Bit, sel builtin.Bit) (out [16]builtin.Bit) {
	out[0] = MUX(a[0], b[0], sel)
	out[1] = MUX(a[1], b[1], sel)
	out[2] = MUX(a[2], b[2], sel)
	out[3] = MUX(a[3], b[3], sel)
	out[4] = MUX(a[4], b[4], sel)
	out[5] = MUX(a[5], b[5], sel)
	out[6] = MUX(a[6], b[6], sel)
	out[7] = MUX(a[7], b[7], sel)
	out[8] = MUX(a[8], b[8], sel)
	out[9] = MUX(a[9], b[9], sel)
	out[10] = MUX(a[10], b[10], sel)
	out[11] = MUX(a[11], b[11], sel)
	out[12] = MUX(a[12], b[12], sel)
	out[13] = MUX(a[13], b[13], sel)
	out[14] = MUX(a[14], b[14], sel)
	out[15] = MUX(a[15], b[15], sel)
	return
}

// MUX4WAY16
//
// a | b | c | d | sel | out
// 1 | 0 | 0 | 0 | 00  | 1
// 0 | 1 | 0 | 0 | 01  | 1
// 0 | 0 | 1 | 0 | 10  | 1
// 0 | 0 | 0 | 1 | 11  | 1
func MUX4WAY16(a, b, c, d [16]builtin.Bit, sel [2]builtin.Bit) (out [16]builtin.Bit) {
	ab := MUX16(a, b, sel[0])
	cd := MUX16(c, d, sel[0])
	out = MUX16(ab, cd, sel[1])
	return
}

// MUX8WAY16
//
// a | b | c | d | e | f | g | h | sel | out
// 1 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 000 | 1
// 0 | 1 | 0 | 0 | 0 | 0 | 0 | 0 | 001 | 1
// 0 | 0 | 1 | 0 | 0 | 0 | 0 | 0 | 010 | 1
// 0 | 0 | 0 | 1 | 0 | 0 | 0 | 0 | 011 | 1
// 0 | 0 | 0 | 0 | 1 | 0 | 0 | 0 | 100 | 1
// 0 | 0 | 0 | 0 | 0 | 1 | 0 | 0 | 101 | 1
// 0 | 0 | 0 | 0 | 0 | 0 | 1 | 0 | 110 | 1
// 0 | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 111 | 1
func MUX8WAY16(a, b, c, d, e, f, g, h [16]builtin.Bit, sel [3]builtin.Bit) (out [16]builtin.Bit) {
	abcd := MUX4WAY16(a, b, c, d, [2]builtin.Bit(sel[0:2]))
	efgh := MUX4WAY16(e, f, g, h, [2]builtin.Bit(sel[0:2]))
	out = MUX16(abcd, efgh, sel[2])
	return
}
