package main

import "fmt"

type prev struct {
	isFirst bool
	age     int
}

func main() {
	input := []int{20, 9, 11, 0, 1, 2}

	resultA := processA(input)
	// resultB := processB(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func processA(input []int) int {
	tracker := map[int]int{}
	turns := make([]int, 2020)
	last := prev{false, 0}

	for i := 1; i <= len(turns); i++ {
		if i < len(input)+1 {
			value := input[i-1]
			tracker[value] = i
			turns[i-1] = value
			last = prev{true, 0}
			continue
		}

		if last.isFirst {
			value := 0
			last = prev{false, i - tracker[value]}

			tracker[value] = i
			turns[i-1] = value
			continue
		}

		value := last.age
		if tracker[value] == 0 {
			last = prev{true, 0}
		} else {
			last = prev{false, i - tracker[value]}
		}
		tracker[value] = i
		turns[i-1] = value
	}

	return turns[2020-1]
}
