package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode/utf8"
)

var tree = rune('#')

type slope struct {
	right int
	down  int
}

func main() {
	data := readDataLines("official.data")

	slopes := []slope{
		slope{right: 1, down: 1},
		slope{right: 3, down: 1},
		slope{right: 5, down: 1},
		slope{right: 7, down: 1},
		slope{right: 1, down: 2},
	}

	trees := 1
	for _, s := range slopes {
		trees *= findTrees(data, s)
	}

	println("Trees:", trees)
}

func findTrees(data []string, slope slope) (trees int) {
	width := len(data[0])
	position := 0

	for _, line := range data[1:] {

		position = (position + slope.right) % width
		charAtPosition, _ := utf8.DecodeRuneInString(line[position:])

		isTree := charAtPosition == tree

		if isTree {
			trees++
		}
	}

	fmt.Printf("Result of run %vx%v: %v trees\n", slope.right, slope.down, trees)

	return trees
}

func replaceAtIndex(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
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
