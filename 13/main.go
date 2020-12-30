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
	timestamp, lines := readData("official.data")

	waits := map[int]int{}
	for _, line := range lines {
		waits[line] = line - timestamp%line
	}

	minWait, line := 100000000000, 0
	for l, wait := range waits {
		if wait < minWait {
			minWait = wait
			line = l
		}
	}

	fmt.Printf("time: %v, lines: %v\n", timestamp, lines)
	fmt.Printf("Wait: %v, line: %v\n", minWait, line)
	fmt.Printf("Answer: %v\n", minWait*line)
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

func min(items []int) int {
	if len(items) == 0 {
		return 0
	}
	min := items[0]

	for _, i := range items {
		if i < min {
			min = i
		}
	}
	return min
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
