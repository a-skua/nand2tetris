package chip

import (
	"github.com/a-skua/nand2tetris/hack/chip/builtin"
	"testing"
)

func toBits16(v int16) (out [16]builtin.Pin) {
	for i := range out {
		out[i] = builtin.Pin((uint16(v) >> i) & 1)
	}
	return
}

func TestALU(t *testing.T) {
	type control struct {
		zx, nx, zy, ny, f, no builtin.Pin
	}
	funcs := []struct {
		name string
		ctrl control
		want func(x, y int16) int16
	}{
		{name: "0", ctrl: control{1, 0, 1, 0, 1, 0}, want: func(x, y int16) int16 { return 0 }},
		{name: "1", ctrl: control{1, 1, 1, 1, 1, 1}, want: func(x, y int16) int16 { return 1 }},
		{name: "-1", ctrl: control{1, 1, 1, 0, 1, 0}, want: func(x, y int16) int16 { return -1 }},
		{name: "x", ctrl: control{0, 0, 1, 1, 0, 0}, want: func(x, y int16) int16 { return x }},
		{name: "y", ctrl: control{1, 1, 0, 0, 0, 0}, want: func(x, y int16) int16 { return y }},
		{name: "!x", ctrl: control{0, 0, 1, 1, 0, 1}, want: func(x, y int16) int16 { return ^x }},
		{name: "!y", ctrl: control{1, 1, 0, 0, 0, 1}, want: func(x, y int16) int16 { return ^y }},
		{name: "-x", ctrl: control{0, 0, 1, 1, 1, 1}, want: func(x, y int16) int16 { return -x }},
		{name: "-y", ctrl: control{1, 1, 0, 0, 1, 1}, want: func(x, y int16) int16 { return -y }},
		{name: "x+1", ctrl: control{0, 1, 1, 1, 1, 1}, want: func(x, y int16) int16 { return x + 1 }},
		{name: "y+1", ctrl: control{1, 1, 0, 1, 1, 1}, want: func(x, y int16) int16 { return y + 1 }},
		{name: "x-1", ctrl: control{0, 0, 1, 1, 1, 0}, want: func(x, y int16) int16 { return x - 1 }},
		{name: "y-1", ctrl: control{1, 1, 0, 0, 1, 0}, want: func(x, y int16) int16 { return y - 1 }},
		{name: "x+y", ctrl: control{0, 0, 0, 0, 1, 0}, want: func(x, y int16) int16 { return x + y }},
		{name: "x-y", ctrl: control{0, 1, 0, 0, 1, 1}, want: func(x, y int16) int16 { return x - y }},
		{name: "y-x", ctrl: control{0, 0, 0, 1, 1, 1}, want: func(x, y int16) int16 { return y - x }},
		{name: "x&y", ctrl: control{0, 0, 0, 0, 0, 0}, want: func(x, y int16) int16 { return x & y }},
		{name: "x|y", ctrl: control{0, 1, 0, 1, 0, 1}, want: func(x, y int16) int16 { return x | y }},
	}
	inputs := []struct{ x, y int16 }{
		{0, 0},
		{0, -1},
		{17, 3},
		{3, 17},
		{-12345, 6789},
		{32767, 1},
		{-32768, -1},
	}

	for _, fn := range funcs {
		for _, in := range inputs {
			t.Run(fn.name, func(t *testing.T) {
				c := fn.ctrl
				out, zr, ng := ALU(toBits16(in.x), toBits16(in.y), c.zx, c.nx, c.zy, c.ny, c.f, c.no)

				v := fn.want(in.x, in.y)
				var wantZr, wantNg builtin.Pin
				if v == 0 {
					wantZr = 1
				}
				if v < 0 {
					wantNg = 1
				}
				if out != toBits16(v) || zr != wantZr || ng != wantNg {
					t.Errorf("ALU(x=%d, y=%d) = (%v, zr=%d, ng=%d); want (%v (%d), zr=%d, ng=%d)",
						in.x, in.y, out, zr, ng, toBits16(v), v, wantZr, wantNg)
				}
			})
		}
	}
}
