package main

import (
	"fmt"
	"testing"
)

func TestGetDir(t *testing.T) {

	dir := getDir("R25")
	if dir != "R" {
		t.Error("Error: expected 'R', got '" + dir + ".'")
	}
}

func TestGetDist(t *testing.T) {

	dist := getDist("R255")
	if dist != 255 {
		t.Error("Error: expected 255, got %d.", dist)
	}
}

func TestAddNextCoords(t *testing.T) {

	t1 := "L12"
	t2 := "U7"
	s := []int{0, 0}
	var ta [][]int
	ta = append(ta, s)

	a := addNextCoords(getDir(t1), getDist(t1), ta)
	b := addNextCoords(getDir(t2), getDist(t2), a)

	if len(a) < 13 || len(a) > 13 {
		t.Error("Failed to load array properly.")
	}
	if a[6][0] != -6 {
		t.Error("Expected index incorrect, array failed to load correctly.")
	}
	if len(b) < 20 || len(b) > 20 {
		t.Error("Failed to Load Array properly, length incorrect.")
	}
	if b[19][0] != -12 && b[19][1] != 7 {
		t.Error("Indexes incorrect, array failed to load correct values.")
	}
}

func TestLoadAllCoords(t *testing.T) {

	p := []string{"R3", "U4", "L2", "D1"}
	s := []int{0, 0}
	var tp [][]int
	tp = append(tp, s)

	a := loadAllCoords(p, tp)

	if len(a) < 11 || len(a) > 11 {
		t.Error("Arrays did not load correctly.")
	}
	if a[4][0] != 3 {
		t.Error("Incorrect x coord, arrays did not load properly.")
	}
}

func TestFindIntersections(t *testing.T) {

	// example case from challenge.
	p1 := []string{"R8", "U5", "L5", "D3"}
	p2 := []string{"U7", "R6", "D4", "L4"}
	r1 := []int{6, 5}
	r2 := []int{3, 3}

	var ta [][]int
	s := []int{0, 0}
	ta = append(ta, s)

	fp1 := loadAllCoords(p1, ta)
	fp2 := loadAllCoords(p2, ta)

	t1 := findIntersections(fp1, fp2)

	if t1[0][0] != r1[0] || t1[0][1] != r1[1] {
		t.Error("Matching coordinates not correctly identified.")
	}
	if t1[1][0] != r2[0] || t1[1][1] != r2[1] {
		t.Error("Matching coordinates not correctly identified.")
	}
}

func TestFindLowest(t *testing.T) {

	// example case from challenge.
	p1 := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	p2 := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}

	var ta [][]int
	s := []int{0, 0}
	ta = append(ta, s)

	fp1 := loadAllCoords(p1, ta)
	fp2 := loadAllCoords(p2, ta)

	t1 := findIntersections(fp1, fp2)
	fmt.Println(t1)

	t2 := findLowest(t1)
	fmt.Println(t2)

	if t2 != 159 {
		t.Error("Correct manhatten distance not returned")
	}

}
