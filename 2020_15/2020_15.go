package main

import "fmt"

func main() {
	input := []int{20, 9, 11, 0, 1, 2}

	resultA := processA(input, 2020)
	resultB := processA(input, 30000000)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func processA(input []int, end int) int {
	tracker := map[int]int{}
	prevDiff := 0
	result := 0

	for i := 1; i <= end; i++ {
		// initialize tracker with input
		if i < len(input)+1 {
			tracker[input[i-1]] = i
			continue
		}

		// prevDiff defaults to 0 if prev is first
		value := prevDiff
		if tracker[value] == 0 {
			prevDiff = 0
		} else {
			prevDiff = i - tracker[value]
		}

		tracker[value] = i
		result = value
	}

	return result
}
