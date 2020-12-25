package main

import "fmt"
import "log"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := sumHomework(input, false)
	resultB := sumHomework(input, true)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func sumHomework(input []string, advanced bool) int {
	result := 0

	for _, line := range input {
		lineValue := processLine(line, advanced)
		value, _ := strconv.Atoi(lineValue)
		result += value
	}

	return result
}

func processLine(input string, advanced bool) string {
	bracketRe := regexp.MustCompile(`\(([+*\d ]*)\)`)
	inOrderRe := regexp.MustCompile(`\d* [+*] \d*`)
	addRe := regexp.MustCompile(`\d* \+ \d*`)
	mulRe := regexp.MustCompile(`\d* \* \d*`)

	for {
		bracketMatches := bracketRe.FindStringSubmatch(input)

		if len(bracketMatches) > 0 {
			value := processLine(bracketMatches[1], advanced)
			input = strings.Replace(input, bracketMatches[0], value, 1)
			continue
		}

		inOrderMatches := inOrderRe.FindString(input)

		if !advanced && len(inOrderMatches) > 0 {
			value := mathEval(inOrderMatches)
			input = strings.Replace(input, inOrderMatches, value, 1)
			continue
		}

		addMatch := addRe.FindString(input)

		if advanced && len(addMatch) > 0 {
			value := mathEval(addMatch)
			input = strings.Replace(input, addMatch, value, 1)
			continue
		}

		mulMatch := mulRe.FindString(input)

		if advanced && len(mulMatch) > 0 {
			value := mathEval(mulMatch)
			input = strings.Replace(input, mulMatch, value, 1)
			continue
		}

		return input
	}
}

func mathEval(input string) string {
	operands := strings.Split(input, " ")
	operandA, _ := strconv.Atoi(operands[0])
	operandB, _ := strconv.Atoi(operands[2])

	if operands[1] == "+" {
		return strconv.Itoa(operandA + operandB)
	}

	return strconv.Itoa(operandA * operandB)
}
