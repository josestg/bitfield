package bitfield

// BitField is a 64-bit bitfield that can be used to store the membership status
// of 64 elements. If the bit at position k is 1, then the element k is in the set.
// Otherwise, the element k is not in the set. Where k is in the range [0, 63].
type BitField uint64

// Empty returns true if all bits in the BitField are 0.
func (f BitField) Empty() bool {
	return f == 0
}

// NotEmpty returns true if any bit in the BitField is 1.
func (f BitField) NotEmpty() bool {
	return f != 0
}

// AllSet returns true if all bits in the BitField are 1.
func (f BitField) AllSet() bool {
	return f == 0xffffffffffffffff
}

// SetBit sets the bit at position k in the BitField to 1.
// If k >= 64, SetBit does nothing.
func (f BitField) SetBit(k uint8) BitField {
	g := safeShiftLeft(k)
	return Union(f, g)
}

// DelBit sets the bit at position k in the BitField to 0.
// If k >= 64, DelBit does nothing.
func (f BitField) DelBit(k uint8) BitField {
	g := safeShiftLeft(k)
	return Difference(f, g)
}

// IsSet returns true if the bit at position k in the BitField is 1.
// If k >= 64, IsSet returns false.
func (f BitField) IsSet(pos uint8) bool {
	g := safeShiftLeft(pos)
	return Intersection(f, g).NotEmpty()
}

// Cardinal returns the number of bits that are set to 1 in the BitField.
func (f BitField) Cardinal() uint8 {
	var count uint8
	for i := uint8(0); i < 64; i++ {
		if f.IsSet(i) {
			count++
		}
	}
	return count
}

// safeShiftLeft shifts left 1 by n bits, returning 0 if n >= 64.
func safeShiftLeft(n uint8) BitField {
	if n >= 64 {
		return 0
	}
	return 1 << n
}

// Invert returns the inverse of a BitField.
func Invert(f BitField) BitField { return ^f }

// Union returns the union of two Fields.
func Union(x, y BitField) BitField { return x | y }

// Intersection returns the intersection of two Fields.
func Intersection(x, y BitField) BitField { return x & y }

// Difference returns a new field that all bits that are set in x, but not set in y (x / y).
// Note that x/y != x/y.
func Difference(x, y BitField) BitField {
	z := Intersection(x, y)
	return x ^ z
}
