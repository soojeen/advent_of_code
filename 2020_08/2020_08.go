package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type instruct struct {
	op  string
	arg int
}

type bootCode struct {
	acc  int
	code []instruct
	idx  int
}

func (c *bootCode) runUntilLoop() {
	tracker := make(map[int]bool)

	for {
		if tracker[c.idx] == true {
			break
		}

		tracker[c.idx] = true

		instruct := c.code[c.idx]
		switch instruct.op {
		case "nop":
			c.idx++
		case "acc":
			c.acc += instruct.arg
			c.idx++
		case "jmp":
			c.idx += instruct.arg
		}
	}
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input, parseError := parseInput(rawInput)
	if parseError != nil {
		log.Fatal(parseError)
	}
	resultA := findLoop(input)
	// resultB := countChildren(input, "shiny gold")

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) (bootCode, error) {
	var err error
	lines := strings.Split(input, "\n")
	code := make([]instruct, len(lines))

	for i, line := range lines {
		codeParts := strings.Split(line, " ")
		arg, error := strconv.Atoi(codeParts[1])
		if error != nil {
			err = error
			break
		}

		code[i] = instruct{codeParts[0], arg}
	}

	return bootCode{0, code, 0}, err
}

func findLoop(input bootCode) int {
	input.runUntilLoop()

	return input.acc
}
