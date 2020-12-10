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
	current := 0

	for _, number := range input {
		diff := number - current
		diffs[diff]++
		current += diff
	}

	// extra 3 diff between last adaptor and device
	return diffs[1] * (diffs[3] + 1)
}

func findArrangements(input []int) int {
	multiplier := map[int]int{1: 1, 2: 2, 3: 4, 4: 7}

	// insert 0 plug source
	numbers := append(input, 0)
	sort.Ints(numbers)
	last := numbers[len(numbers)-1]
	// padding to help check next diffs
	numbers = append(numbers, last+3, last+6)

	result := 1
	diffCount := 0

	for i := 0; i < len(input); i++ {
		number := numbers[i]
		next := numbers[i+1]
		diff := next - number
		nextDiff := numbers[i+2] - next

		if diff == 3 {
			continue
		}

		if diffCount == 0 && diff != nextDiff {
			continue
		}

		diffCount++

		if diff != nextDiff {
			result *= multiplier[diffCount]
			diffCount = 0
		}

	}

	return result
}
