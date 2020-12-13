package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type point struct {
	x int
	y int
}
type ship struct {
	direction byte
	point     point
}

const north = 'N'
const south = 'S'
const east = 'E'
const west = 'W'
const left = 'L'
const right = 'R'
const forward = 'F'

func turn(input nav, direction byte) byte {
	directions := [4]byte{north, east, south, west}
	indexer := map[byte]int{north: 0, east: 1, south: 2, west: 3}
	shift := input.value / 90
	index := 0

	if input.action == left {
		index = indexer[direction] - shift
	} else {
		index = indexer[direction] + shift
	}

	if index < 0 {
		index += 4
	} else if index > 3 {
		index -= 4
	}

	return directions[index]
}

func (s *ship) doAction(input nav) {
	if input.action == left || input.action == right {

		s.direction = turn(input, s.direction)
		return
	}

	if input.action == forward {
		s.point = directionMove(s.point, nav{s.direction, input.value})
		return
	}

	s.point = directionMove(s.point, input)
}

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
	ship := ship{east, point{0, 0}}

	for _, nav := range input {
		ship.doAction(nav)
	}

	return absolute(ship.point.x) + absolute(ship.point.y)
}

func directionMove(input point, nav nav) point {
	result := point{input.x, input.y}

	switch nav.action {

	case north:
		result.y += nav.value
	case south:
		result.y -= nav.value
	case east:
		result.x += nav.value
	case west:
		result.x -= nav.value
	}

	return result
}

func absolute(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
