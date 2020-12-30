package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const debug = false

// [][]bool -> row, seat
type WaitingArea struct {
	rowCount  int
	seatCount int
	seats     [][]bool
	occupied  [][]bool
}

func main() {
	data := readData("official.data")

	waitingArea := fill(data)

	if debug {
		waitingArea.draw()
	}

	iterations := 0
	for {
		newWaitingArea := waitingArea.takeTurn()
		differences := waitingArea.diffOccupied(newWaitingArea)
		if differences == 0 {
			break
		}
		waitingArea = newWaitingArea
		iterations++
	}

	fmt.Printf("Iterations: %v\n", iterations)
	fmt.Printf("Occupied: %v\n", waitingArea.occupiedCount())
}

func (area WaitingArea) takeTurn() WaitingArea {
	newWaitingArea := fill(area.seats)

	for i, row := range area.seats {
		for j, seat := range row {
			if !seat {
				continue
			}

			adjacent := area.countOccupied(i, j)
			if area.occupied[i][j] {
				newWaitingArea.occupied[i][j] = adjacent < 5
			} else {
				newWaitingArea.occupied[i][j] = adjacent == 0
			}
		}
	}

	if debug {
		newWaitingArea.draw()
	}
	return newWaitingArea
}

func (area WaitingArea) castRay(row int, seat int, dirRow int, dirSeat int) bool {
	distance := 1
	checkRow, checkSeat := row+dirRow, seat+dirSeat
	for checkRow >= 0 && checkRow < area.rowCount && checkSeat >= 0 && checkSeat < area.seatCount {
		if area.occupied[checkRow][checkSeat] {
			return true
		}
		if area.seats[checkRow][checkSeat] {
			return false
		}
		distance++
		checkRow, checkSeat = row+(distance*dirRow), seat+(distance*dirSeat)
	}

	return false
}

const (
	up    = -1
	down  = 1
	left  = -1
	right = 1
	same  = 0
)

func (area WaitingArea) countOccupied(row int, seat int) int {
	occupied := 0

	if area.castRay(row, seat, up, left) {
		occupied++
	}
	if area.castRay(row, seat, up, same) {
		occupied++
	}
	if area.castRay(row, seat, up, right) {
		occupied++
	}
	if area.castRay(row, seat, same, left) {
		occupied++
	}
	if area.castRay(row, seat, same, right) {
		occupied++
	}
	if area.castRay(row, seat, down, left) {
		occupied++
	}
	if area.castRay(row, seat, down, same) {
		occupied++
	}
	if area.castRay(row, seat, down, right) {
		occupied++
	}

	return occupied
}

func (area WaitingArea) countAdjacent(row int, seat int) int {
	adjacent := 0

	if row > 0 {
		row := area.occupied[row-1]
		if seat > 0 && row[seat-1] {
			adjacent++
		}
		if row[seat] {
			adjacent++
		}
		if seat < area.seatCount-1 && row[seat+1] {
			adjacent++
		}
	}

	if seat > 0 && area.occupied[row][seat-1] {
		adjacent++
	}
	if seat < area.seatCount-1 && area.occupied[row][seat+1] {
		adjacent++
	}

	if row < area.rowCount-1 {
		row := area.occupied[row+1]
		if seat > 0 && row[seat-1] {
			adjacent++
		}
		if row[seat] {
			adjacent++
		}
		if seat < area.seatCount-1 && row[seat+1] {
			adjacent++
		}
	}

	return adjacent
}

func (area WaitingArea) occupiedCount() int {
	occupied := 0
	for _, row := range area.occupied {
		for _, seat := range row {
			if seat {
				occupied++
			}
		}
	}
	return occupied
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
	fmt.Println("------------")
}

func (area WaitingArea) diffOccupied(area2 WaitingArea) int {
	differences := 0
	for i, row := range area.occupied {
		for j, seat := range row {
			if area2.occupied[i][j] != seat {
				differences++
			}
		}
	}
	return differences
}

func fill(seats [][]bool) WaitingArea {
	area := WaitingArea{
		seats:    seats,
		occupied: make([][]bool, len(seats)),
		rowCount: len(seats),
	}
	area.seatCount = len(seats[0])
	for i := range area.seats {
		area.occupied[i] = make([]bool, area.seatCount)
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
