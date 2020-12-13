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
		next := (currentTime / bus * bus) + bus
		diff := next - currentTime

		if timeToNext < 0 || diff < timeToNext {
			timeToNext = diff
			nextBus = bus
		}
	}

	return nextBus * timeToNext
}

func processB(input []int) int {
	buses := input[2:]
	earliest := 0

	for i := 0; ; i++ {
		firstBus := i * input[1]

		isAllValid := true
		for j, bus := range buses {
			if bus == 0 {
				continue
			}

			nextBusDiff := nextBusArrival(firstBus, bus)
			// fmt.Println("b:", bus, nextBusDiff, j)

			if nextBusDiff != j+1 {
				isAllValid = false
				break
			}
		}

		if isAllValid {
			earliest = firstBus
			break
		}

		// fmt.Println("b:", firstBus)
		// if i > 100 {
		// 	break
		// }

	}

	// for each departure time of first bus
	// 	for each subsequent bus
	// 		check diff against index constraint

	return earliest
}

func nextBusArrival(currentTime int, bus int) int {
	next := (currentTime / bus * bus) + bus
	return next - currentTime
}

// 299 high

// 1068781
// 1068781
