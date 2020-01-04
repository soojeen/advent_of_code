package main

import "fmt"
import "strconv"
import "strings"

func main() {
	input := []int{128392, 643281}

	resultA := processA(input)
	resultB := processB(resultA)

	fmt.Println(len(resultA))
	fmt.Println(len(resultB))
}

func processA(input []int) []int {
	valid := make([]int, 0)

	for i := input[0]; i <= input[1]; i++ {
		if isValid(i) {
			valid = append(valid, i)
		}
	}

	return valid
}

func processB(input []int) []int {
	valid := make([]int, 0)

	for _, password := range input {
		if hasValidAdjacent(password) {
			valid = append(valid, password)
		}
	}

	return valid
}

func isValid(password int) bool {
	passwordString := strconv.Itoa(password)
	passwordLetters := strings.Split(passwordString, "")

	neverDecreases := true
	hasAdjacent := false

	for i, letter := range passwordLetters {
		if i == len(passwordLetters)-1 {
			continue
		}

		if passwordLetters[i+1] < letter {
			neverDecreases = false
			break
		}

		if passwordLetters[i+1] == letter {
			hasAdjacent = true
		}
	}

	return neverDecreases && hasAdjacent
}

func hasValidAdjacent(password int) bool {
	passwordString := strconv.Itoa(password)
	passwordLetters := strings.Split(passwordString, "")

	hasValidRepeat := false
	repeatCount := 1

	for i, letter := range passwordLetters {
		eol := i == len(passwordLetters)-1

		if eol {
			hasValidRepeat = repeatCount == 2
			break
		}

		nextLetterRepeats := letter == passwordLetters[i+1]

		if nextLetterRepeats {
			repeatCount++
			continue
		}

		if !nextLetterRepeats && repeatCount == 1 {
			continue
		}

		if repeatCount == 2 {
			hasValidRepeat = true
			break
		}

		if repeatCount != 2 {
			repeatCount = 1
		}

	}

	return hasValidRepeat
}
