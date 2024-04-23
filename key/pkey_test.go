package key

import (
	"testing"
)

func Test_NewPkey0(t *testing.T) {

	pk := NewPKey()
	if len(pk) != 0 {
		t.Fatal("cannot create a new PKey that is empty.  Actual length is not equal to zero.")
	}
}

func Test_NewPkey1(t *testing.T) {

	pk := NewPKey(12)
	if len(pk) != 1 {
		t.Fatal("Creating a new PKey failed.  Actual length is not equal to one.")
	}
}

func Test_Equal0(t *testing.T) {

	pk := NewPKey()

	if !pk.EqualTo(NewPKey()) {
		t.Fatal("Two empty PKeys should be equal.  They are not.")
	}
}

func Test_Equal1(t *testing.T) {

	pk1 := NewPKey(2, 5, 1)
	pk2 := NewPKey(2, 5, 1)

	if !pk1.EqualTo(pk2) {
		t.Fatal("Two identical PKeys should be equal.  They are not.")
	}
}

func Test_Equal2(t *testing.T) {

	pk1 := NewPKey(2, 5, 1)
	pk2 := NewPKey(2, 1, 5)

	if pk1.EqualTo(pk2) {
		t.Fatal("Two PKeys with same number of levels but different indices should not be equal.  They are equal.")
	}
}

func Test_Equal3(t *testing.T) {

	pk1 := NewPKey(2, 5, 1)
	pk2 := NewPKey(2, 1)

	if pk1.EqualTo(pk2) {
		t.Fatal("Two PKeys with different number of levels should not be equal.  They are equal.")
	}
}

// Created an indepent function than PKey.EqualTo to avoid coupling tests
func inequalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}

func Test_AddLevel0(t *testing.T) {

	pk1 := NewPKey()
	pk2 := pk1.AddLevel(89)

	if inequalIntSlices(pk2, []int{89}) {
		t.Fatal("Adding a level to an empty PKey failed.")
	}
}

func Test_AddLevel1(t *testing.T) {

	pk1 := NewPKey(56, 9, 1)
	pk2 := pk1.AddLevel(90)

	if inequalIntSlices(pk2, []int{56, 9, 1, 90}) {
		t.Fatal("Adding a level to an existing PKey failed.")
	}
}

func Test_DescendsFrom0(t *testing.T) {

	pk := NewPKey(56)

	if !pk.DescendsFrom(NewPKey()) {
		t.Fatal("PKey with an index does not descend from empty PKey.")
	}
}

func Test_DescendsFrom1(t *testing.T) {

	pk := NewPKey(56, 9, 1)

	if !pk.DescendsFrom(NewPKey(56, 9)) {
		t.Fatal("PKey doesn't descend from a proper parent PKey.")
	}
}

func Test_DescendsFrom2(t *testing.T) {

	if NewPKey().DescendsFrom(NewPKey()) {
		t.Fatal("A PKey cannot descend from an equivalent PKey.")
	}

	if NewPKey(4, 5, 9).DescendsFrom(NewPKey(4, 5, 9)) {
		t.Fatal("A PKey cannot descend from an equivalent PKey.")
	}
}

func Test_DescendsFrom3(t *testing.T) {

	if NewPKey(2, 90, 1).DescendsFrom(NewPKey(2, 70)) {
		t.Fatal("A PKey with a different upper level index cannot descend.")
	}
}

func Test_IndexAtLevel0(t *testing.T) {

	pk := NewPKey()

	if pk.IndexAtLevel(0) != INVALID_INDEX {
		t.Fatal("Function didn't return an invalid index.")
	}
}

func Test_IndexAtLevel1(t *testing.T) {

	pk := NewPKey(23)

	if pk.IndexAtLevel(0) != 23 {
		t.Error("Function didn't return correct key at level 0.")
	}

	if pk.IndexAtLevel(1) != INVALID_INDEX {
		t.Fatal("Function didn't return an invalid index.")
	}
}

func Test_IndexAtLevel2(t *testing.T) {

	pk := NewPKey(34, 9, 73)

	if pk.IndexAtLevel(0) != 34 {
		t.Error("Function didn't return correct key at level 0.")
	}

	if pk.IndexAtLevel(1) != 9 {
		t.Error("Function didn't return correct key at level 1.")
	}

	if pk.IndexAtLevel(2) != 73 {
		t.Error("Function didn't return correct key at level 2.")
	}

	if pk.IndexAtLevel(3) != INVALID_INDEX {
		t.Error("Function didn't return an invalid index for level 3.")
	}
}
