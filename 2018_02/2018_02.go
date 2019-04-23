package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(rawInput)

	// resultA := applyFrequencies(input)
	// resultB := findTwiceFrequency(input)

	// fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}
