package main

import (
	"fmt"
	"strconv"
)

// challenge: https://adventofcode.com/2019/day/4
// challenge input: 171309-643603
func main() {

	fmt.Println(countPwds(loadChallenge(171309, 643603)))
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

func notDecrease(i []int) bool {

	for j := 0; j < len(i)-1; j++ {

		if !gTeT(i[j+1], i[j]) {

			return false
		}
	}

	return true
}

func hasPair(i []int) bool {

	for j := 0; j < len(i)-1; j++ {

		if i[j] == i[j+1] {

			return true
		}
	}

	return false
}

func countPwds(i []int) int {

	c := 0
	for j := 0; j < len(i); j++ {

		k := cutUpNum(i[j])
		if notDecrease(k) && hasPair(k) {

			c++
		}
	}

	return c
}

func loadChallenge(begin, end int) []int {

	a := make([]int, end-begin+1)
	for i := range a {

		a[i] = begin + i
	}
	return a
}
