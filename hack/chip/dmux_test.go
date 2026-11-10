package chip

import (
	"github.com/a-skua/nand2tetris/hack/builtin"
	"testing"
)

func TestDMUX(t *testing.T) {
	tests := []struct {
		name         string
		in, sel      builtin.Pin
		wantA, wantB builtin.Pin
	}{
		{name: "DMUX(0,0)", in: 0, sel: 0, wantA: 0, wantB: 0},
		{name: "DMUX(0,1)", in: 0, sel: 1, wantA: 0, wantB: 0},
		{name: "DMUX(1,0)", in: 1, sel: 0, wantA: 1, wantB: 0},
		{name: "DMUX(1,1)", in: 1, sel: 1, wantA: 0, wantB: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := DMUX(tt.in, tt.sel)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("DMUX(%d,%d) = (%d,%d); want (%d,%d)", tt.in, tt.sel, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestDMUX4WAY(t *testing.T) {
	tests := []struct {
		name                       string
		in                         builtin.Pin
		sel                        [2]builtin.Pin
		wantA, wantB, wantC, wantD builtin.Pin
	}{
		{name: "DMUX4WAY(0,[0,0])", in: 0, sel: [2]builtin.Pin{0, 0}, wantA: 0, wantB: 0, wantC: 0, wantD: 0},
		{name: "DMUX4WAY(0,[0,1])", in: 0, sel: [2]builtin.Pin{1, 0}, wantA: 0, wantB: 0, wantC: 0, wantD: 0},
		{name: "DMUX4WAY(0,[1,0])", in: 0, sel: [2]builtin.Pin{0, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 0},
		{name: "DMUX4WAY(0,[1,1])", in: 0, sel: [2]builtin.Pin{1, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 0},
		{name: "DMUX4WAY(1,[0,0])", in: 1, sel: [2]builtin.Pin{0, 0}, wantA: 1, wantB: 0, wantC: 0, wantD: 0},
		{name: "DMUX4WAY(1,[0,1])", in: 1, sel: [2]builtin.Pin{1, 0}, wantA: 0, wantB: 1, wantC: 0, wantD: 0},
		{name: "DMUX4WAY(1,[1,0])", in: 1, sel: [2]builtin.Pin{0, 1}, wantA: 0, wantB: 0, wantC: 1, wantD: 0},
		{name: "DMUX4WAY(1,[1,1])", in: 1, sel: [2]builtin.Pin{1, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB, gotC, gotD := DMUX4WAY(tt.in, tt.sel)
			if gotA != tt.wantA || gotB != tt.wantB || gotC != tt.wantC || gotD != tt.wantD {
				t.Errorf("DMUX4WAY(%d,%v) = (%d,%d,%d,%d); want (%d,%d,%d,%d)", tt.in, tt.sel, gotA, gotB, gotC, gotD, tt.wantA, tt.wantB, tt.wantC, tt.wantD)
			}
		})
	}
}

func TestDMUX8WAY(t *testing.T) {
	tests := []struct {
		name                                                   string
		in                                                     builtin.Pin
		sel                                                    [3]builtin.Pin
		wantA, wantB, wantC, wantD, wantE, wantF, wantG, wantH builtin.Pin
	}{
		{name: "DMUX8WAY(0,[0,0,0])", in: 0, sel: [3]builtin.Pin{0, 0, 0}, wantA: 0, wantB: 0, wantC: 0, wantD: 0, wantE: 0, wantF: 0, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[0,0,0])", in: 1, sel: [3]builtin.Pin{0, 0, 0}, wantA: 1, wantB: 0, wantC: 0, wantD: 0, wantE: 0, wantF: 0, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[0,0,1])", in: 1, sel: [3]builtin.Pin{1, 0, 0}, wantA: 0, wantB: 1, wantC: 0, wantD: 0, wantE: 0, wantF: 0, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[0,1,0])", in: 1, sel: [3]builtin.Pin{0, 1, 0}, wantA: 0, wantB: 0, wantC: 1, wantD: 0, wantE: 0, wantF: 0, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[0,1,1])", in: 1, sel: [3]builtin.Pin{1, 1, 0}, wantA: 0, wantB: 0, wantC: 0, wantD: 1, wantE: 0, wantF: 0, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[1,0,0])", in: 1, sel: [3]builtin.Pin{0, 0, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 0, wantE: 1, wantF: 0, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[1,0,1])", in: 1, sel: [3]builtin.Pin{1, 0, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 0, wantE: 0, wantF: 1, wantG: 0, wantH: 0},
		{name: "DMUX8WAY(1,[1,1,0])", in: 1, sel: [3]builtin.Pin{0, 1, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 0, wantE: 0, wantF: 0, wantG: 1, wantH: 0},
		{name: "DMUX8WAY(1,[1,1,1])", in: 1, sel: [3]builtin.Pin{1, 1, 1}, wantA: 0, wantB: 0, wantC: 0, wantD: 0, wantE: 0, wantF: 0, wantG: 0, wantH: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB, gotC, gotD, gotE, gotF, gotG, gotH := DMUX8WAY(tt.in, tt.sel)
			if gotA != tt.wantA || gotB != tt.wantB || gotC != tt.wantC || gotD != tt.wantD || gotE != tt.wantE || gotF != tt.wantF || gotG != tt.wantG || gotH != tt.wantH {
				t.Errorf("DMUX8WAY(%d,%v) = (%d,%d,%d,%d,%d,%d,%d,%d); want (%d,%d,%d,%d,%d,%d,%d,%d)", tt.in, tt.sel, gotA, gotB, gotC, gotD, gotE, gotF, gotG, gotH, tt.wantA, tt.wantB, tt.wantC, tt.wantD, tt.wantE, tt.wantF, tt.wantG, tt.wantH)
			}
		})
	}
}
