package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestMUX(t *testing.T) {
	tests := []struct {
		name      string
		a, b, sel builtin.Pin
		want      builtin.Pin
	}{
		{name: "MUX(0,0,0)", a: 0, b: 0, sel: 0, want: 0},
		{name: "MUX(0,1,0)", a: 0, b: 1, sel: 0, want: 0},
		{name: "MUX(1,0,0)", a: 1, b: 0, sel: 0, want: 1},
		{name: "MUX(1,1,0)", a: 1, b: 1, sel: 0, want: 1},
		{name: "MUX(0,0,1)", a: 0, b: 0, sel: 1, want: 0},
		{name: "MUX(0,1,1)", a: 0, b: 1, sel: 1, want: 1},
		{name: "MUX(1,0,1)", a: 1, b: 0, sel: 1, want: 0},
		{name: "MUX(1,1,1)", a: 1, b: 1, sel: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MUX(tt.a, tt.b, tt.sel)
			if got != tt.want {
				t.Errorf("MUX(%d,%d,%d) = %d; want %d", tt.a, tt.b, tt.sel, got, tt.want)
			}
		})
	}
}

func TestMUX16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Pin
		sel  builtin.Pin
		want [16]builtin.Pin
	}{
		{
			name: "MUX16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 0)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			sel:  0,
			want: [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "MUX16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 1)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			sel:  1,
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MUX16(tt.a, tt.b, tt.sel)
			if got != tt.want {
				t.Errorf("MUX16(%v,%v,%d) = %v; want %v", tt.a, tt.b, tt.sel, got, tt.want)
			}
		})
	}
}

func TestMUX4WAY16(t *testing.T) {
	tests := []struct {
		name       string
		a, b, c, d [16]builtin.Pin
		sel        [2]builtin.Pin
		want       [16]builtin.Pin
	}{
		{
			name: "MUX4WAY16(1, 0, 0, 0, 00)",
			a:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			c:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			d:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			sel:  [2]builtin.Pin{0, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "MUX4WAY16(0, 1, 0, 0, 01)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			c:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			d:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			sel:  [2]builtin.Pin{1, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "MUX4WAY16(0, 0, 1, 0, 10)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			c:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			d:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			sel:  [2]builtin.Pin{0, 1},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "MUX4WAY16(0, 0, 0, 1, 11)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			c:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			d:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			sel:  [2]builtin.Pin{1, 1},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MUX4WAY16(tt.a, tt.b, tt.c, tt.d, tt.sel)
			if got != tt.want {
				t.Errorf("MUX4WAY16(%v,%v,%v,%v,%v) = %v; want %v", tt.a, tt.b, tt.c, tt.d, tt.sel, got, tt.want)
			}
		})
	}
}

func TestMUX8WAY16(t *testing.T) {
	tests := []struct {
		name                   string
		a, b, c, d, e, f, g, h [16]builtin.Pin
		sel                    [3]builtin.Pin
		want                   [16]builtin.Pin
	}{
		{
			name: "MUX8WAY16(1, 0, 0, 0, 0, 0, 0, 0, 000)",
			a:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			c:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			d:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			e:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			f:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			g:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			h:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			sel:  [3]builtin.Pin{0, 0, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "MUX8WAY16(0, 1, 0, 0, 0, 0, 0, 0, 001)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			c:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			d:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			e:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			f:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			g:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			h:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			sel:  [3]builtin.Pin{1, 0, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "MUX8WAY16(0, 0, 0, 0, 0, 0, 0, 1, 111)",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			c:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			d:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			e:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			f:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			g:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			h:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			sel:  [3]builtin.Pin{1, 1, 1},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MUX8WAY16(tt.a, tt.b, tt.c, tt.d, tt.e, tt.f, tt.g, tt.h, tt.sel)
			if got != tt.want {
				t.Errorf("MUX8WAY16(%v,%v,%v,%v,%v,%v,%v,%v,%v) = %v; want %v", tt.a, tt.b, tt.c, tt.d, tt.e, tt.f, tt.g, tt.h, tt.sel, got, tt.want)
			}
		})
	}
}
