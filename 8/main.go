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

const debug = false

type Program struct {
	instructions []Instruction
}

type Instruction struct {
	op  Op
	arg int
}

func main() {
	program := readProgram("official.data")
	printProgram(program, -1)
	fmt.Println("-------------")
	change := 0
	for _, inst := range program.instructions {
		if inst.op == nop {
			change++
		}
	}

	terminated := false
	for i, inst := range program.instructions {
		if inst.op != acc {
			terminated = execute(modifyProgram(program, i))
			if terminated {
				break
			}
		}
	}
}

func execute(program Program) bool {
	executed := map[int]bool{}
	termination := len(program.instructions)
	a := 0
	pc := 0
	for !executed[pc] && pc < termination {
		inst := program.instructions[pc]
		executed[pc] = true
		if inst.op == jmp {
			pc += inst.arg
		} else if inst.op == acc {
			a += inst.arg
			pc++
		} else if inst.op == nop {
			pc++
		}
	}

	if pc >= termination {
		fmt.Printf("PC: %v, acc: %v\n", pc, a)
		fmt.Println("Terminated")
		return true
	} else {
		return false
	}
}

func modifyProgram(program Program, changeOp int) Program {
	newProgram := Program{
		instructions: make([]Instruction, len(program.instructions)),
	}

	for i, inst := range program.instructions {
		if i == changeOp {
			newOp := nop
			if inst.op == nop {
				newOp = jmp
			}
			newProgram.instructions[i] = Instruction{op: newOp, arg: inst.arg}
		} else {
			newProgram.instructions[i] = Instruction{op: inst.op, arg: inst.arg}
		}
	}

	printProgram(newProgram, changeOp)
	return newProgram
}

func printProgram(program Program, changeOp int) {
	if debug {
		for i, inst := range program.instructions {
			if i == changeOp {
				fmt.Print("-> ")
			}
			fmt.Printf("%v: %v\n", i, inst)
		}
	}
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
