package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Op string

const (
	nop Op = "nop"
	acc Op = "acc"
	jmp Op = "jmp"
)

type Program struct {
	instructions []Instruction
}

type Instruction struct {
	op  Op
	arg int
}

func main() {
	program := readProgram("official.data")
	execute(program)
}

func execute(program Program) {
	executed := map[int]bool{}

	a := 0
	pc := 0
	for !executed[pc] {
		inst := program.instructions[pc]
		executed[pc] = true
		fmt.Printf("%v\n", inst)
		if inst.op == jmp {
			pc += inst.arg
		} else if inst.op == acc {
			a += inst.arg
			pc++
		} else if inst.op == nop {
			pc++
		}
	}

	fmt.Printf("Acc before loop: %v\n", a)
}

func readProgram(file string) Program {
	lines := readDataLines(file)

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}
	return Program{
		instructions: instructions,
	}
}

func parseInstruction(instruction string) Instruction {
	tokens := strings.Split(instruction, " ")

	arg, _ := strconv.Atoi(tokens[1])
	return Instruction{
		op:  Op(tokens[0]),
		arg: arg,
	}
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
