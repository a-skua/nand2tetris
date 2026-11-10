package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestAND(t *testing.T) {
	tests := []struct {
		name string
		a, b builtin.Pin
		want builtin.Pin
	}{
		{name: "AND(0,0)", a: 0, b: 0, want: 0},
		{name: "AND(0,1)", a: 0, b: 1, want: 0},
		{name: "AND(1,0)", a: 1, b: 0, want: 0},
		{name: "AND(1,1)", a: 1, b: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AND(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("AND(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestAND16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Pin
		want [16]builtin.Pin
	}{
		{
			name: "AND16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "AND16([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], [1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0])",
			a:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			want: [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AND16(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("AND16(%v,%v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
