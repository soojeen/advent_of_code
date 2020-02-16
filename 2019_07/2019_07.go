package main

import "fmt"
import "log"
import "advent_of_code/intcomp"
import "advent_of_code/permutations"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	program := intcomp.Parse(rawInput)

	resultA := processA(program)
	fmt.Println(resultA)

	resultB := processB(program)
	fmt.Println(resultB)
}

func processA(program []int) int {
	base := []int{0, 1, 2, 3, 4}
	maxOutput := 0

	permutations.IteratePermutation(base, func(phaseSettings []int) {
		input := 0

		for _, phaseSetting := range phaseSettings {
			programCopy := make([]int, len(program))
			copy(programCopy, program)
			inputs := []int{phaseSetting, input}

			computer := intcomp.Computer{
				Program: programCopy,
				Inputs:  inputs,
			}

			newInput := computer.RunProgram()
			input = newInput
		}

		if input > maxOutput {
			maxOutput = input
		}
	})

	return maxOutput
}

func processB(program []int) int {
	return 2
}
