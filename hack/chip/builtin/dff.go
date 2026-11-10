package builtin

type DFF struct {
	bit Bit
}

func (d *DFF) Clock(load, in Bit) (out Bit) {
	out = d.bit
	if load == ON {
		d.bit = in
	}
	return
}
