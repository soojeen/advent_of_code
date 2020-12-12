package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type ship struct {
	direction byte
	x         int
	y         int
}

func (s *ship) doAction(input nav) {

}

const north = 'N'
const south = 'S'
const east = 'E'
const west = 'W'
const left = 'L'
const right = 'R'

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input, parseError := parseInput(rawInput)
	if parseError != nil {
		log.Fatal(readError)
	}

	resultA := run(input)
	// resultB := run(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

type nav struct {
	action byte
	value  int
}

func parseInput(input string) ([]nav, error) {
	var err error

	lines := strings.Split(input, "\n")
	result := make([]nav, len(lines))

	for i, line := range lines {
		action := line[0]
		value, e := strconv.Atoi(line[1:])
		if e != nil {
			err = e
			break
		}

		result[i] = nav{action, value}
	}

	return result, err
}

func run(input []nav) int {
	ship := ship{east, 0, 0}

	for _, nav := range input {
		ship.doAction(nav)
		fmt.Println("a:", ship, nav)
	}

	return ship.x + ship.y
}

// func generateRotation(currentDirection rune) {
// 	rotation := map[rune]map[int]map[rune]rune{}

// 	rotation[north][90][left] = west
// 	rotation[north][180][left] = south
// 	rotation[north][270][left] = east

// 	rotation[north][90][right] = east
// 	rotation[north][180][right] = south
// 	rotation[north][270][right] = west

// 	rotation[]
// }
