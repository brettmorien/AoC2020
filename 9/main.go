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
	windowSize := 25

	window := make([]int, windowSize)

	for i := 0; i < windowSize; i++ {
		window[i] = data[i]
	}

	for i, val := range data[windowSize:] {
		valid := validate(val, window)
		if !valid {
			fmt.Printf("Not valid: %v\n", val)
		}
		window[i%windowSize] = val
	}

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
