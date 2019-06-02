package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(rawInput)

	// resultA := correctOrder(reqs)
	// resultB := multiWorkers(reqs)
	fmt.Println("a:", input)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) []int {
	values := strings.Split(input, " ")
	result := make([]int, len(values))

	for i, value := range values {
		iValue, _ := strconv.Atoi(value)
		result[i] = iValue
	}

	return result
}
