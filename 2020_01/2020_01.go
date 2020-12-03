package main

import "fmt"
import "log"
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

	resultA := findTwo(input)
	resultB := findThree(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) ([]int, error) {
	expenses := strings.Split(input, "\n")

	result := make([]int, len(expenses))
	var err error

	for i, rawExpense := range expenses {
		expense, e := strconv.Atoi(rawExpense)
		if e != nil {
			err = e
			break
		}

		result[i] = expense
	}

	return result, err
}

func findTwo(input []int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}

			if input[i]+input[j] == 2020 {
				result = input[i] * input[j]
				break
			}
		}
		if result > 0 {
			break
		}
	}

	return result
}

func findThree(input []int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}

			for k := 0; k < len(input); k++ {
				if (i == k) || (j == k) {
					continue
				}

				if input[i]+input[j]+input[k] == 2020 {
					result = input[i] * input[j] * input[k]
					break
				}
			}

			if result > 0 {
				break
			}
		}

		if result > 0 {
			break
		}
	}

	return result
}
