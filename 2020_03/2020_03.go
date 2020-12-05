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

	resultA := countTrees(input, 3, 1)
	resultB := moreTrails(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func countTrees(input []string, right int, down int) int {
	count := 0
	marker := 0
	moduloOperator := len(input[0])

	for i := 0; i < len(input); i += down {
		relativeMarker := marker % moduloOperator
		value := string(input[i][relativeMarker])

		if value == "#" {
			count++
		}

		marker += right
	}

	return count
}

func moreTrails(input []string) int {
	a := countTrees(input, 1, 1)
	b := countTrees(input, 3, 1)
	c := countTrees(input, 5, 1)
	d := countTrees(input, 7, 1)
	e := countTrees(input, 1, 2)

	return a * b * c * d * e
}
