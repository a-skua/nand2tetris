package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestHalfADDER(t *testing.T) {
	tests := []struct {
		name               string
		a, b               builtin.Pin
		wantSum, wantCarry builtin.Pin
	}{
		{name: "HalfADDER(0,0)", a: 0, b: 0, wantSum: 0, wantCarry: 0},
		{name: "HalfADDER(0,1)", a: 0, b: 1, wantSum: 1, wantCarry: 0},
		{name: "HalfADDER(1,0)", a: 1, b: 0, wantSum: 1, wantCarry: 0},
		{name: "HalfADDER(1,1)", a: 1, b: 1, wantSum: 0, wantCarry: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, carry := HalfADDER(tt.a, tt.b)
			if sum != tt.wantSum || carry != tt.wantCarry {
				t.Errorf("HalfADDER(%d,%d) = (%d,%d); want (%d,%d)", tt.a, tt.b, sum, carry, tt.wantSum, tt.wantCarry)
			}
		})
	}
}

func TestFullADDER(t *testing.T) {
	tests := []struct {
		name               string
		a, b, c            builtin.Pin
		wantSum, wantCarry builtin.Pin
	}{
		{name: "FullADDER(0,0,0)", a: 0, b: 0, c: 0, wantSum: 0, wantCarry: 0},
		{name: "FullADDER(0,0,1)", a: 0, b: 0, c: 1, wantSum: 1, wantCarry: 0},
		{name: "FullADDER(0,1,0)", a: 0, b: 1, c: 0, wantSum: 1, wantCarry: 0},
		{name: "FullADDER(0,1,1)", a: 0, b: 1, c: 1, wantSum: 0, wantCarry: 1},
		{name: "FullADDER(1,0,0)", a: 1, b: 0, c: 0, wantSum: 1, wantCarry: 0},
		{name: "FullADDER(1,0,1)", a: 1, b: 0, c: 1, wantSum: 0, wantCarry: 1},
		{name: "FullADDER(1,1,0)", a: 1, b: 1, c: 0, wantSum: 0, wantCarry: 1},
		{name: "FullADDER(1,1,1)", a: 1, b: 1, c: 1, wantSum: 1, wantCarry: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, carry := FullADDER(tt.a, tt.b, tt.c)
			if sum != tt.wantSum || carry != tt.wantCarry {
				t.Errorf("FullADDER(%d,%d,%d) = (%d,%d); want (%d,%d)", tt.a, tt.b, tt.c, sum, carry, tt.wantSum, tt.wantCarry)
			}
		})
	}
}

func TestADD16(t *testing.T) {
	tests := []struct {
		name string
		a, b [16]builtin.Pin
		want [16]builtin.Pin
	}{
		{
			name: "ADD16(0, 0) = 0",
			a:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "ADD16(1, 1) = 2",
			a:    [16]builtin.Pin{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			b:    [16]builtin.Pin{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Pin{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "ADD16(0x1234, 0x9876) = 0xAAAA",
			a:    [16]builtin.Pin{0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0},
			b:    [16]builtin.Pin{0, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1},
			want: [16]builtin.Pin{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
		},
		{
			name: "ADD16(-1, 1) = 0 (overflow is ignored)",
			a:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Pin{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [16]builtin.Pin{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "ADD16(-1, -1) = -2",
			a:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			b:    [16]builtin.Pin{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: [16]builtin.Pin{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ADD16(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("ADD16(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
