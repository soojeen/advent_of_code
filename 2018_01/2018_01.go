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

	input, err := parseInput(rawInput)
	if err != nil {
		log.Fatal(err)
	}

	resultA := applyFrequencies(input)
	resultB := findTwiceFrequency(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) ([]int, error) {
	changes := strings.Split(input, "\n")

	new := make([]int, len(changes))
	var err error

	for i, changeStr := range changes {
		changeInt, e := strconv.Atoi(changeStr)
		if e != nil {
			err = e
			break
		}

		new[i] = changeInt
	}

	return new, err
}

func applyFrequencies(changes []int) int {
	frequency := 0
	for _, change := range changes {
		frequency += change
	}

	return frequency
}

func findTwiceFrequency(changes []int) int {
	frequencies := make(map[int]bool)
	frequency := 0
	outerBreak := false

	for {
		for _, change := range changes {
			frequency += change

			if frequencies[frequency] == true {
				outerBreak = true
				break
			}

			frequencies[frequency] = true
		}

		if outerBreak == true {
			break
		}
	}

	return frequency
}
