package main

import "fmt"
import "log"
import "sort"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input, parseError := parseInput(rawInput)
	if parseError != nil {
		log.Fatal(parseError)
	}

	resultA := findDistribution(input)
	// resultB := processB(input, resultA, upperIndex)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) ([]int, error) {
	var err error
	lines := strings.Split(input, "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		number, error := strconv.Atoi(line)
		if error != nil {
			err = error
			break
		}

		numbers[i] = number
	}

	return numbers, err
}

func findDistribution(input []int) int {
	sort.Ints(input)

	diffs := make(map[int]int)
	current := 0

	for _, number := range input {
		diff := number - current
		diffs[diff]++
		current += diff
	}
	fmt.Println("a:", diffs)

	// extra 3 diff between last adaptor and device
	return diffs[1] * (diffs[3] + 1)
}
