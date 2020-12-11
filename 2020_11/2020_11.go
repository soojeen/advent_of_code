package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

type grid [][]string
type tracker struct {
	x int
	y int
}
type seatLayout struct {
	grid    grid
	tracker tracker
}

const empty = "L"
const occupied = "#"

func (sl *seatLayout) runRound() bool {
	hasChange := false

	return hasChange
}

func (sl *seatLayout) checkSeat() (string, bool) {
	seat := sl.grid[sl.tracker.x][sl.tracker.y]
	adjacent := sl.getAdjacent()

	switch seat {
	case empty:
		if hasNoOccupiedAdjacent(adjacent) {
			return occupied, true
		}

	case occupied:
		if hasOccupiedAdjacent(adjacent) {
			return empty, true
		}
	}

	return seat, false

}

func (sl *seatLayout) getAdjacent() []string {
	x := sl.tracker.x
	y := sl.tracker.y
	maxX := len(sl.grid[0]) - 1
	maxY := len(sl.grid) - 1

	if x == 0 && y == 0 {
		return []string{sl.grid[y+1][x], sl.grid[y][x+1], sl.grid[y+1][x+1]}
	}

	if x == 0 && y == maxY {
		return []string{sl.grid[y][x+1], sl.grid[y-1][x], sl.grid[y-1][x+1]}
	}

	if x == maxX && y == maxY {
		return []string{sl.grid[y-1][x], sl.grid[y][x-1], sl.grid[y-1][x-1]}
	}

	if x == maxX && y == 0 {
		return []string{sl.grid[y][x-1], sl.grid[y+1][x], sl.grid[y+1][x-1]}
	}

	if x == 0 {
		return []string{sl.grid[y-1][x], sl.grid[y+1][x], sl.grid[y][x+1], sl.grid[y-1][x+1], sl.grid[y+1][x+1]}
	}

	if x == maxX {
		return []string{sl.grid[y-1][x], sl.grid[y+1][x], sl.grid[y][x-1], sl.grid[y-1][x-1], sl.grid[y+1][x-1]}
	}

	if y == 0 {
		return []string{sl.grid[y][x-1], sl.grid[y][x+1], sl.grid[y+1][x], sl.grid[y+1][x-1], sl.grid[y+1][x+1]}
	}

	if y == maxY {
		return []string{sl.grid[y][x-1], sl.grid[y][x+1], sl.grid[y-1][x], sl.grid[y-1][x-1], sl.grid[y-1][x+1]}
	}

	top := sl.grid[y-1][x-1 : x+2]
	bottom := sl.grid[y+1][x-1 : x+2]

	result := append(top, bottom...)

	return append(result, sl.grid[y][x-1], sl.grid[y][x+1])
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := stabalise(input)
	// resultB := findArrangements(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) seatLayout {
	lines := strings.Split(input, "\n")
	grid := make(grid, len(lines))

	for i, line := range lines {
		row := strings.Split(line, "")
		grid[i] = row
	}

	return seatLayout{grid, tracker{0, 0}}
}

func stabalise(input seatLayout) int {

	fmt.Println("a: ", input.grid[0])

	return 0
}

func hasNoOccupiedAdjacent(input []string) bool {
	for _, space := range input {
		if space == occupied {
			return false
		}
	}

	return true
}

func hasOccupiedAdjacent(input []string) bool {
	count := 0

	for _, space := range input {
		if space == occupied {
			count++
		}
	}

	return count >= 4
}
