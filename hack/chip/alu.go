package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// ALU computes out from x and y according to the control bits.
// zr is 1 if out == 0, and ng is 1 if out < 0.
//
//	zx | nx | zy | ny | f | no | out
//	1  | 0  | 1  | 0  | 1 | 0  | 0
//	1  | 1  | 1  | 1  | 1 | 1  | 1
//	1  | 1  | 1  | 0  | 1 | 0  | -1
//	0  | 0  | 1  | 1  | 0 | 0  | x
//	1  | 1  | 0  | 0  | 0 | 0  | y
//	0  | 0  | 1  | 1  | 0 | 1  | !x
//	1  | 1  | 0  | 0  | 0 | 1  | !y
//	0  | 0  | 1  | 1  | 1 | 1  | -x
//	1  | 1  | 0  | 0  | 1 | 1  | -y
//	0  | 1  | 1  | 1  | 1 | 1  | x+1
//	1  | 1  | 0  | 1  | 1 | 1  | y+1
//	0  | 0  | 1  | 1  | 1 | 0  | x-1
//	1  | 1  | 0  | 0  | 1 | 0  | y-1
//	0  | 0  | 0  | 0  | 1 | 0  | x+y
//	0  | 1  | 0  | 0  | 1 | 1  | x-y
//	0  | 0  | 0  | 1  | 1 | 1  | y-x
//	0  | 0  | 0  | 0  | 0 | 0  | x&y
//	0  | 1  | 0  | 1  | 0 | 1  | x|y
func ALU(x, y [16]builtin.Pin, zx, nx, zy, ny, f, no builtin.Pin) (out [16]builtin.Pin, zr, ng builtin.Pin) {
	var zero [16]builtin.Pin

	// if zx then x = 0
	// if nx then x = !x
	x = MUX16(x, zero, zx)
	x = MUX16(x, NOT16(x), nx)

	// if zy then y = 0
	// if ny then y = !y
	y = MUX16(y, zero, zy)
	y = MUX16(y, NOT16(y), ny)

	// if f then out = x + y else out = x & y
	out = MUX16(AND16(x, y), ADD16(x, y), f)

	// if no then out = !out
	out = MUX16(out, NOT16(out), no)

	zr = NOT(OR(
		OR8WAY([8]builtin.Pin(out[0:8])),
		OR8WAY([8]builtin.Pin(out[8:16])),
	))
	ng = out[15]
	return
}
