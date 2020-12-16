package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule struct {
	bagColor    string
	constraints []Constraint
}

type Constraint struct {
	count int
	color string
}

func main() {
	data := readDataLines("sample.data")

	rules := make([]Rule, len(data))

	for i, line := range data {
		rules[i] = parseRule(line)
	}

	for _, rule := range rules {
		fmt.Printf("%v\n", rule)

		// for _, c := range rule.constraints {
		// 	println("shiny gold ?", c.color)
		// 	if c.color == "shiny gold" {
		// 		total++
		// 	}
		// }
	}

}

func parseRule(ruleString string) Rule {
	rule := Rule{
		bagColor: ruleString[:strings.Index(ruleString, "bags")-1],
	}

	ruleString = ruleString[strings.Index(ruleString, "contain")+8:] // everything after contain

	if strings.Contains(ruleString, "no other bags") {
		return rule
	}

	constraintTokens := strings.Split(ruleString, ", ")

	rule.constraints = make([]Constraint, len(constraintTokens))
	for i, token := range constraintTokens {
		count, _ := strconv.Atoi(token[:strings.Index(token, " ")])
		rule.constraints[i] = Constraint{
			count: count,
			color: token[strings.Index(token, " ")+1 : strings.Index(token, "bag")-1],
		}
	}
	return rule
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
