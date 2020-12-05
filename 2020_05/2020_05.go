package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := highestSeatID(input)
	// resultB := countValid(input, createValidatorB())

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n\n")
}

func highestSeatID(input []string) int {
	max := 0
	return max
}
