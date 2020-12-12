package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

type grid map[int]map[int]rune
type seatLayout struct {
	grid grid
	x    int
	y    int
}

const empty = 'L'
const occupied = '#'

func (sl *seatLayout) runRound() bool {
	isDirty := false
	newGrid := make(grid, len(sl.grid))

	for y := 0; y < len(sl.grid); y++ {
		row := sl.grid[y]
		newRow := make(map[int]rune, len(row))

		for x := 0; x < len(row); x++ {
			sl.x = x
			sl.y = y
			value, isDirtySeat := sl.checkSeat()

			newRow[x] = value
			isDirty = isDirty || isDirtySeat
		}

		newGrid[y] = newRow
	}

	sl.grid = newGrid

	return isDirty
}

func (sl *seatLayout) checkSeat() (rune, bool) {
	seat := sl.grid[sl.y][sl.x]
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

func (sl *seatLayout) getAdjacent() [8]rune {
	result := [8]rune{}

	// top
	result[0] = sl.grid[sl.y-1][sl.x-1]
	result[1] = sl.grid[sl.y-1][sl.x]
	result[2] = sl.grid[sl.y-1][sl.x+1]
	// bottom
	result[3] = sl.grid[sl.y+1][sl.x-1]
	result[4] = sl.grid[sl.y+1][sl.x]
	result[5] = sl.grid[sl.y+1][sl.x+1]
	// mid
	result[6] = sl.grid[sl.y][sl.x-1]
	result[7] = sl.grid[sl.y][sl.x+1]

	return result
}

func (sl *seatLayout) countOccupied() int {
	count := 0

	for _, row := range sl.grid {
		for _, char := range row {
			if char == occupied {
				count++
			}
		}
	}

	return count
}

func (sl *seatLayout) prettyPrint() string {
	prettyGrid := ""

	for y := 0; y < len(sl.grid); y++ {
		row := sl.grid[y]
		prettyRow := ""

		for x := 0; x < len(row); x++ {
			prettyRow = prettyRow + string(row[x])
		}

		prettyGrid = prettyGrid + prettyRow + "\n"
	}

	return prettyGrid
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

	for y, line := range lines {
		row := make(map[int]rune)

		for x, char := range line {
			row[x] = char
		}

		grid[y] = row
	}

	return seatLayout{grid, 0, 0}
}

func stabalise(input seatLayout) int {
	for {
		isDirty := input.runRound()

		if !isDirty {
			return input.countOccupied()
		}
	}
}

func hasNoOccupiedAdjacent(input [8]rune) bool {
	for _, space := range input {
		if space == occupied {
			return false
		}
	}

	return true
}

func hasOccupiedAdjacent(input [8]rune) bool {
	count := 0

	for _, space := range input {
		if space == occupied {
			count++
		}
	}

	return count >= 4
}
