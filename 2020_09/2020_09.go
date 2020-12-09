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

	resultA, upperIndex := findInvalid(input)
	resultB := processB(input, resultA, upperIndex)

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

func findInvalid(input []int) (int, int) {
	for i := 25; i < len(input); i++ {
		number := input[i]
		isValid := checkValid(input[i-25:i], number)

		if !isValid {
			return number, i - 25
		}
	}

	return 0, 0
}

func checkValid(input []int, value int) bool {
	numbers := make(map[int]bool)
	for _, number := range input {
		numbers[number] = true
	}

	for _, number := range input {
		if numbers[value-number] {
			return true
		}
	}

	return false
}

func processB(numbers []int, value int, upperIndex int) int {
	group := findGroup(numbers, value, upperIndex)
	sort.Ints(group)

	return group[0] + group[len(group)-1]
}

func findGroup(numbers []int, value int, upperIndex int) []int {
	group := []int{numbers[upperIndex]}
	sum := numbers[upperIndex]

	for next := upperIndex - 1; next >= 0; {
		number := numbers[next]

		if sum+number > value {
			lastIndex := len(group) - 1
			sum -= group[lastIndex]
			group = group[:lastIndex]
			continue
		}

		group = append([]int{number}, group...)
		sum += number
		next--

		if sum == value {
			return group
		}
	}

	return []int{}
}
