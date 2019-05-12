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

	resultA := checksumGo(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

// func checksum(input []string) int {
// 	// TODO: swap to concurrent go routines and use channels
// 	twos := 0
// 	threes := 0
// 	for _, id := range input {
// 		charCounts := uniqueCharCount(id)

// 		if charCountHasCount(charCounts, 2) {
// 			twos += 1
// 		}

// 		if charCountHasCount(charCounts, 3) {
// 			threes += 1
// 		}
// 	}

// 	return twos * threes
// }

func checksumGo(input []string) int {
	type TwoThree struct {
		two   bool
		three bool
	}

	inputCount := len(input)
	c := make(chan TwoThree, inputCount)

	for _, id := range input {
		id := id
		go func() {
			charCounts := uniqueCharCount(id)
			two := charCountHasCount(charCounts, 2)
			three := charCountHasCount(charCounts, 3)

			c <- TwoThree{two, three}
		}()
	}

	twos := 0
	threes := 0
	for i := 0; i < inputCount; i++ {
		twoThree := <-c
		twos += boolToCount(twoThree.two)
		threes += boolToCount(twoThree.three)
	}

	return twos * threes
}

func boolToCount(booly bool) int {
	if booly {
		return 1
	}
	return 0
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
