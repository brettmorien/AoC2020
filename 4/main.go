package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

func (r *record) validate() bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.

	if !validateYear(r.byr, 1920, 2002) {
		return false
	}
	if !validateYear(r.iyr, 2010, 2020) {
		return false
	}
	if !validateYear(r.eyr, 2020, 2030) {
		return false
	}

	if strings.HasSuffix(r.hgt, "cm") {
		hgt, err := strconv.Atoi(r.hgt[:len(r.hgt)-2])
		if err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.HasSuffix(r.hgt, "in") {
		hgt, err := strconv.Atoi(r.hgt[:len(r.hgt)-2])
		if err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	if match, _ := regexp.MatchString("^#[a-f0-9]{6}$", r.hcl); !match {
		return false
	}

	eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	colorMatch := false
	for _, color := range eyeColors {
		if r.ecl == color {
			colorMatch = true
		}
	}
	if !colorMatch {
		return false
	}

	if match, _ := regexp.MatchString("^[0-9]{9}$", r.pid); !match {
		return false
	}

	return true
}

func validateYear(yearString string, min int, max int) bool {
	if len(yearString) != 4 {
		return false
	}
	year, err := strconv.Atoi(yearString)

	if err != nil || year < min || year > max {
		return false
	}

	return true
}

func main() {
	// data := readData("sample2.data")
	data := readData("official.data")

	records := splitRecords(data)

	valid := 0
	for _, record := range records {
		if record.validate() {
			valid++
			// fmt.Printf("%v\n", record)
		}
	}

	fmt.Printf("Valid passports: %v/%v\n", valid, len(records))
}

func splitRecords(data []string) []record {
	records := make([]record, len(data))

	for i, line := range data {
		records[i] = parseRecord(line)
	}

	return records
}

func parseRecord(data string) record {
	values := map[string]string{}

	for _, token := range strings.Fields(data) {
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
