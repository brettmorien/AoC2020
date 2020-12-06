package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	bonus := false
	if len(os.Args) > 1 {
		bonus = true
	}

	fmt.Println("Bonus ", bonus)

	right := 3
	down := 1

	tree := rune('#')

	fmt.Printf("Slope - %v right, %v down\n", right, down)

	data := readDataLines("official.data")
	width := len(data[0])

	trees := 0
	position := 0

	for _, line := range data[1:] {

		position = (position + right) % width
		charAtPosition, _ := utf8.DecodeRuneInString(line[position:])

		isTree := charAtPosition == tree

		if isTree {
			trees++
		}
	}

	fmt.Printf("Result: %v trees\n", trees)
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
