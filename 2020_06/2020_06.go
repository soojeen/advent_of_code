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

	resultA := countAll(input)
	// resultB := findSeatID(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n\n")
}

func countAll(input []string) int {
	count := 0

	for _, group := range input {
		count += countGroup(group)
	}

	return count
}

func countGroup(input string) int {
	return 0
}
