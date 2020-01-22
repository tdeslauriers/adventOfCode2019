package main

import (
	"testing"
)

func Test_gTeT(t *testing.T) {

	w := 3
	x := 3
	y := 4
	z := 5

	if gTeT(x, y) && !gTeT(z, y) && !gTeT(w, x) {
		t.Error("Greater-than equal-to check incorrect.")
	}

}

func Test_cutUpNum(t *testing.T) {

	// challenge example inputs.
	a := cutUpNum(111111) // meets these criteria (double 11, never decreases).
	b := cutUpNum(223450) // does not meet these criteria (decreasing pair of digits 50).
	c := cutUpNum(123789) // does not meet these criteria (no double).

	if len(a) != 6 || len(b) != 6 || len(c) != 6 {
		t.Error("Arrays not loaded correctly:\n",
			"\texpected array length 7, got ", len(a), "\n",
			"\texpected array length 7, got ", len(b), "\n",
			"\texpected array length 7, got ", len(c))
	}
	if a[0] != 1 || b[1] != 2 || c[2] != 3 {
		t.Error("Arrays not loaded correctly:\n",
			"\texpected a[0] == 1, got ", a[0], "\n",
			"\texpected b[1] == 2, got ", b[1], "\n",
			"\texpected c[2] == 3, got ", c[2])
	}
}

func Test_notDecrease(t *testing.T) {

	// challenge example inputs.
	a := 111111 // meets these criteria (double 11, never decreases).
	b := 223450 // does not meet these criteria (decreasing pair of digits 50).
	c := 123789 // does not meet these criteria (no double).

	if !notDecrease(cutUpNum(a)) {
		t.Error("Expected input of 111111 to yield true.")
	}

	if notDecrease(cutUpNum(b)) {
		t.Error("Expected input of 223450 to yield false.")
	}

	if !notDecrease(cutUpNum(c)) {
		t.Error("Expected input of 123789 to yield true.")
	}
}

func Test_hasPair(t *testing.T) {

	// challenge example inputs
	a := 111111 // meets these criteria (double 11, never decreases).
	b := 223450 // does not meet these criteria (decreasing pair of digits 50).
	c := 123789 // does not meet these criteria (no double).

	if !hasPair(cutUpNum(a)) {
		t.Error("Expected input of 111111 to yield true.")
	}

	if !hasPair(cutUpNum(b)) {
		t.Error("Expected input of 223450 to yield true.")
	}

	if hasPair(cutUpNum(c)) {
		t.Error("Expected input of 123789 to yield false.")
	}

}

func Test_countPwds(t *testing.T) {

	a := []int{111111, 223450, 123789, 123799}
	b := countPwds(a)

	if b != 2 {
		t.Error("Given the input ", a,
			"\ncount of possible passwords should be 2: ",
			"\npositions a[0]: ", a[0], "and a[3]: ", a[3])
	}
}

func Test_loadChallenge(t *testing.T) {

	b := 5
	e := 15

	a := loadChallenge(b, e)

	if len(a) != 11 {
		t.Error("Expected length of a to be 11, got ", len(a))
	}

	if a[0] != 5 || a[2] != 7 || a[10] != 15 {
		t.Error("Challenge load failed: ",
			"\nExpected a[0] = 5, got ", a[0],
			"\nExpected a[2] = 7, got ", a[2],
			"\nExpected a[11] = 15, got ", a[10])
	}

}
