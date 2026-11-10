package builtin

// NAND
//
// a | b | out
// 0 | 0 | 1
// 0 | 1 | 1
// 1 | 0 | 1
// 1 | 1 | 0
func NAND(a, b Bit) (out Bit) {
	out = NewBit(^(a & b))
	return
}
