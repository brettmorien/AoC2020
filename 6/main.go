package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	records := readRecords("official.data")

	total := 0
	for _, record := range records {
		answers := getYesAnswers(record)
		total += answers
	}
	fmt.Println("Total answers", total)
}

func getYesAnswers(record string) int {
	values := map[string]bool{}

	m1 := regexp.MustCompile(`[^a-z]`)

	answers := m1.ReplaceAllString(record, "")
	for _, a := range answers {
		values[string(a)] = true
	}

	return len(values)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readRecords(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	records := strings.Split(string(fileBytes), "\n\n")

	return records
}
