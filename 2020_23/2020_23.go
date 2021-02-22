package main

import "fmt"

func main() {
	// input := [9]int{3, 1, 5, 6, 7, 9, 8, 2, 4}
	// test
	input := [9]int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	resultA := partA(input)

	fmt.Println("a:", resultA)
}

func partA(input [9]int) [9]int {
	currentIndex := 0

	for i := 0; i < 10; i++ {
		currentValue := input[currentIndex]

		picks := input[currentIndex+1 : currentIndex+4]
		destination := getDestination(input, picks)
		fmt.Println("a:", picks, currentValue)
	}

	return input
}

func getDestination(input [9]int, picks []int, currentValue int) (int, int) {
	destinationIndex := 0
	destinationValue := currentValue

	for i := 0; i < 10; i++ {
	}

	return destinationIndex, destinationValue
}
