package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(rawInput, "\n")

	resultA := checksum(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func checksum(input []string) int {
	// TODO: swap to concurrent go routines and use channels
	twos := 0
	threes := 0
	for _, id := range input {
		charCounts := uniqueCharCount(id)

		if charCountHasCount(charCounts, 2) {
			twos += 1
		}

		if charCountHasCount(charCounts, 3) {
			threes += 1
		}
	}

	return twos * threes
}

func charCountHasCount(charCounts map[string]int, count int) bool {
	for _, charCount := range charCounts {
		if count == charCount {
			return true
		}
	}

	return false
}

func uniqueCharCount(str string) map[string]int {
	charCounts := make(map[string]int)

	for _, rune := range str {
		char := string(rune)
		if charCounts[char] == 0 {
			charCounts[char] = 1
		} else {
			charCounts[char] += 1
		}
	}

	return charCounts
}
