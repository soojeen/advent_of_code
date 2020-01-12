package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	program := parseInput(rawInput)

	resultA := process(program, 1)
	fmt.Println(resultA)

	resultB := process(program, 5)
	fmt.Println(resultB)
}

func parseInput(rawInput string) []int {
	_positions := strings.Split(rawInput, ",")

	positions := make([]int, len(_positions))
	for i := range positions {
		position, e := strconv.Atoi(_positions[i])
		if e != nil {
			log.Fatal(e)
		}

		positions[i] = position
	}

	return positions
}

type computer struct {
	program         []int
	currentPosition int
	halt            bool
	input           int
	output          int
}

func (c *computer) getRawParameter(index int) int {
	return c.program[c.currentPosition+index+1]
}

func (c *computer) getParameter(index int) int {
	value := c.getRawParameter(index)
	rawOpCode := c.program[c.currentPosition]
	mode := 0

	switch index {
	case 0:
		mode = (rawOpCode % 1000) / 100
	case 1:
		mode = (rawOpCode % 10000) / 1000
	}

	if mode == 0 {
		return c.program[value]
	}

	return value
}

func (c *computer) processInstruction() {
	rawOpCode := c.program[c.currentPosition]
	opCode := rawOpCode % 100
	_program := c.program

	switch opCode {
	case 1:
		_program[c.getRawParameter(2)] = c.getParameter(0) + c.getParameter(1)
		c.currentPosition += 4
	case 2:
		_program[c.getRawParameter(2)] = c.getParameter(0) * c.getParameter(1)
		c.currentPosition += 4
	case 3:
		_program[c.getRawParameter(0)] = c.input
		c.currentPosition += 2
	case 4:
		c.output = c.getParameter(0)
		c.currentPosition += 2
	case 5:
		if c.getParameter(0) != 0 {
			c.currentPosition = c.getParameter(1)
		} else {
			c.currentPosition += 3
		}
	case 6:
		if c.getParameter(0) == 0 {
			c.currentPosition = c.getParameter(1)
		} else {
			c.currentPosition += 3
		}
	case 7:
		value := 0
		if c.getParameter(0) < c.getParameter(1) {
			value = 1
		}

		_program[c.getRawParameter(2)] = value
		c.currentPosition += 4
	case 8:
		value := 0
		if c.getParameter(0) == c.getParameter(1) {
			value = 1
		}

		_program[c.getRawParameter(2)] = value
		c.currentPosition += 4
	case 99:
		fallthrough
	default:
		c.halt = true
	}

	return
}

func process(program []int, input int) int {
	programCopy := append(program[:0:0], program...)
	computer := computer{programCopy, 0, false, input, 0}

	for {
		computer.processInstruction()

		if computer.halt {
			break
		}
	}

	return computer.output
}
