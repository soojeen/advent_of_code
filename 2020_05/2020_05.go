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
	rawRow := input[:7]
	rawCol := input[7:]

	rowBin := binaryPartition(rawRow, 'B', 'F')
	colBin := binaryPartition(rawCol, 'R', 'L')

	row, parseError := strconv.ParseInt(rowBin, 2, 0)
	if parseError != nil {
		log.Fatal(parseError)
	}

	col, parseError := strconv.ParseInt(colBin, 2, 0)
	if parseError != nil {
		log.Fatal(parseError)
	}

	return (int(row) * 8) + int(col)
}

func binaryPartition(input string, top rune, bottom rune) string {
	binary := ""

	for _, char := range input {
		switch char {
		case top:
			binary += "1"
		case bottom:
			binary += "0"
		}
	}

	return binary
}
