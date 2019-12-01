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

	input := parseInput(rawInput)
	fmt.Println(input)

	resultA := sumFuelRequirementsA(input)
	resultB := sumFuelRequirementsB(input)
	fmt.Println(resultA)
	fmt.Println(resultB)
}

func parseInput(rawInput string) []int {
	massInputs := strings.Split(rawInput, "\n")

	masses := make([]int, len(massInputs))
	for i := range massInputs {
		massInput, e := strconv.Atoi(massInputs[i])
		if e != nil {
			log.Fatal(e)
		}

		masses[i] = massInput
	}

	return masses
}

func sumFuelRequirementsA(input []int) int {
	sum := 0
	for i := range input {
		fuelRequirement := (input[i] / 3) - 2
		sum += fuelRequirement
	}

	return sum
}

func recursiveFuelRequirement(input int) int {
	fuelRequirment := (input / 3) - 2

	if fuelRequirment <= 0 {
		return 0
	}

	return fuelRequirment + recursiveFuelRequirement(fuelRequirment)
}

func sumFuelRequirementsB(input []int) int {
	sum := 0
	for i := range input {
		fuelRequirement := recursiveFuelRequirement(input[i])
		sum += fuelRequirement
	}

	return sum
}
