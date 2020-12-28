package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const debug = false

func main() {
	data := readData("official.data")

	if debug {
		fmt.Printf("%v\n", data)
	}

	findGaps(data)
}

func findGaps(data []int) {
	gaps := map[int]int{}
	gaps[3] = 1
	gaps[data[0]]++

	for i, val := range data[1:] {
		gaps[val-data[i]]++
	}

	fmt.Printf("Gaps: %v\n", gaps)
	fmt.Printf("Answer: %v\n", gaps[1]*gaps[3])
}

func readData(filename string) []int {
	dataLines := readDataLines(filename)

	result := make([]int, len(dataLines))
	for i, line := range dataLines {
		result[i], _ = strconv.Atoi(line)
	}
	sort.Ints(result)
	return result
}

func readDataLines(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	return strings.Split(string(fileBytes), "\n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
