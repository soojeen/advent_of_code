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

	wirePaths := parseInput(rawInput)

	gridState := buildGrid(wirePaths)
	resultA, resultB := process(gridState)

	fmt.Println(resultA)
	fmt.Println(resultB)
}

func parseInput(rawInput string) [][]string {
	rawWirePaths := strings.Split(rawInput, "\n")
	wirePaths := make([][]string, 2)

	for i, rawWirePath := range rawWirePaths {
		wirePaths[i] = strings.Split(rawWirePath, ",")
	}

	return wirePaths
}

type point struct {
	x, y int
}

type pointContent [2]int

type gridState map[point]pointContent

func buildGrid(paths [][]string) gridState {
	gridState := make(gridState)

	for i, path := range paths {
		gridState = tracePath(path, i, gridState)
	}

	return gridState
}

func tracePath(path []string, pointContentKey int, gridState gridState) gridState {
	current := point{0, 0}
	stepCount := 0

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

			stepCount++
			currentPoint := gridState[current]
			// ignore multiple contacts and keep lowest step count
			if currentPoint[pointContentKey] == 0 {
				currentPoint[pointContentKey] = stepCount
			}
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

func process(gridState gridState) (int, int) {
	intersectionPoints := make([]point, 0)

	for point, pointContent := range gridState {
		if pointContent[0] > 0 && pointContent[1] > 0 {
			intersectionPoints = append(intersectionPoints, point)
		}
	}

	var closestManhattan int
	var closestSteps int

	for i, point := range intersectionPoints {
		manhattan := point.x + point.y
		if i == 0 || manhattan < closestManhattan {
			closestManhattan = manhattan
		}

		pointContent := gridState[point]
		steps := pointContent[0] + pointContent[1]
		if i == 0 || steps < closestSteps {
			closestSteps = steps
		}
	}

	return closestManhattan, closestSteps
}
