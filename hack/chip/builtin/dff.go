package builtin

// DFF is a data flip-flop that holds one [Bit].
type DFF struct {
	bit Bit
}

// Clock returns the stored bit, then stores in.
// That is, out(t) = in(t-1).
func (d *DFF) Clock(in Bit) (out Bit) {
	out = d.bit
	d.bit = in
	return
}
