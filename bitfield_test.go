package bitfield_test

import (
	"github.com/josestg/bitfield"
	"math"
	"testing"
)

func TestInvert(t *testing.T) {
	f := bitfield.BitField(0)
	g := bitfield.Invert(f)
	if g != math.MaxUint64 {
		t.Error("Expect all bits set")
	}

	// Invert of inverted bitfield should be back to original.
	if bitfield.Invert(g) != f {
		t.Error("Expect all bits unset")
	}
}

func TestUnion(t *testing.T) {
	f := bitfield.BitField(0b0011_0011)
	g := bitfield.BitField(0b0101_0101)
	h := bitfield.Union(f, g)
	if h != 0b0111_0111 {
		t.Error("Expect 0b0111_0111")
	}
}

func TestIntersection(t *testing.T) {
	f := bitfield.BitField(0b0011_0011)
	g := bitfield.BitField(0b0101_0101)
	h := bitfield.Intersection(f, g)
	if h != 0b0001_0001 {
		t.Error("Expect 0b0001_0001")
	}
}

func TestDifference(t *testing.T) {
	f := bitfield.BitField(0b0011_0011)
	g := bitfield.BitField(0b0101_0101)
	h := bitfield.Difference(f, g)
	if h != 0b0010_0010 {
		t.Error("Expect 0b0010_0010")
	}
}

func TestBitField_AllSet(t *testing.T) {
	f := bitfield.BitField(0)
	if f.AllSet() {
		t.Error("Expect false")
	}
	f = bitfield.BitField(math.MaxUint64)
	if !f.AllSet() {
		t.Error("Expect true")
	}
}

func TestBitField_NotEmpty(t *testing.T) {
	f := bitfield.BitField(0)
	if f.NotEmpty() {
		t.Error("Expect false")
	}
	f = bitfield.BitField(0b0000_1000)
	if !f.NotEmpty() {
		t.Error("Expect true")
	}
}

func TestBitField_Empty(t *testing.T) {
	f := bitfield.BitField(0)
	if !f.Empty() {
		t.Error("Expect true")
	}
	f = bitfield.BitField(0b0000_1000)
	if f.Empty() {
		t.Error("Expect false")
	}
}

func TestBitField_SetBit(t *testing.T) {
	var f bitfield.BitField
	f = f.SetBit(0).SetBit(1).SetBit(2)
	if f != 0b0000_0111 {
		t.Error("Expect 0b0000_0111")
	}

	// position out of range should be ignored.
	f = f.SetBit(64)
	if f != 0b0000_0111 {
		t.Error("Expect no effect")
	}
}

func TestBitField_IsSet(t *testing.T) {
	// empty bitfield should return false for all positions.
	var f bitfield.BitField
	for i := 0; i < 64; i++ {
		if f.IsSet(uint8(i)) {
			t.Errorf("Expect false for %d", i)
		}
	}

	// expect true for positions 0, 1, 2.
	f = 0b0000_0111
	for i := 0; i < 3; i++ {
		if !f.IsSet(uint8(i)) {
			t.Errorf("Expect true for %d", i)
		}
	}

	// position out of range should return false.
	if f.IsSet(64) {
		t.Error("Expect false for 64")
	}
}

func TestBitField_DelBit(t *testing.T) {
	f := bitfield.BitField(0b1010_1010)
	f = f.DelBit(1)
	if f != 0b1010_1000 {
		t.Error("Expect 0b1010_1000")
	}

	f = f.DelBit(7)
	if f != 0b0010_1000 {
		t.Error("Expect 0b0010_1000")
	}

	f = f.DelBit(3)
	if f != 0b0010_0000 {
		t.Error("Expect 0b0010_0000")
	}

	// position out of range should be ignored.
	f = f.DelBit(64)
	if f != 0b0010_0000 {
		t.Error("Expect no effect")
	}
}

func TestBitField_Cardinal(t *testing.T) {
	var f bitfield.BitField
	if f.Cardinal() != 0 {
		t.Error("Expect cardinality of enpty bitfield to be 0")
	}

	f = 0b0000_0111
	if f.Cardinal() != 3 {
		t.Error("Expect cardinality of 0b0000_0111 to be 3")
	}

	g := bitfield.Invert(f)
	if g.Cardinal() != (64 - f.Cardinal()) {
		t.Error("Expect cardinality of inverted bitfield to be 64 - cardinality of original bitfield")
	}
}
