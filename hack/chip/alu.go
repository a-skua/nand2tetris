package chip

import "github.com/a-skua/nand2tetris/hack/chip/builtin"

// ALU
//
// zx | nx | zy | ny | f | no | out
// 1  | 0  | 1  | 0  | 1 | 0  | 0
// 1  | 1  | 1  | 1  | 1 | 1  | 1
// 1  | 1  | 1  | 0  | 1 | 0  | -1
// 0  | 0  | 1  | 1  | 0 | 0  | x
// 1  | 1  | 0  | 0  | 0 | 0  | y
// 0  | 0  | 1  | 1  | 0 | 1  | !x
// 1  | 1  | 0  | 0  | 0 | 1  | !y
// 0  | 0  | 1  | 1  | 1 | 1  | -x
// 1  | 1  | 0  | 0  | 1 | 1  | -y
// 0  | 1  | 1  | 1  | 1 | 1  | x+1
// 1  | 1  | 0  | 1  | 1 | 1  | y+1
// 0  | 0  | 1  | 1  | 1 | 0  | x-1
// 1  | 1  | 0  | 0  | 1 | 0  | y-1
// 0  | 0  | 0  | 0  | 1 | 0  | x+y
// 0  | 1  | 0  | 0  | 1 | 1  | x-y
// 0  | 0  | 0  | 1  | 1 | 1  | y-x
// 0  | 0  | 0  | 0  | 0 | 0  | x&y
// 0  | 1  | 0  | 1  | 0 | 1  | x|y
//
// zr = 1 if out == 0
// ng = 1 if out < 0
func ALU(x, y [16]builtin.Bit, zx, nx, zy, ny, f, no builtin.Bit) (out [16]builtin.Bit, zr, ng builtin.Bit) {
	var zero [16]builtin.Bit

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
		OR8WAY([8]builtin.Bit(out[0:8])),
		OR8WAY([8]builtin.Bit(out[8:16])),
	))
	ng = out[15]
	return
}
