package main

import "fmt"
import "log"

// import "strconv"
import "sort"
import "strings"
import "advent_of_code/utils"

type GuardShift struct {
	date string
	id string
	isSleeping [60]bool
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(strings.Split(rawInput, "\n"))

	resultA := mostSleepGuard(input)
	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func mostSleepGuard (input []GuardShift) int {
	fmt.Println("a:", input[0])

	return 1
}

func parseInput (input []string) []GuardShift {
	sort.Strings(input)

	
	return []GuardShift{GuardShift{}}
}
