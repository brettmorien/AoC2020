package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const debug = true

type TreeNode struct {
	value    int
	children []*TreeNode
}

func main() {
	data := readData("sample.data")

	if debug {
		fmt.Printf("%v\n", data)
	}

	findGaps(data)

	count := 0
	buildTree(0, data, &count)

	fmt.Printf("Count: %v", count)
}

func buildTree(value int, data []int, count *int) *TreeNode {
	node := TreeNode{
		value: value,
	}
	for _, val := range data {
		if val <= value {
			continue
		} else if val > value+3 {
			break
		}

		node.children = append(node.children, buildTree(val, data, count))
	}

	if len(node.children) == 0 {
		*count = *count + 1
	}
	return &node
}

func findGaps(data []int) {
	gaps := map[int]int{}

	for i, val := range data[1:] {
		gaps[val-data[i]]++
	}

	fmt.Printf("Gaps: %v\n", gaps)
	fmt.Printf("Answer: %v\n", gaps[1]*gaps[3])
}

func readData(filename string) []int {
	dataLines := readDataLines(filename)

	result := make([]int, len(dataLines)+2)
	for i, line := range dataLines {
		result[i], _ = strconv.Atoi(line)
	}
	result[len(result)-1] = max(result) + 3
	sort.Ints(result)
	return result
}

func readDataLines(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	return strings.Split(string(fileBytes), "\n")
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
