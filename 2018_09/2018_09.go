package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type treePartial struct {
	metaDataSum int
	value       int
	subTree     []int
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(rawInput)
	// result := parseTreeSum(input)

	fmt.Println("a:", input)
	// fmt.Println("b:", result.value)
}

type input struct {
	players int
	points  int
}

func parseInput(rawInput string) input {
	values := strings.Split(rawInput, " ")
	players, _ := strconv.Atoi(values[0])
	points, _ := strconv.Atoi(values[6])

	return input{players, points}
}
