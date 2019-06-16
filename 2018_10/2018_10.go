package main

import "fmt"
import "log"

// import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(rawInput)

	// resultA := winningScore(input, 1)
	// resultB := winningScore(input, 100)

	fmt.Println("a:", input)
	// fmt.Println("b:", resultB)
}

type light struct {
	position [2]int
	velocity [2]int
}

func parseInput(rawInput string) []light {
	lines := strings.Split(rawInput, "\n")
	lights := make([]light, len(lines))

	for _, line := range lines {
		fmt.Println("line:", line)

	}

	return lights
}
