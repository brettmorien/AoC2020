package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	letter string
	min    int
	max    int
}

func main() {
	data := readDataLines("official.data")

	valid := 0

	for _, line := range data {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) < 2 {
			panic("A line had an improperly formatted policy")
		}

		policy := parsePolicy(parts[0])
		if checkPasswordAgainstPolicy(parts[1], policy) {
			valid++
		}
	}

	fmt.Printf("Valid passwords: %v/%v\n", valid, len(data))
}

func parsePolicy(policyString string) passwordPolicy {
	parts := strings.Split(policyString, " ")
	allowedRange := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(allowedRange[0])
	max, _ := strconv.Atoi(allowedRange[1])

	return passwordPolicy{
		letter: parts[1],
		min:    min,
		max:    max,
	}
}

func checkPasswordAgainstPolicy(password string, policy passwordPolicy) bool {
	count := strings.Count(password, policy.letter)

	return policy.min <= count && count <= policy.max
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
