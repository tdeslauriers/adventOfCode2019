package main

import (
	"strconv"
)

// challenge: https://adventofcode.com/2019/day/4
// challenge input: 171309-643603
func main() {

	modUpNum(123456)
}

func gTeT(s0, s1 int) bool {

	if s0 >= s1 {
		return true
	}

	return false
}

func cutUpNum(i int) []int {

	s := strconv.Itoa(i)
	r := []rune(s)
	var j []int

	for k := 0; k < len(r); k++ {

		d := string(r[k])
		e, _ := strconv.Atoi(d)
		j = append(j, e)
	}

	return j
}
