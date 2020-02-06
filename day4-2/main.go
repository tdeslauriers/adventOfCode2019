package main

import (
	"fmt"
	"strconv"
)

// challenge: https://adventofcode.com/2019/day/4
// challenge input: 171309-643603
// second challenge
func main() {

	fmt.Println(countPwds(loadChallenge(171309, 643603)))
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

		if !(i[j+1] >= i[j]) {

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

func pairIsLegit(i []int) bool {

	if !hasPair(i) {
		return false
	}

	switch {
	// 112222
	case i[0] == i[1] && i[1] != i[2]:
		return true
	// 122333
	case i[0] != i[1] && i[1] == i[2] && i[2] != i[3]:
		return true
	// 112233
	case i[1] != i[2] && i[2] == i[3] && i[3] != i[4]:
		return true
	// 111223
	case i[2] != i[3] && i[3] == i[4] && i[4] != i[5]:
		return true
	// 111122
	case i[4] == i[5] && i[4] != i[3]:
		return true
	default:
		return false
	}

}

func hasThreeOfAKind(i []int) bool {

	for j := 0; j < len(i)-2; j++ {

		if i[j] == i[j+1] && i[j+1] == i[j+2] {

			return true
		}
	}

	return false
}

func countPwds(i []int) int {

	c := 0
	for j := 0; j < len(i); j++ {

		k := cutUpNum(i[j])
		if notDecrease(k) && pairIsLegit(k) {

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
