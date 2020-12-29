package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const debug = true

func main() {
	data := readData("sample.data")

	if debug {
		fmt.Printf("%v items: %v\n", len(data), data)
	}

	findGaps(data)

	count := 0
	traverseTree(0, makeMap(data), &count)

	fmt.Printf("Count: %v\n", count)
}

func traverseTree(value int, data map[int]bool, count *int) {
	childCount := 0
	for i := 1; i <= 3; i++ {
		if data[value+i] {
			childCount++
			traverseTree(value+i, data, count)
		}
	}

	if childCount == 0 {
		*count = *count + 1
	}
}

func makeMap(data []int) map[int]bool {
	valueMap := map[int]bool{}
	for _, i := range data {
		valueMap[i] = true
	}
	return valueMap
}

func findGaps(data []int) {
	gaps := map[int]int{}

	for i, val := range data[1:] {
		gaps[val-data[i]]++
	}

	fmt.Printf("Gaps: %v\n", gaps)
	fmt.Printf("Answer: %v\n", gaps[1]*gaps[3])
}

func readData(filename string) []int {
	dataLines := readDataLines(filename)

	result := make([]int, len(dataLines)+2)
	for i, line := range dataLines {
		result[i], _ = strconv.Atoi(line)
	}
	result[len(result)-1] = max(result) + 3
	sort.Ints(result)
	return result
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
