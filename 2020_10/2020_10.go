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
	resultB := findArrangements(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
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
	prev := 0

	for _, number := range input {
		diff := number - prev
		diffs[diff]++
		prev = number
	}

	// extra 3 diff between last adaptor and device
	return diffs[1] * (diffs[3] + 1)
}

func findArrangements(input []int) int {
	multiplier := map[int]int{1: 1, 2: 2, 3: 4, 4: 7}

	sort.Ints(input)

	prevDiff := 0
	prevNumber := 0
	result := 1
	sameDiffCount := 1

	for i, number := range input {
		diff := number - prevNumber

		if prevDiff == 1 && diff == 1 {
			sameDiffCount++
		}

		if prevDiff == 1 && diff == 3 {
			result *= multiplier[sameDiffCount]
			sameDiffCount = 1
		}

		// final element
		if i == len(input)-1 {
			result *= multiplier[sameDiffCount]
		}

		prevDiff = diff
		prevNumber = number
	}

	return result
}

// 86812553324672
// 86812553324672
