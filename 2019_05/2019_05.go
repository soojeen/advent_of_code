package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"
import "advent_of_code/intcomp"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	program := parseInput(rawInput)

	resultA := process(program, 1)
	fmt.Println(resultA)

	resultB := process(program, 5)
	fmt.Println(resultB)
}

func parseInput(rawInput string) []int {
	_positions := strings.Split(rawInput, ",")

	positions := make([]int, len(_positions))
	for i := range positions {
		position, e := strconv.Atoi(_positions[i])
		if e != nil {
			log.Fatal(e)
		}

		positions[i] = position
	}

	return positions
}

func process(program []int, input int) int {
	programCopy := append(program[:0:0], program...)
	computer := intcomp.Computer{
		Program: programCopy,
		Input:   input,
	}

	return computer.RunProgram()
}
