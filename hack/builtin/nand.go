package builtin

// NAND returns NOT (a AND b).
//
//	a | b | out
//	0 | 0 | 1
//	0 | 1 | 1
//	1 | 0 | 1
//	1 | 1 | 0
func NAND(a, b Pin) (out Pin) {
	out = NewPin(^(a & b))
	return
}
