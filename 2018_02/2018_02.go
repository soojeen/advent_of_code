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
	resultB := findCommonLetters(input)
	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
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

func findCommonLetters(input []string) string {
	var commonChars string
	var isMatch bool

	for idx, id1 := range input {
		restOfInput := input[(idx + 1):]
		for _, id2 := range restOfInput {
			isMatch, commonChars = compare(id1, id2)
			if isMatch {
				break
			}
		}

		if isMatch {
			break
		}
	}

	return commonChars
}

func compare(id1 string, id2 string) (bool, string) {
	offChars := 0
	offCharIndex := 0
	for i := range id1 {
		if id1[i] != id2[i] {
			offChars++
			offCharIndex = i
		}
	}

	var commonChars string
	if offChars == 1 {
		commonChars = id1[:offCharIndex] + id1[(offCharIndex+1):]
		return true, commonChars
	}

	return false, ""
}

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

	for _, char := range str {
		charS := string(char)
		if charCounts[charS] == 0 {
			charCounts[charS] = 1
		} else {
			charCounts[charS]++
		}
	}

	return charCounts
}
