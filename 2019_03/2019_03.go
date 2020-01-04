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

	pathA, pathB := parseInput(rawInput)

	resultA := process(pathA, pathB)

	fmt.Println(resultA)
}

func parseInput(rawInput string) ([]string, []string) {
	rawWirePaths := strings.Split(rawInput, "\n")
	wirePaths := make([][]string, 2)

	for i, rawWirePath := range rawWirePaths {
		wirePaths[i] = strings.Split(rawWirePath, ",")
	}

	return wirePaths[0], wirePaths[1]
}

type point struct {
	x, y int
}

type pointContent struct {
	a, b bool
}

type gridState map[point]map[rune]bool

func process(pathA []string, pathB []string) int {
	gridState := make(gridState)

	gridState = tracePath(pathA, 'a', gridState)
	gridState = tracePath(pathB, 'b', gridState)

	return findResult(gridState)
}

func tracePath(path []string, gridKey rune, gridState gridState) gridState {
	current := point{0, 0}

	for _, step := range path {
		direction, distance := parseStep(step)
		for i := 0; i < distance; i++ {
			switch direction {
			case "U":
				current.y++
			case "D":
				current.y--
			case "R":
				current.x++
			case "L":
				current.x--
			default:
				fmt.Println("No Direction")
			}

			currentPoint := gridState[current]
			if len(currentPoint) == 0 {
				currentPoint = make(map[rune]bool)
			}
			currentPoint[gridKey] = true
			gridState[current] = currentPoint
		}
	}

	return gridState
}

func parseStep(rawStep string) (string, int) {
	step := strings.SplitN(rawStep, "", 2)
	distance, e := strconv.Atoi(step[1])
	if e != nil {
		log.Fatal(e)
	}

	return step[0], distance
}

func findResult(gridState gridState) int {
	intersections := make([]point, 0)

	for key, value := range gridState {
		if value['a'] && value['b'] {
			intersections = append(intersections, key)
		}
	}

	closestDistance := -1

	for _, intersection := range intersections {
		distance := intersection.x + intersection.y
		if closestDistance == -1 || distance < closestDistance {
			closestDistance = distance
		}
	}

	return closestDistance
}
