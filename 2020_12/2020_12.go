package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type nav struct {
	action byte
	value  int
}
type point struct {
	x int
	y int
}
type ship struct {
	direction point
	location  point
	waypoint  point
}

const north = 'N'
const south = 'S'
const east = 'E'
const west = 'W'
const left = 'L'
const right = 'R'
const forward = 'F'

func (s *ship) doActionA(input nav) {
	if input.action == left || input.action == right {
		s.direction = turnVector(s.direction, input)
		return
	}

	if input.action == forward {
		vector := mulitplyVector(s.direction, input.value)
		s.location = addVector(s.location, vector)
		return
	}

	vector := getDirectionVector(input.action)
	vector = mulitplyVector(vector, input.value)
	s.location = addVector(s.location, vector)
}

// func (s *ship) doActionB(input nav) {
// 	if input.action == left || input.action == right {
// 		s.waypoint = turnWaypoint(s.point, s.waypoint, input)
// 		return
// 	}

// 	if input.action == forward {
// 		s.point, s.waypoint = moveToWayPoint(s.point, s.waypoint, input.value)
// 		return
// 	}

// 	s.waypoint = directionMoveLegacy(s.waypoint, input)
// }

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input, parseError := parseInput(rawInput)
	if parseError != nil {
		log.Fatal(readError)
	}

	resultA := runA(input)
	// resultB := runB(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
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

func runA(input []nav) int {
	ship := ship{getDirectionVector('E'), point{0, 0}, point{0, 0}}

	for _, nav := range input {
		ship.doActionA(nav)
	}

	return absolute(ship.location.x) + absolute(ship.location.y)
}

// func runB(input []nav) int {
// 	ship := ship{east, point{0, 0}, point{10, 1}}

// 	for _, nav := range input {
// 		ship.doActionB(nav)
// 	}

// 	fmt.Println("vector:", turnVector90L(point{1, 0}))
// 	fmt.Println("vector:", turnVector90L(point{0, 1}))
// 	fmt.Println("vector:", turnVector90L(point{-1, 0}))
// 	fmt.Println("vector:", turnVector90L(point{0, -1}))

// 	return absolute(ship.point.x) + absolute(ship.point.y)
// }

// func directionMoveLegacy(input point, nav nav) point {

// 	result := point{input.x, input.y}

// 	switch nav.action {

// 	case north:
// 		result.y += nav.value
// 	case south:
// 		result.y -= nav.value
// 	case east:
// 		result.x += nav.value
// 	case west:
// 		result.x -= nav.value
// 	}

// 	return result
// }

// func turn(input nav, direction byte) byte {
// 	directions := [4]byte{north, east, south, west}
// 	indexer := map[byte]int{north: 0, east: 1, south: 2, west: 3}
// 	shift := input.value / 90
// 	index := 0

// 	if input.action == left {
// 		index = indexer[direction] - shift
// 	} else {
// 		index = indexer[direction] + shift
// 	}

// 	if index < 0 {
// 		index += 4
// 	} else if index > 3 {
// 		index -= 4
// 	}

// 	return directions[index]
// }

// func moveToWayPoint(shipPoint point, waypoint point, value int) (point, point) {
// 	xDiff := value * (waypoint.x - shipPoint.x)
// 	yDiff := value * (waypoint.y - shipPoint.y)

// 	newShipPoint := point{shipPoint.x + xDiff, shipPoint.y + yDiff}
// 	newWaypoint := point{waypoint.x + xDiff, waypoint.y + yDiff}

// 	return newShipPoint, newWaypoint
// }

// func turnWaypoint(shipPoint point, waypoint point, input nav) point {
// 	vector := point{waypoint.x - shipPoint.x, waypoint.y - shipPoint.y}

// 	if input.value == 180 {
// 		return point{shipPoint.x - vector.x, shipPoint.y - vector.y}
// 	}

// 	direction := north
// 	if vector.x > 0 && vector.y <= 0 {
// 		direction = east
// 	} else if vector.x <= 0 && vector.y < 0 {
// 		direction = south
// 	} else if vector.x < 0 && vector.y >= 0 {
// 		direction = west
// 	}

// 	newDirection := turn(input, string(direction)[0])
// 	aVector := point{absolute(vector.x), absolute(vector.y)}
// 	if newDirection == north {
// 		return point{shipPoint.x + aVector.y, shipPoint.y + aVector.x}
// 	} else if newDirection == east {
// 		return point{shipPoint.x + aVector.y, shipPoint.y - aVector.x}
// 	} else if newDirection == south {
// 		return point{shipPoint.x - aVector.y, shipPoint.y - aVector.x}
// 	} else {
// 		return point{shipPoint.x - aVector.y, shipPoint.y + aVector.x}
// 	}
// }

func turnVector(input point, nav nav) point {
	result := input
	turns := nav.value / 90

	for i := 0; i < turns; i++ {
		result = turnVector90(result, nav.action)
	}

	return result
}

func absolute(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func turnVector90(input point, direction byte) point {
	if direction == left {
		return point{-input.y, input.x}
	}

	return point{input.y, -input.x}
}

func getDirectionVector(direction byte) point {
	switch direction {
	case north:
		return point{0, 1}
	case south:
		return point{0, -1}
	case east:
		return point{1, 0}
	case west:
		return point{-1, 0}
	default:
		return point{}
	}
}

func addVector(input point, vector point) point {
	return point{input.x + vector.x, input.y + vector.y}
}

func mulitplyVector(input point, value int) point {
	return point{input.x * value, input.y * value}
}
