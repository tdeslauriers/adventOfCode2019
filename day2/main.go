package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	file, _ := filepath.Abs("./src/adventOfCode2019/day2/dayTwoInput.txt")
	fp := processInts(readFileToArray(file))
	fmt.Println(fp)
}

func readFileToArray(filename string) []int {

	c, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	n := strings.Split(string(c), ",")

	var m []int
	for _, t := range n {
		q, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}
		m = append(m, q)
	}
	return m
}

func processInts(arr []int) []int {

	for i := 0; arr[i] != 99; i += 4 {

		x := arr[i+1]
		y := arr[i+2]
		z := arr[i+3]

		if arr[i] == 1 {

			r := arr[x] + arr[y]
			arr[z] = r
		} else if arr[i] == 2 {

			r := arr[x] * arr[y]
			arr[z] = r
		}
	}
	return arr
}
