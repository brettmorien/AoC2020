package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const debug = true

type WaitingArea struct {
	seats    [][]bool
	occupied [][]bool
}

func main() {
	data := readData("sample.data")

	waitingArea := fill(data)

	if debug {
		waitingArea.draw()
	}

	count := 0

	fmt.Printf("Count: %v\n", count)
}

func (area WaitingArea) draw() {
	for i, row := range area.seats {
		for j, seat := range row {
			if !seat {
				fmt.Print(".")
			} else {
				if area.occupied[i][j] {
					fmt.Print("#")
				} else {
					fmt.Print("L")
				}
			}
		}
		fmt.Println()
	}
}

func fill(seats [][]bool) WaitingArea {
	area := WaitingArea{
		seats:    seats,
		occupied: make([][]bool, len(seats)),
	}
	width := len(seats[0])
	for i := range area.seats {
		area.occupied[i] = make([]bool, width)
	}

	return area
}

func readData(filename string) [][]bool {
	dataLines := readDataLines(filename)

	width := len(dataLines[0])

	result := make([][]bool, len(dataLines))
	for i, line := range dataLines {
		result[i] = make([]bool, width)
		for j, seat := range line {
			if seat == 'L' {
				result[i][j] = true
			}
		}
	}
	return result
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
