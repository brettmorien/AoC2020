package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const debug = false

func main() {
	data := readData("official.data")

	weakness := findWeakness(data, 25)
	fmt.Printf("Weakness: %v\n", weakness)

	set := findContiguous(weakness, data)

	encWeakness := min(set) + max(set)
	fmt.Printf("Weakness: %v - Found set: %v\n", encWeakness, set)
}

func findWeakness(data []int, windowSize int) int {
	window := make([]int, windowSize)

	for i := 0; i < windowSize; i++ {
		window[i] = data[i]
	}

	for i, val := range data[windowSize:] {
		valid := validate(val, window)
		if !valid {
			return val
		}
		window[i%windowSize] = val
	}
	return 0
}

func findContiguous(total int, data []int) []int {
	for begin := range data {
		for end := range data[begin+2:] {
			sum := 0
			for _, i := range data[begin : end+begin] {
				sum += i
			}
			if sum == total {
				return data[begin : end+begin]
			}
			if sum > total {
				break
			}
		}
	}
	return []int{}
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

func validate(number int, window []int) bool {
	for _, val := range window {
		for _, val2 := range window {
			if val+val2 == number {
				return true
			}
		}
	}

	return false
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readData(filename string) []int {
	dataLines := readDataLines(filename)

	result := make([]int, len(dataLines))
	for i, line := range dataLines {
		result[i], _ = strconv.Atoi(line)
	}
	return result
}

func readDataLines(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	return strings.Split(string(fileBytes), "\n")
}
