package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

type Claim struct {
	id string
	cornerX int
	cornerY int
	width int
	height int
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(strings.Split(rawInput, "\n"))

	fmt.Println("a:", input)

// #1 @ 596,731: 11x27

	// resultA := checksumGo(input)
	// resultB := findCommonLetters(input)
	// fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input []string) []Claim {
	return []Claim{Claim{"1", 1, 1, 1, 1}}
}