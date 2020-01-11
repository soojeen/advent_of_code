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

	// resultA := process(program, 1)
	// fmt.Println(resultA)

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

type instruction struct {
	opCode        int
	rawOpCode     int
	rawParameters []int
}

func (i *instruction) getParameter(position int) (int, int) {
	parameter := i.rawParameters[position]
	mode := 0

	switch position {
	case 0:
		mode = (i.rawOpCode % 1000) / 100
		return parameter, mode
	case 1:
		mode = (i.rawOpCode % 10000) / 1000
		return parameter, mode
	case 2:
		return parameter, mode
	default:
		return 0, 0
	}
}

func (c *computer) parseInstruction() instruction {
	rawOpCode := c.program[c.currentPosition]
	opCode := rawOpCode % 100

	switch opCode {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 7:
		fallthrough
	case 8:
		rawParameters := make([]int, 3)
		for i := range rawParameters {
			rawParameters[i] = c.program[c.currentPosition+i+1]
		}
		return instruction{opCode, rawOpCode, rawParameters}
	case 3:
		fallthrough
	case 4:
		rawParameters := []int{c.program[c.currentPosition+1]}
		return instruction{opCode, rawOpCode, rawParameters}
	case 5:
		fallthrough
	case 6:
		rawParameters := make([]int, 2)
		for i := range rawParameters {
			rawParameters[i] = c.program[c.currentPosition+i+1]
		}
		return instruction{opCode, rawOpCode, rawParameters}
	case 99:
		fallthrough
	default:
		return instruction{opCode, rawOpCode, []int{}}
	}
}

func (c *computer) processInstruction(instruction instruction) {
	_program := c.program

	switch instruction.opCode {
	case 1:
		augend, augendMode := instruction.getParameter(0)
		if augendMode == 0 {
			augend = _program[augend]
		}

		addend, addendMode := instruction.getParameter(1)
		if addendMode == 0 {
			addend = _program[addend]
		}

		targetParameter, _ := instruction.getParameter(2)

		_program[targetParameter] = augend + addend
		c.currentPosition += 4
	case 2:
		multiplicand, multiplicandMode := instruction.getParameter(0)
		if multiplicandMode == 0 {
			multiplicand = _program[multiplicand]
		}

		multiplier, multiplierMode := instruction.getParameter(1)
		if multiplierMode == 0 {
			multiplier = _program[multiplier]
		}

		targetParameter, _ := instruction.getParameter(2)

		_program[targetParameter] = multiplicand * multiplier
		c.currentPosition += 4
	case 3:
		targetParameter, _ := instruction.getParameter(0)

		_program[targetParameter] = c.input
		c.currentPosition += 2
	case 4:
		value, mode := instruction.getParameter(0)
		if mode == 0 {
			value = _program[value]
		}

		c.output = value
		c.currentPosition += 2
	case 5:
		value, mode := instruction.getParameter(0)
		if mode == 0 {
			value = _program[value]
		}

		if value != 0 {
			instruction, instructionMode := instruction.getParameter(1)
			if instructionMode == 0 {
				instruction = _program[instruction]
			}

			_program[c.currentPosition] = instruction
		} else {
			c.currentPosition += 3
		}
	case 6:
		value, mode := instruction.getParameter(0)
		if mode == 0 {
			value = _program[value]
		}

		if value == 0 {
			instruction, instructionMode := instruction.getParameter(1)
			if instructionMode == 0 {
				instruction = _program[instruction]
			}

			_program[c.currentPosition] = instruction
		} else {
			c.currentPosition += 3
		}
	case 7:
		operandA, operandAMode := instruction.getParameter(0)
		if operandAMode == 0 {
			operandA = _program[operandA]
		}

		operandB, operandBMode := instruction.getParameter(1)
		if operandBMode == 0 {
			operandB = _program[operandB]
		}

		targetParameter, _ := instruction.getParameter(2)
		if operandA < operandB {
			_program[targetParameter] = 1
		} else {
			_program[targetParameter] = 0
		}
	case 8:
		operandA, operandAMode := instruction.getParameter(0)
		if operandAMode == 0 {
			operandA = _program[operandA]
		}

		operandB, operandBMode := instruction.getParameter(1)
		if operandBMode == 0 {
			operandB = _program[operandB]
		}

		targetParameter, _ := instruction.getParameter(2)
		if operandA == operandB {
			_program[targetParameter] = 1
		} else {
			_program[targetParameter] = 0
		}
	case 99:
		c.halt = true
	default:
		c.halt = true
	}

	c.program = _program

	return
}

func process(program []int, input int) int {
	computer := computer{program, 0, false, input, 0}

	for {
		instruction := computer.parseInstruction()
		computer.processInstruction(instruction)

		fmt.Println(instruction, computer.output)

		if computer.halt {
			break
		}
	}

	return computer.output
}
