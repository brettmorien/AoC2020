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
	timestamp, lines := readData("sample.data")

	fmt.Printf("time: %v, lines: %v\n", timestamp, lines)
}

func readData(filename string) (int, []int) {
	dataLines := readDataLines(filename)

	timestamp, _ := strconv.Atoi(dataLines[0])
	lines := strings.Split(dataLines[1], ",")
	inService := make([]int, 0, len(lines))
	for _, line := range lines {
		if line != "x" {
			lineNumber, _ := strconv.Atoi(line)
			inService = append(inService, lineNumber)
		}
	}

	sort.Ints(inService)
	return timestamp, inService
}

func readDataLines(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	return strings.Split(string(fileBytes), "\n")
}

func makeSet(data []int) map[int]bool {
	valueMap := map[int]bool{}
	for _, i := range data {
		valueMap[i] = true
	}
	return valueMap
}

func max(items []int) int {
	if len(items) == 0 {
		return 0
	}
	max := items[0]

	for _, i := range items {
		if i > max {
			max = i
		}
	}
	return max
}

func abs(i int) int {
	if i < 0 {
		return 0 - i
	}
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
