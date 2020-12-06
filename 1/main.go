package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	depth := 2
	if len(os.Args) > 1 {
		depth, _ = strconv.Atoi(os.Args[1])
	}

	fmt.Println("Depth ", depth)

	fileBytes, error := ioutil.ReadFile("official.data")
	check(error)
	dataLines := strings.Split(string(fileBytes), "\n")
	var data []int = make([]int, len(dataLines))
	for i, v := range dataLines {
		data[i], _ = strconv.Atoi(v)
	}

	if depth == 2 {
		for i, a := range data {
			for _, b := range data[i:] {
				result := a + b
				// fmt.Printf("%v + %v = %v\n", a, b, result)
				if result == 2020 {
					fmt.Printf("Found 2020 result: %v + %v: %v\n", a, b, a*b)
					return
				}
			}
		}
	} else if depth == 3 {
		for i, a := range data {
			for _, b := range data[i:] {
				for _, c := range data[i:] {
					result := a + b + c
					// fmt.Printf("%v + %v = %v\n", a, b, result)
					if result == 2020 {
						fmt.Printf("Found 2020 result: %v + %v + %v: %v\n", a, b, c, a*b*c)
						return
					}
				}
			}
		}
	} else {
		fmt.Printf("Unsupported depth %v\n", depth)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
