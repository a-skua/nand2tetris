package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestOR(t *testing.T) {
	tests := []struct {
		name string
		a, b builtin.Pin
		want builtin.Pin
	}{
		{name: "OR(0,0)", a: 0, b: 0, want: 0},
		{name: "OR(0,1)", a: 0, b: 1, want: 1},
		{name: "OR(1,0)", a: 1, b: 0, want: 1},
		{name: "OR(1,1)", a: 1, b: 1, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OR(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("OR(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestOR16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Pin
		want [16]builtin.Pin
	}{
		{
			name: "OR16([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0], [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1])",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "OR16([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], [0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0])",
			a:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "OR16([0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1], [1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0])",
			a:    [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			b:    [16]builtin.Pin{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0},
			want: [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OR16(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("OR16(%v,%v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestOR8WAY(t *testing.T) {
	tests := []struct {
		name string
		a    [8]builtin.Pin
		want builtin.Pin
	}{
		{
			name: "OR8WAY([0,0,0,0,0,0,0,0])",
			a:    [8]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0},
			want: 0,
		},
		{
			name: "OR8WAY([1,1,1,1,1,1,1,1])",
			a:    [8]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1},
			want: 1,
		},
		{
			name: "OR8WAY([0,0,0,0,0,0,0,1])",
			a:    [8]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 1},
			want: 1,
		},
		{
			name: "OR8WAY([1,0,0,0,0,0,0,0])",
			a:    [8]builtin.Pin{1, 0, 0, 0, 0, 0, 0, 0},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OR8WAY(tt.a)
			if got != tt.want {
				t.Errorf("OR8WAY(%v) = %v; want %v", tt.a, got, tt.want)
			}
		})
	}
}
