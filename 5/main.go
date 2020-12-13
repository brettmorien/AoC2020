package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Seat struct {
	row    int
	column int
}

func (seat *Seat) seatID() int {
	return seat.row*8 + seat.column
}

func main() {
	data := readDataLines("official.data")

	seats := make([]Seat, len(data))

	for i, code := range data {
		seats[i] = parseSeat(code)
	}

	// Part 1
	maxSeatID := 0
	for _, seat := range seats {
		seatID := seat.seatID()
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Printf("maximum seat Id: %v\n", maxSeatID)
}

func parseSeat(code string) Seat {
	rowData := strings.ReplaceAll(strings.ReplaceAll(code[:7], "F", "0"), "B", "1")
	columnData := strings.ReplaceAll(strings.ReplaceAll(code[7:], "L", "0"), "R", "1")

	rowNumber, _ := strconv.ParseInt(rowData, 2, 32)
	columnNumber, _ := strconv.ParseInt(columnData, 2, 32)

	return Seat{
		row:    int(rowNumber),
		column: int(columnNumber),
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readDataLines(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	return strings.Split(string(fileBytes), "\n")
}
