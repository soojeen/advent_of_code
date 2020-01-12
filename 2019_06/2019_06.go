package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	orbits := parseInput(rawInput)

	resultA := process(orbits)
	fmt.Println(resultA)
}

func parseInput(rawInput string) [][]string {
	orbitPairs := strings.Split(rawInput, "\n")

	orbits := make([][]string, len(orbitPairs))
	for i := range orbits {
		orbitPair := strings.Split(orbitPairs[i], ")")

		orbits[i] = orbitPair
	}

	return orbits
}

func process(orbits [][]string) int {
	return 1
}
