package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("reading")

	fileBytes, error := ioutil.ReadFile("official.data")
	check(error)
	dataLines := strings.Split(string(fileBytes), "\n")
	var data []int = make([]int, len(dataLines))
	for i, v := range dataLines {
		data[i], _ = strconv.Atoi(v)
	}

	fmt.Println(len(data))

	for i, a := range data {
		for _, b := range data[i:] {
			result := a + b
			fmt.Printf("%v + %v = %v\n", a, b, result)
			if result == 2020 {
				fmt.Printf("Found 2020 result: %v + %v: %v\n", a, b, a*b)
				return
			}
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
