package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	letter string
	first  int
	second int
}

func main() {
	policyType := 1
	if len(os.Args) > 1 {
		policyType, _ = strconv.Atoi(os.Args[1])
	}

	if policyType != 1 && policyType != 2 {
		panic("Unsupported password policy type")
	}

	fmt.Println("Policy Type ", policyType)

	data := readDataLines("official.data")

	valid := 0

	for _, line := range data {
		parts := strings.SplitN(line, ":", 2)
		password := parts[1]
		if len(parts) < 2 {
			panic("A line had an improperly formatted policy")
		}

		policy := parsePolicy(parts[0])
		if policyType == 1 {
			if checkPasswordAgainstPolicy1(password, policy) {
				valid++
			}
		} else if policyType == 2 {
			if checkPasswordAgainstPolicy2(password, policy) {
				valid++
			}
		}

	}

	fmt.Printf("Valid passwords: %v/%v\n", valid, len(data))
}

func parsePolicy(policyString string) passwordPolicy {
	parts := strings.Fields(policyString)
	allowedRange := strings.Split(parts[0], "-")
	first, _ := strconv.Atoi(allowedRange[0])
	second, _ := strconv.Atoi(allowedRange[1])

	return passwordPolicy{
		letter: parts[1],
		first:  first,
		second: second,
	}
}

func checkPasswordAgainstPolicy1(password string, policy passwordPolicy) bool {
	count := strings.Count(password, policy.letter)

	return policy.first <= count && count <= policy.second
}

func checkPasswordAgainstPolicy2(password string, policy passwordPolicy) bool {
	return (string(password[policy.first]) == policy.letter) != (string(password[policy.second]) == policy.letter)
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
