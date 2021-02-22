package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

type axial struct {
	q int
	r int
}
type pendingValue struct {
	current bool
	pending bool
}
type axialMap map[axial]pendingValue

func (a *axial) translate(input axial) {
	a.q += input.q
	a.r += input.r
}

func (a *axialMap) count() int {
	result := 0

	for _, value := range *a {
		if value.current {
			result++
		}
	}

	return result
}

func (a *axialMap) getAdjacent(input axial) axialMap {
	result := axialMap{}
	compassAxial := getCompassAxial()

	for _, directionAxial := range compassAxial {
		adjacentAxial := translateAxial(input, directionAxial)
		result[adjacentAxial] = (*a)[adjacentAxial]
	}

	return result
}

func (a *axialMap) dayFlip() {
	for axial, axialValue := range *a {
		if axialValue.current == false {
			continue
		}

		adjacent := a.getAdjacent(axial)
		outerPendingValue := pendingValue{true, true}

		adjacentCount := adjacent.count()
		if adjacentCount == 0 || adjacentCount > 2 {
			outerPendingValue.pending = false
		}
		(*a)[axial] = outerPendingValue

		for adjacentAxial, adjacentAxialValue := range adjacent {
			if adjacentAxialValue.current == true {
				continue
			}

			nestedAdjacent := a.getAdjacent(adjacentAxial)
			nestedPendingValue := pendingValue{false, false}

			if nestedAdjacent.count() == 2 {
				nestedPendingValue.pending = true
			}

			(*a)[adjacentAxial] = nestedPendingValue
		}
	}
}

func (a *axialMap) applyPending() {
	for axial, axialValue := range *a {
		(*a)[axial] = pendingValue{axialValue.pending, false}
	}
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA, initialState := flipTiles(input)
	resultB := multipleFlips(initialState)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func flipTiles(input []string) (int, axialMap) {
	axialMap := axialMap{}

	for _, line := range input {
		axial := findTile(line)
		pendingValue := pendingValue{false, true}

		if axialMap[axial].pending {
			pendingValue.pending = false
		}

		axialMap[axial] = pendingValue
	}

	axialMap.applyPending()

	return axialMap.count(), axialMap
}

func findTile(input string) axial {
	compassAxial := getCompassAxial()
	current := axial{0, 0}
	direction := ""
	skipNext := false

	for i, char := range input {
		if skipNext {
			skipNext = false
			continue
		}

		if char == 'e' || char == 'w' {
			direction = string(char)
		} else {
			direction = string([]rune{char, rune(input[i+1])})
			skipNext = true
		}

		axialTranslation := compassAxial[direction]
		current.translate(axialTranslation)
	}

	return current
}

func getCompassAxial() map[string]axial {
	result := map[string]axial{
		"e":  axial{1, 0},
		"w":  axial{-1, 0},
		"se": axial{0, 1},
		"sw": axial{-1, 1},
		"ne": axial{1, -1},
		"nw": axial{0, -1},
	}

	return result
}

func translateAxial(input axial, translator axial) axial {
	return axial{input.q + translator.q, input.r + translator.r}
}

func multipleFlips(input axialMap) int {
	for i := 0; i < 100; i++ {
		input.dayFlip()
		input.applyPending()
	}

	return input.count()
}
