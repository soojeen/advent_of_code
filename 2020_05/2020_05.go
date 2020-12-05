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

	input := parseInput(rawInput)

	resultA := highestSeatID(input)
	resultB := findSeatID(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func highestSeatID(input []string) int {
	max := 0

	for _, seat := range input {
		seatID := getSeatID(seat)

		if seatID > max {
			max = seatID
		}
	}

	return max
}

func findSeatID(input []string) int {
	trackingMap := make(map[int]int)

	for _, seat := range input {
		seatID := getSeatID(seat)
		trackingMap[seatID] = seatID
	}

	for key := range trackingMap {
		isSeat := trackingMap[key+1] == 0 && trackingMap[key+2] != 0

		if isSeat {
			return key + 1
		}
	}

	return 0
}

func getSeatID(input string) int {
	binary := ""

	for _, char := range input {
		binary += converter(char)
	}

	seatID, parseError := strconv.ParseInt(binary, 2, 0)
	if parseError != nil {
		log.Fatal(parseError)
	}

	return int(seatID)
}

func converter(input rune) string {
	switch input {
	case 'F':
		return "0"
	case 'B':
		return "1"
	case 'L':
		return "0"
	case 'R':
		return "1"
	}

	return ""
}
