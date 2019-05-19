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

	resultA := maxArea(input)
	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

type step struct {
	prereq *step
	id     string
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
