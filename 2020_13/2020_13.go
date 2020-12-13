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
	resultB := processB(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
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
		diff := nextBusArrival(currentTime, bus)

		if timeToNext < 0 || diff < timeToNext {
			timeToNext = diff
			nextBus = bus
		}
	}

	return nextBus * timeToNext
}

func processB(input []int) int {
	buses := input[2:]
	baseDiff := input[1]
	currentDiff := input[1]

	for i, bus := range buses {
		if bus == 0 {
			continue
		}

		for j := baseDiff; ; j += currentDiff {
			nextBusDiff := nextBusArrival(j, bus)

			if nextBusDiff == (i%bus)+1 {
				baseDiff = j
				currentDiff *= bus
				break
			}
		}
	}

	return baseDiff
}

func nextBusArrival(currentTime int, bus int) int {
	next := (currentTime / bus * bus) + bus
	return next - currentTime
}
