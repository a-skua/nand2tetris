package builtin

// DFF is a data flip-flop that holds one [Pin].
type DFF struct {
	Pin
}

// Clock returns the stored bit, then stores in.
// That is, out(t) = in(t-1).
func (d *DFF) Clock(in Pin) (out Pin) {
	out = d.Pin
	d.Pin = in
	return
}
