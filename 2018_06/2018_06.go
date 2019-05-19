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

	input := parsePoints(rawInput)

	resultA, resultB := maxArea(input)
	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

type point struct {
	x int
	y int
}

func parsePoints(input string) []point {
	rawPoints := strings.Split(input, "\n")
	points := make([]point, len(rawPoints))

	for i, rawPoint := range rawPoints {
		currentPoint := point{}

		_, error := fmt.Sscanf(rawPoint, "%d, %d", &currentPoint.x, &currentPoint.y)
		if error != nil {
			panic(error)
		}

		points[i] = currentPoint
	}

	return points
}

func maxArea(pointsInput []point) (int, int) {
	minPoint := pointsInput[0]
	maxPoint := pointsInput[0]

	for _, point := range pointsInput {
		if point.x < minPoint.x {
			minPoint.x = point.x
		}
		if point.x > maxPoint.x {
			maxPoint.x = point.x
		}
		if point.y < minPoint.y {
			minPoint.y = point.y
		}
		if point.y > maxPoint.y {
			maxPoint.y = point.y
		}
	}

	infinitePoints := make(map[point]bool)
	pointAreas := make(map[point]int)
	areaB := 0

	for x := minPoint.x; x <= maxPoint.x; x++ {
		for y := minPoint.y; y <= maxPoint.y; y++ {
			nearestDistance := -1
			nearestInputPoint := point{-1, -1}
			sumDistanceToInputs := 0

			for _, pointInput := range pointsInput {
				distance := absolute(x-pointInput.x) + absolute(y-pointInput.y)

				if distance < nearestDistance || nearestDistance == -1 {
					nearestDistance = distance
					nearestInputPoint = pointInput
				} else if distance == nearestDistance {
					nearestInputPoint = point{-1, -1}
				}

				sumDistanceToInputs += distance
			}

			if x == minPoint.x || y == minPoint.y || x == maxPoint.x || y == maxPoint.y {
				infinitePoints[nearestInputPoint] = true
			}

			pointAreas[nearestInputPoint]++

			if sumDistanceToInputs < 10000 {
				areaB++
			}
		}
	}

	maxArea := 0
	for point, area := range pointAreas {
		if !infinitePoints[point] && area > maxArea {
			maxArea = area
		}
	}

	return maxArea, areaB
}

func absolute(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
