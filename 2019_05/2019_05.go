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

	input := parseInput(rawInput)

	input[1] = 12
	input[2] = 2

	resultA := process(input)
	fmt.Println(resultA[0])

	noun, verb := findInputs(input)
	fmt.Println(noun*100 + verb)
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

func process(codes []int) []int {
	_codes := append([]int(nil), codes...)
	currentPosition := 0

	for _codes[currentPosition] != 99 {
		opCode := _codes[currentPosition]
		inputA := _codes[_codes[currentPosition+1]]
		inputB := _codes[_codes[currentPosition+2]]
		outputPosition := _codes[currentPosition+3]

		switch opCode {
		case 1:
			_codes[outputPosition] = inputA + inputB
		case 2:
			_codes[outputPosition] = inputA * inputB
		}

		currentPosition += 4
	}

	return _codes
}

func findInputs(input []int) (int, int) {
	var noun int
	var verb int
	done := false

	for i := 0; i < 100; i++ {
		if done {
			break
		}

		for j := 0; j < 100; j++ {
			if done {
				break
			}

			input[1] = i
			input[2] = j

			result := process(input)

			if result[0] == 19690720 {
				noun = i
				verb = j
				done = true
			}
		}
	}

	return noun, verb
}
