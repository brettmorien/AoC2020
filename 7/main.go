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
	data := readDataLines("official.data")

	rules := make([]Rule, len(data))

	for i, line := range data {
		rules[i] = parseRule(line)
	}

	ruleMap := map[string]Rule{}

	for _, rule := range rules {
		fmt.Printf("%v\n", rule)

		ruleMap[rule.bagColor] = rule
	}

	desiredColor := "shiny gold"
	colorMap := map[string]bool{}
	colorMap[desiredColor] = true
	count := len(colorMap)
	for {
		for color := range colorMap {
			containingColors := findBagsContainingColor(color, rules)
			for _, c := range containingColors {
				colorMap[c] = true
			}
		}

		// harmlessly remove initial color from collection
		delete(colorMap, desiredColor)

		// leave if the list hasn't grown
		if len(colorMap) == count {
			break
		}
		count = len(colorMap)
	}

	colors := make([]string, 0, len(colorMap))
	for c := range colorMap {
		colors = append(colors, c)
	}

	fmt.Println("------------")
	fmt.Printf("%v colors: %#v\n", len(colors), colors)
}

func findBagsContainingColor(color string, rules []Rule) []string {
	var results []string
	for _, rule := range rules {
		for _, c := range rule.constraints {
			if c.color == color {
				results = append(results, rule.bagColor)
			}
		}
	}

	return results
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
