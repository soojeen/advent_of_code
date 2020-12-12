package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"
import "advent_of_code/grid"

type seatLayout struct {
	grid        grid.Rune
	isAdjacent  bool
	maxOccupied int
}

const empty = 'L'
const floor = '.'
const occupied = '#'

func (s *seatLayout) runRound() bool {
	isDirty := false
	newGrid := make(grid.Rune, len(s.grid))

	for y, row := range s.grid {
		newRow := make(grid.RuneRow, len(row))

		for x := range row {
			seat := s.grid[y][x]

			if s.isAdjacent {
				adjacent := s.grid.GetAllAdjacent(grid.Point{X: x, Y: y})
				value, isDirtySeat := processSeat(seat, adjacent, s.maxOccupied)

				newRow[x] = value
				isDirty = isDirty || isDirtySeat
			} else {
				visibleSeats := [8]rune{}
				callback := func(direction grid.Point, index int) {
					for i := 1; ; i++ {
						nextLineOfSight := s.grid[y+i*direction.Y][x+i*direction.X]
						if nextLineOfSight == floor {
							continue
						}

						visibleSeats[index] = nextLineOfSight
						break
					}
				}
				grid.ForEachDirection(grid.Point{X: x, Y: y}, callback)
				value, isDirtySeat := processSeat(seat, visibleSeats, s.maxOccupied)

				newRow[x] = value
				isDirty = isDirty || isDirtySeat
			}
		}

		newGrid[y] = newRow
	}

	s.grid = newGrid

	return isDirty
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := run(input)

	input.maxOccupied = 5
	input.isAdjacent = false
	resultB := run(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) seatLayout {
	lines := strings.Split(input, "\n")
	seatGrid := make(grid.Rune, len(lines))

	for y, line := range lines {
		row := make(grid.RuneRow)

		for x, char := range line {
			row[x] = char
		}

		seatGrid[y] = row
	}

	return seatLayout{seatGrid, true, 4}
}

func run(input seatLayout) int {
	for {
		isDirty := input.runRound()

		if !isDirty {
			counts := input.grid.CountAll()
			return counts[occupied]
		}
	}
}

func processSeat(seat rune, adjacent [8]rune, maxOccupied int) (rune, bool) {
	seatDistribution := checkSeats(adjacent, seat)

	if seat == empty && seatDistribution[occupied] == 0 {
		return occupied, true
	}

	if seat == occupied && seatDistribution[occupied] >= maxOccupied {
		return empty, true
	}

	return seat, false
}

func checkSeats(input [8]rune, seat rune) map[rune]int {
	result := map[rune]int{}
	for _, space := range input {
		result[space]++
	}

	return result
}
