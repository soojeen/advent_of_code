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

	resultA := sumHomework(input)
	resultB := sumHomework(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func sumHomework(input []string) int {
	result := 0

	for _, line := range input {
		lineValue := processLine(line)
		value, _ := strconv.Atoi(lineValue)
		result += value
	}

	return result
}

func processLine(input string) string {
	bracketRe := regexp.MustCompile(`\(([+*\d ]*)\)`)
	inOrderRe := regexp.MustCompile(`\d* [+*] \d*`)

	for {
		bracketMatches := bracketRe.FindStringSubmatch(input)

		if len(bracketMatches) > 0 {
			value := processLine(bracketMatches[1])
			input = strings.Replace(input, bracketMatches[0], value, 1)
			continue
		}

		inOrderMatches := inOrderRe.FindString(input)

		if len(inOrderMatches) > 0 {
			value := mathEval(inOrderMatches)
			input = strings.Replace(input, inOrderMatches, value, 1)
			continue
		}

		return input
	}
}

func mathEval(input string) string {
	operands := strings.Split(input, " ")
	operandA, _ := strconv.Atoi(operands[0])
	operandB, _ := strconv.Atoi(operands[2])

	if operands[1] == add {
		return strconv.Itoa(operandA + operandB)
	}

	return strconv.Itoa(operandA * operandB)
}
