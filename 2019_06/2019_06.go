package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

const root = "COM"

type stringMap map[string]string

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parentMap := parseInput(rawInput)

	resultA := processA(parentMap)
	fmt.Println(resultA)

	resultB := processB(parentMap)
	fmt.Println(resultB)
}

func parseInput(rawInput string) stringMap {
	parentMap := make(map[string]string)
	orbitPairs := strings.Split(rawInput, "\n")

	for _, orbitPair := range orbitPairs {
		orbitPair := strings.Split(orbitPair, ")")

		parentMap[orbitPair[1]] = orbitPair[0]
	}

	return parentMap
}

func processA(parentMap stringMap) int {
	total := 0

	for _, parent := range parentMap {
		current := parent
		currentParentCount := 1

		for {
			if current == root {
				break
			}

			currentParentCount++
			current = parentMap[current]
		}

		total += currentParentCount
	}

	return total
}

type traceMap map[string]int

func traceMapForTarget(parentMap stringMap, target string) traceMap {
	traceMap := make(traceMap)
	current := parentMap[target]
	distance := 0

	for {
		if current == root {
			break
		}

		traceMap[current] = distance
		distance++
		current = parentMap[current]
	}

	return traceMap
}

func processB(parentMap stringMap) int {
	youTrace := traceMapForTarget(parentMap, "YOU")
	santaTrace := traceMapForTarget(parentMap, "SAN")

	closest := len(santaTrace) + len(parentMap)

	for key, value := range santaTrace {
		youValue := youTrace[key]
		total := value + youValue

		if youValue == 0 || total > closest {
			continue
		}

		closest = total
	}

	return closest
}
