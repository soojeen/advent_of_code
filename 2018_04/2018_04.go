package main

import "fmt"
import "log"

// import "strconv"
import "sort"
import "strings"
import "advent_of_code/utils"

type guardShift struct {
	date       string
	id         string
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

func mostSleepGuard(input []guardShift) int {
	fmt.Println("a:", input[0])

	return 1
}

func parseInput(input []string) []guardShift {
	sort.Strings(input)

	// currentGuard := ""
	// currentDate := ""
	for _, rawData := range input {
		data := strings.Split(rawData, "] ")
		rawDate, rawAction := data[0], data[1]
		action := strings.Split(rawAction, " ")

		switch action[0] {
		case "Guard":
			fmt.Println("g:", rawDate, action)
			// currentGuard = "new guard"
			// currentDate = "new date"
		case "falls":
			fmt.Println("f:", rawDate, action)
		case "wakes":
			fmt.Println("w:", rawDate, action)
		}
	}

	return []guardShift{guardShift{}}
}
