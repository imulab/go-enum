package enum

type bitmap uint

func (f bitmap) Has(flag bitmap) bool {
	return f&flag != 0
}

func (f *bitmap) Set(flag bitmap) {
	*f |= flag
}

func (f *bitmap) Clear(flag bitmap) {
	*f &= ^flag
}

func (f *bitmap) Toggle(flag bitmap) {
	*f ^= flag
}

func (f bitmap) UInt() uint {
	return uint(f)
}
