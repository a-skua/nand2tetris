package chip

import "github.com/a-skua/nand2tetris/hack/builtin"

// DMUX routes in to a or b according to sel.
//
//	sel | a  | b
//	0   | in | 0
//	1   | 0  | in
func DMUX(in, sel builtin.Pin) (a, b builtin.Pin) {
	a = AND(NOT(sel), in)
	b = AND(sel, in)
	return
}

// DMUX4WAY routes in to one of a-d according to sel.
// sel is written MSB first (sel[0] is the LSB).
//
//	sel | a  | b  | c  | d
//	00  | in | 0  | 0  | 0
//	01  | 0  | in | 0  | 0
//	10  | 0  | 0  | in | 0
//	11  | 0  | 0  | 0  | in
func DMUX4WAY(in builtin.Pin, sel [2]builtin.Pin) (a, b, c, d builtin.Pin) {
	ab, cd := DMUX(in, sel[1])
	a, b = DMUX(ab, sel[0])
	c, d = DMUX(cd, sel[0])
	return
}

// DMUX8WAY routes in to one of a-h according to sel.
// sel is written MSB first (sel[0] is the LSB).
//
//	sel | a  | b  | c  | d  | e  | f  | g  | h
//	000 | in | 0  | 0  | 0  | 0  | 0  | 0  | 0
//	001 | 0  | in | 0  | 0  | 0  | 0  | 0  | 0
//	010 | 0  | 0  | in | 0  | 0  | 0  | 0  | 0
//	011 | 0  | 0  | 0  | in | 0  | 0  | 0  | 0
//	100 | 0  | 0  | 0  | 0  | in | 0  | 0  | 0
//	101 | 0  | 0  | 0  | 0  | 0  | in | 0  | 0
//	110 | 0  | 0  | 0  | 0  | 0  | 0  | in | 0
//	111 | 0  | 0  | 0  | 0  | 0  | 0  | 0  | in
func DMUX8WAY(in builtin.Pin, sel [3]builtin.Pin) (a, b, c, d, e, f, g, h builtin.Pin) {
	abcd, efgh := DMUX(in, sel[2])
	a, b, c, d = DMUX4WAY(abcd, [2]builtin.Pin(sel[0:2]))
	e, f, g, h = DMUX4WAY(efgh, [2]builtin.Pin(sel[0:2]))
	return
}
