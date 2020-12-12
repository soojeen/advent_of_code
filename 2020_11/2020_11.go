package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"
import "advent_of_code/grid"

type seatLayout struct {
	grid grid.Rune
}

// type seatLayout struct {
// 	grid    grid.Rune
// 	tracker grid.Point
// }

const empty = 'L'
const floor = '.'
const occupied = '#'

// func (sl *seatLayout) runRound() bool {
// 	isDirty := false
// 	newGrid := make(grid.Rune, len(sl.grid))

// 	for y := 0; y < len(sl.grid); y++ {
// 		row := sl.grid[y]
// 		newRow := make(grid.RuneRow, len(row))

// 		for x := 0; x < len(row); x++ {
// 			sl.x = x
// 			sl.y = y
// 			value, isDirtySeat := sl.checkSeat()

// 			newRow[x] = value
// 			isDirty = isDirty || isDirtySeat
// 		}

// 		newGrid[y] = newRow
// 	}

// 	sl.grid = newGrid

// 	return isDirty
// }

// func (sl *seatLayout) runRoundLineOfSight() bool {
// 	isDirty := false
// 	newGrid := make(grid.Rune, len(sl.grid))

// 	for y := 0; y < len(sl.grid); y++ {
// 		row := sl.grid[y]
// 		newRow := make(grid.RuneRow, len(row))

// 		for x := 0; x < len(row); x++ {
// 			sl.x = x
// 			sl.y = y
// 			value, isDirtySeat := sl.checkSeatLineOfSight()

// 			newRow[x] = value
// 			isDirty = isDirty || isDirtySeat
// 		}

// 		newGrid[y] = newRow
// 	}

// 	sl.grid = newGrid

// 	return isDirty
// }

// func (sl *seatLayout) checkSeatLineOfSight() (rune, bool) {
// 	seat := sl.grid[sl.y][sl.x]
// 	adjacent := sl.getLightOfSight()

// 	switch seat {
// 	case empty:
// 		if hasNoOccupiedAdjacent(adjacent) {
// 			return occupied, true
// 		}

// 	case occupied:
// 		if hasOccupiedAdjacent(adjacent) {
// 			return empty, true
// 		}
// 	}

// 	return seat, false
// }

// func (sl *seatLayout) checkSeat() (rune, bool) {
// 	seat := sl.grid[sl.y][sl.x]
// 	adjacent := sl.getAdjacent()

// 	switch seat {
// 	case empty:
// 		if hasNoOccupiedAdjacent(adjacent) {
// 			return occupied, true
// 		}

// 	case occupied:
// 		if hasOccupiedAdjacent(adjacent) {
// 			return empty, true
// 		}
// 	}

// 	return seat, false
// }

// func (sl *seatLayout) getAdjacent() [8]rune {
// 	result := [8]rune{}

// 	// top
// 	result[0] = sl.grid[sl.y-1][sl.x-1]
// 	result[1] = sl.grid[sl.y-1][sl.x]
// 	result[2] = sl.grid[sl.y-1][sl.x+1]
// 	// bottom
// 	result[3] = sl.grid[sl.y+1][sl.x-1]
// 	result[4] = sl.grid[sl.y+1][sl.x]
// 	result[5] = sl.grid[sl.y+1][sl.x+1]
// 	// mid
// 	result[6] = sl.grid[sl.y][sl.x-1]
// 	result[7] = sl.grid[sl.y][sl.x+1]

// 	return result
// }

// func (sl *seatLayout) getLightOfSight() [8]rune {
// 	result := [8]rune{}
// 	directions := make(map[int][2]int, 8)
// 	directions[0] = [2]int{-1, -1}
// 	directions[1] = [2]int{-1, 0}
// 	directions[2] = [2]int{-1, 1}
// 	directions[3] = [2]int{1, -1}
// 	directions[4] = [2]int{1, 0}
// 	directions[5] = [2]int{1, 1}
// 	directions[6] = [2]int{0, -1}
// 	directions[7] = [2]int{0, 1}

// 	for i, direction := range directions {
// 		for j := 1; ; j++ {
// 			seat := sl.grid[sl.y+j*direction[1]][sl.x+j*direction[0]]

// 			if seat == floor {
// 				continue
// 			}

// 			result[i] = seat
// 			break
// 		}
// 	}
// }

// func (sl *seatLayout) countOccupied() int {
// 	count := 0

// 	for _, row := range sl.grid {
// 		for _, char := range row {
// 			if char == occupied {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := runA(input)
	// resultB := lineOfSight(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
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

	return seatLayout{seatGrid}
}

func runA(input seatLayout) int {
	for {
		isDirty := input.runRound(false)

		if !isDirty {
			counts := input.grid.CountAll()
			return counts[occupied]
		}
	}
}

func (s *seatLayout) runRound(lineOfSight bool) bool {
	isDirty := false
	newGrid := make(grid.Rune, len(s.grid))

	for y, row := range s.grid {
		newRow := make(grid.RuneRow, len(row))

		for x := range row {
			if !lineOfSight {
				seat := s.grid[y][x]

				adjacent := s.grid.GetAllAdjacent(grid.Point{X: x, Y: y})

				value, isDirtySeat := checkAdjacent(adjacent, seat)

				newRow[x] = value
				isDirty = isDirty || isDirtySeat
			}
		}

		newGrid[y] = newRow
	}

	s.grid = newGrid

	return isDirty
}

func checkAdjacent(adjacent [8]rune, seat rune) (rune, bool) {
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

// func stabalise(input seatLayout) int {
// 	for {
// 		isDirty := input.runRound()

// 		if !isDirty {
// 			return input.countOccupied()
// 		}
// 	}
// }

// func lineOfSight(input seatLayout) int {
// 	for {
// 		isDirty := input.runRoundLineOfSight()

// 		if !isDirty {
// 			return input.countOccupied()
// 		}
// 	}
// }

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
