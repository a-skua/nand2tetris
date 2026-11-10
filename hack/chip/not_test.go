package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestNOT(t *testing.T) {
	tests := []struct {
		name string
		in   builtin.Pin
		want builtin.Pin
	}{
		{name: "NOT(0)", in: 0, want: 1},
		{name: "NOT(1)", in: 1, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NOT(tt.in)
			if got != tt.want {
				t.Errorf("NOT(%d) = %d; want %d", tt.in, got, tt.want)
			}
		})
	}
}

func TestNOT16(t *testing.T) {
	tests := []struct {
		name string
		in   [16]builtin.Pin
		want [16]builtin.Pin
	}{
		{
			name: "NOT16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0])",
			in:   [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "NOT16([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
			in:   [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "NOT16([0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1])",
			in:   [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			want: [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NOT16(tt.in)
			if got != tt.want {
				t.Errorf("NOT16(%v) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}
}
