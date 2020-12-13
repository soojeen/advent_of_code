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

	resultA := processA(input)
	// resultB := findSeatID(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) ([]int, error) {
	lines := strings.Split(input, "\n")
	buses := strings.Split(lines[1], ",")
	result := make([]int, len(buses)+1)

	timestamp, e := strconv.Atoi(lines[0])
	if e != nil {
		return result, e
	}
	result[0] = timestamp

	var err error
	for i, bus := range buses {
		if bus == "x" {
			continue
		}

		busNumber, e := strconv.Atoi(bus)

		if e != nil {
			err = e
			break
		}

		result[i+1] = busNumber
	}

	return result, err
}

func processA(input []int) int {
	currentTime := input[0]
	buses := input[1:]
	nextBus := 0
	timeToNext := -1

	for _, bus := range buses {
		if bus == 0 {
			continue
		}

		// assume no bus leaving exactly at curentTime
		next := (currentTime / bus * bus) + bus
		diff := next - currentTime

		if timeToNext < 0 || diff < timeToNext {
			timeToNext = diff
			nextBus = bus
		}
	}
	fmt.Println("a:", nextBus, timeToNext)

	return nextBus * timeToNext
}

// 299 high
