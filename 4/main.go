package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type record struct {
	byr string // (Birth Year)
	iyr string // (Issue Year)
	eyr string // (Expiration Year)
	hgt string // (Height)
	hcl string // (Hair Color)
	ecl string // (Eye Color)
	pid string // (Passport ID)
	cid string // (Country ID)
}

func main() {
	data := readData("sample.data")

	records := splitRecords(data)

	fmt.Printf("%v\n", records[1])
}

func splitRecords(data []string) []record {
	records := make([]record, len(data))

	for i, line := range data {
		records[i] = parseRecord(line)
	}

	return records
}

func parseRecord(data string) record {
	tokens := strings.Fields(data)

	values := map[string]string{}

	for _, token := range tokens {
		parts := strings.Split(token, ":")

		values[parts[0]] = parts[1]
	}

	return record{
		byr: values["byr"],
		iyr: values["iyr"],
		eyr: values["eyr"],
		hgt: values["hgt"],
		hcl: values["hcl"],
		ecl: values["ecl"],
		pid: values["pid"],
		cid: values["cid"],
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readData(filename string) []string {
	fileBytes, error := ioutil.ReadFile(filename)
	check(error)
	records := strings.Split(string(fileBytes), "\n\n")

	return records
}
