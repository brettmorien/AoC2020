package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const debug = false

type Point struct {
	x int
	y int
}

type Ship struct {
	position Point
	waypoint Point
}

type Instruction struct {
	op  string
	arg int
}

func main() {
	instructions := readData("official.data")

	ship := Ship{
		position: Point{},
		waypoint: Point{x: 10, y: 1},
	}

	for _, instruction := range instructions {
		ship.move(instruction)
	}

	fmt.Printf("Manhattan Distance: %v\n", abs(ship.position.x)+abs(ship.position.y))
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
		ship.waypoint.y += instruction.arg
	} else if instruction.op == "E" {
		ship.waypoint.x += instruction.arg
	} else if instruction.op == "S" {
		ship.waypoint.y -= instruction.arg
	} else if instruction.op == "W" {
		ship.waypoint.x -= instruction.arg
	} else if instruction.op == "R" {
		ship.rotateWaypoint(-instruction.arg)
	} else if instruction.op == "L" {
		ship.rotateWaypoint(instruction.arg)
	} else if instruction.op == "F" {
		ship.position.x += ship.waypoint.x * instruction.arg
		ship.position.y += ship.waypoint.y * instruction.arg
	}

	if debug {
		fmt.Printf("%v\n", instruction)
		fmt.Printf("%v\n", *ship)
	}
}

func (ship *Ship) rotateWaypoint(degrees int) {
	if degrees < 0 {
		degrees += 360
	}
	if degrees == 90 {
		ship.waypoint = Point{-ship.waypoint.y, ship.waypoint.x}
	}
	if degrees == 180 {
		ship.waypoint = Point{-ship.waypoint.x, -ship.waypoint.y}
	}
	if degrees == 270 {
		ship.waypoint = Point{ship.waypoint.y, -ship.waypoint.x}
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
