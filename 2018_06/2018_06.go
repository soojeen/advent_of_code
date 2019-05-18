package main

import "fmt"

import "log"

// import "regexp"
// import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	resultA := funcA(rawInput)
	// resultB := shortestPolymer(rawInput)
	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func funcA(input string) int {
	return 1
}
