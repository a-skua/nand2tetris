package builtin

// DFF is a data flip-flop that holds one [Bit].
type DFF struct {
	bit Bit
}

// Clock returns the stored bit, then stores in if load is ON.
func (d *DFF) Clock(load, in Bit) (out Bit) {
	out = d.bit
	if load == ON {
		d.bit = in
	}
	return
}
