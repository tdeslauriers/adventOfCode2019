package main

import (
	"testing"
)

func Test_gTeT(t *testing.T) {

	x := 3
	y := 4
	z := 5

	if gTeT(x, y) && !gTeT(z, y) {
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

func Test_NotDecrease(t *testing.T) {

	// challenge example inputs.
	// a := 111111 // meets these criteria (double 11, never decreases).
	// b := 223450 // does not meet these criteria (decreasing pair of digits 50).
	// c := 123789 // does not meet these criteria (no double).
}
