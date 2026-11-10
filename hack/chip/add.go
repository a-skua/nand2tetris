package chip

import "github.com/a-skua/nand2tetris/hack/builtin"

// HalfADDER returns the sum and carry of a + b.
//
//	a | b | sum | carry
//	0 | 0 |  0  |  0
//	0 | 1 |  1  |  0
//	1 | 0 |  1  |  0
//	1 | 1 |  0  |  1
func HalfADDER(a, b builtin.Pin) (sum, carry builtin.Pin) {
	sum = XOR(a, b)
	carry = AND(a, b)
	return
}

// FullADDER returns the sum and carry of a + b + c.
//
//	a | b | c | sum | carry
//	0 | 0 | 0 |  0  |  0
//	0 | 0 | 1 |  1  |  0
//	0 | 1 | 0 |  1  |  0
//	0 | 1 | 1 |  0  |  1
//	1 | 0 | 0 |  1  |  0
//	1 | 0 | 1 |  0  |  1
//	1 | 1 | 0 |  0  |  1
//	1 | 1 | 1 |  1  |  1
func FullADDER(a, b, c builtin.Pin) (sum, carry builtin.Pin) {
	ab, abCarry := HalfADDER(a, b)
	sum, abcCarry := HalfADDER(ab, c)
	carry = OR(abCarry, abcCarry)
	return
}

// ADD16 returns a + b as 16-bit two's complement.
// Index 0 is the LSB. Overflow is ignored.
func ADD16(a, b [16]builtin.Pin) (out [16]builtin.Pin) {
	var c builtin.Pin
	out[0], c = HalfADDER(a[0], b[0])
	out[1], c = FullADDER(a[1], b[1], c)
	out[2], c = FullADDER(a[2], b[2], c)
	out[3], c = FullADDER(a[3], b[3], c)
	out[4], c = FullADDER(a[4], b[4], c)
	out[5], c = FullADDER(a[5], b[5], c)
	out[6], c = FullADDER(a[6], b[6], c)
	out[7], c = FullADDER(a[7], b[7], c)
	out[8], c = FullADDER(a[8], b[8], c)
	out[9], c = FullADDER(a[9], b[9], c)
	out[10], c = FullADDER(a[10], b[10], c)
	out[11], c = FullADDER(a[11], b[11], c)
	out[12], c = FullADDER(a[12], b[12], c)
	out[13], c = FullADDER(a[13], b[13], c)
	out[14], c = FullADDER(a[14], b[14], c)
	out[15], _ = FullADDER(a[15], b[15], c)
	return
}
