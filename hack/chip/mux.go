package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// MUX returns a if sel is 0, otherwise b.
//
//	sel | out
//	0   | a
//	1   | b
func MUX(a, b, sel builtin.Pin) (out builtin.Pin) {
	out = OR(AND(NOT(sel), a), AND(sel, b))
	return
}

// MUX16 is a 16-bit [MUX].
func MUX16(a, b [16]builtin.Pin, sel builtin.Pin) (out [16]builtin.Pin) {
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

// MUX4WAY16 selects one of a-d according to sel.
// sel is written MSB first (sel[0] is the LSB).
//
//	sel | out
//	00  | a
//	01  | b
//	10  | c
//	11  | d
func MUX4WAY16(a, b, c, d [16]builtin.Pin, sel [2]builtin.Pin) (out [16]builtin.Pin) {
	ab := MUX16(a, b, sel[0])
	cd := MUX16(c, d, sel[0])
	out = MUX16(ab, cd, sel[1])
	return
}

// MUX8WAY16 selects one of a-h according to sel.
// sel is written MSB first (sel[0] is the LSB).
//
//	sel | out
//	000 | a
//	001 | b
//	010 | c
//	011 | d
//	100 | e
//	101 | f
//	110 | g
//	111 | h
func MUX8WAY16(a, b, c, d, e, f, g, h [16]builtin.Pin, sel [3]builtin.Pin) (out [16]builtin.Pin) {
	abcd := MUX4WAY16(a, b, c, d, [2]builtin.Pin(sel[0:2]))
	efgh := MUX4WAY16(e, f, g, h, [2]builtin.Pin(sel[0:2]))
	out = MUX16(abcd, efgh, sel[2])
	return
}
