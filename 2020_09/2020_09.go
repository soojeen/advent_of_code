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

	return -1, 0
}

func checkValid(numbers []int, value int) bool {
	input := make(map[int]bool)
	for _, number := range numbers {
		input[number] = true
	}

	for _, number := range numbers {
		diff := value - number

		if input[diff] {
			return true
		}
	}

	return false
}

func processB(numbers []int, value int, upperIndex int) int {
	set := findSet(numbers, value, upperIndex)
	sort.Ints(set)

	return set[0] + set[len(set)-1]
}

func findSet(numbers []int, value int, upperIndex int) []int {
	set := []int{numbers[upperIndex]}
	sum := numbers[upperIndex]
	next := upperIndex - 1
	isCollect := true

	for {
		if isCollect {
			number := numbers[next]

			if sum+number > value {
				isCollect = false
				continue
			}

			if (sum + number) == value {
				return set
			}

			set = append([]int{number}, set...)
			sum += number
			next--

		} else {
			isCollect = true
			sum -= set[len(set)-1]
			set = set[:len(set)-1]
		}

		if next == 0 {
			break
		}
	}

	return []int{}
}
