package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const debug = true

type Ship struct {
	heading int
	x       int
	y       int
}

type Instruction struct {
	op  string
	arg int
}

func main() {
	instructions := readData("official.data")

	ship := Ship{
		heading: 0,
		x:       0,
		y:       0,
	}

	for _, instruction := range instructions {
		ship.move(instruction)
	}

	fmt.Printf("Manhattan Distance: %v\n", abs(ship.x)+abs(ship.y))

}

func toXY(degrees int) (x int, y int) {
	if degrees == 0 {
		return 1, 0
	} else if degrees == 90 {
		return 0, 1
	} else if degrees == 180 {
		return -1, 0
	} else if degrees == 270 {
		return 0, -1
	} else {
		panic(fmt.Sprintf("Unknown direction: %v", degrees))
	}
}

func (ship *Ship) move(instruction Instruction) {
	if instruction.op == "N" {
		ship.y += instruction.arg
	} else if instruction.op == "E" {
		ship.x += instruction.arg
	} else if instruction.op == "S" {
		ship.y -= instruction.arg
	} else if instruction.op == "W" {
		ship.x -= instruction.arg
	} else if instruction.op == "R" {
		ship.heading = (ship.heading - instruction.arg) % 360
	} else if instruction.op == "L" {
		ship.heading = (ship.heading + instruction.arg) % 360
	} else if instruction.op == "F" {
		x, y := toXY(ship.heading)
		ship.x += x * instruction.arg
		ship.y += y * instruction.arg
	}

	if ship.heading < 0 {
		ship.heading += 360
	}

	if debug {
		fmt.Printf("%v\n", instruction)
		fmt.Printf("%v\n", *ship)
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
