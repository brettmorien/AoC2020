package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	policyType := "any"
	if len(os.Args) > 1 {
		policyType = os.Args[1]
	}

	if policyType != "any" && policyType != "all" {
		panic("Unsupported answer policy type. Use \"any\" or \"all\"")
	}

	records := readRecords("official.data")

	total := 0

	getAnswers := getAnyYesAnswers
	if policyType == "all" {
		getAnswers = getAllYesAnswers
	}

	for _, record := range records {
		answers := getAnswers(record)
		total += answers
	}
	fmt.Println("Total answers", total)
}

func getAllYesAnswers(record string) int {
	values := map[string]bool{}

	m1 := regexp.MustCompile(`[^a-z]`)
	allAnswers := m1.ReplaceAllString(record, "")

	responses := strings.Split(record, "\n")
	for _, a := range allAnswers {
		all := true
		for _, r := range responses {
			all = all && (strings.IndexRune(r, a) != -1)
		}
		if all {
			values[string(a)] = true
		}
	}

	return len(values)
}

func getAnyYesAnswers(record string) int {
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
