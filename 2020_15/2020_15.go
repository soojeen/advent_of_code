package main

import "fmt"

type meta struct {
	prevTurn int
}

func main() {
	input := []int{20, 9, 11, 0, 1, 2}

	resultA := processA(input)
	// resultB := processB(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func processA(input []int) int {
	tracker := map[int]meta{}
	turns := make([]int, 2020)
	prevDiff := 0
	prev := 0
	prevIsFirst := true

	for i := 1; i <= 2020; i++ {
		if i < len(input)+1 {
			value := input[i-1]
			tracker[value] = meta{i}
			turns[i-1] = value
			prev = value
			continue
		}

		if prevIsFirst {
			prev = 0
			prevDiff = i - tracker[0].prevTurn
			prevIsFirst = false
			tracker[0] = meta{i}
			turns[i-1] = 0
			continue
		}

		if i > 20 {
			break
		}
	}
	fmt.Println("a:", tracker, turns, prev, prevDiff)

	return 0
}
