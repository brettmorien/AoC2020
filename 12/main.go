package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const debug = false

type ShipPosition struct {
	heading int
	x       int
	y       int
}

type Instruction struct {
	op  string
	arg int
}

func main() {
	data := readData("sample.data")

	for _, line := range data {
		fmt.Printf("%v\n", line)
	}
}

func readData(filename string) []Instruction {
	dataLines := readDataLines(filename)

	result := make([]Instruction, len(dataLines))
	for i, line := range dataLines {
		result[i] = parseInstruction(line)
	}
	return result
}

func parseInstruction(line string) Instruction {
	arg, _ := strconv.Atoi(line[1:])
	return Instruction{
		op:  line[0:1],
		arg: arg,
	}
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
