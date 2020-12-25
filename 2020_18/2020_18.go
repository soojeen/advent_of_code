package main

import "fmt"
import "log"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

const add = "+"
const mul = "*"

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := processA(input)
	resultB := processB(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func processA(input []string) int {
	result := 0

	for _, line := range input {
		value := processLine(line, mathEvalInOrder)
		result += value
	}

	return result
}

func processB(input []string) int {
	result := 0

	for _, line := range input {
		value := processLine(line, mathEvalB)
		result += value
	}

	return result
}

func processLine(input string, mathEval func(string) int) int {
	bracketRe := regexp.MustCompile(`\(([+*\d ]*)\)`)

	for {
		matches := bracketRe.FindStringSubmatch(input)

		if len(matches) == 0 {
			return mathEval(input)
		}

		value := mathEval(matches[1])
		replacer := strconv.Itoa(value)
		input = strings.Replace(input, matches[0], replacer, 1)
	}
}

func mathEvalInOrder(input string) int {
	operands := strings.Split(input, " ")
	operator := add
	result := 0

	for _, operand := range operands {
		if operand == add || operand == mul {
			operator = operand
			continue
		}

		value, _ := strconv.Atoi(operand)

		if operator == add {
			result += value
		} else if operator == mul {
			result *= value
		}

		operator = ""
	}

	return result
}

func mathEvalB(input string) int {
	operands := strings.Split(input, " ")
	operator := add
	result := 0

	for _, operand := range operands {
		if operand == add || operand == mul {
			operator = operand
			continue
		}

		value, _ := strconv.Atoi(operand)

		if operator == add {
			result += value
		} else if operator == mul {
			result *= value
		}

		operator = ""
	}

	return result
}
