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

	resultA := countAll(input, countGroupA)
	resultB := countAll(input, countGroupB)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n\n")
}

func countAll(input []string, counter func(string) int) int {
	count := 0

	for _, group := range input {
		count += counter(group)
	}

	return count
}

func countGroupA(input string) int {
	rawResult := strings.Replace(input, "\n", "", -1)

	charTracker := make(map[rune]bool)
	for _, char := range rawResult {
		charTracker[char] = true
	}

	return len(charTracker)
}

func countGroupB(input string) int {
	personCount := strings.Count(input, "\n") + 1
	rawResult := strings.Replace(input, "\n", "", -1)

	charTracker := make(map[rune]int)
	count := 0

	for _, char := range rawResult {
		charTracker[char]++

		if charTracker[char] == personCount {
			count++
		}
	}

	return count
}
