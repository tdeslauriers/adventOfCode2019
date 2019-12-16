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

	file, _ := filepath.Abs("./src/adventOfCode2019/Day1/dayOneInput.txt")
	total := calculateWeight(readFileToArray(file))
	fmt.Println(total)
}

func readFileToArray(filename string) []string {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	return lines
}

func calculateWeight(modules []string) int {

	total := 0
	for _, s := range modules {
		t, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		t = t / 3
		t = t - 2
		total += t
	}

	return total
}
