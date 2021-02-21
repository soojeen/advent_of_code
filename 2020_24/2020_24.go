package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

type axial [2]int

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := flipTiles(input)
	// resultB := sumHomework(input, true)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func flipTiles(input []string) int {
	axials := make([]axial, len(input))

	for i, line := range input {
		axials[i] = findTile(line)
	}

	result := countTileFlips(axials)
	return result
}

func findTile(input string) axial {
	axial := axial{0, 0}
	var prev rune

	for _, char := range input {
		switch char {
		case 'e':
			if prev != 's' {
				axial[0]--
			}
		case 'w':
			if prev != 'n' {
				axial[0]++
			}
		case 's':
			axial[1]--
		case 'n':
			axial[1]++
		}

		prev = char
	}

	return axial
}

func countTileFlips(input []axial) int {
	counts := make(map[axial]int)

	for _, axial := range input {
		counts[axial]++
	}

	flipCount := 0

	for _, count := range counts {
		if count%2 == 1 {
			flipCount++
		}
	}

	return flipCount
}
