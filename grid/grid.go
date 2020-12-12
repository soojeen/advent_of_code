package grid

// RuneRow - row in grid
type RuneRow map[int]rune

// Rune - 2-D grid
type Rune map[int]RuneRow

// Point - point
type Point struct {
	X int
	Y int
}

// GetAllAdjacent - get all immediately adjacent values
func (gR *Rune) GetAllAdjacent(input Point) [8]rune {
	result := [8]rune{}
	callback := func(direction Point, index int) {
		result[index] = (*gR)[input.Y+direction.Y][input.X+direction.X]
	}

	ForEachDirection(input, callback)

	return result
}

// RunePrettyPrint - pretty grid
func (gR *Rune) RunePrettyPrint() string {
	prettyGrid := ""

	for y := 0; y < len((*gR)); y++ {
		row := (*gR)[y]
		prettyRow := ""

		for x := 0; x < len(row); x++ {
			prettyRow = prettyRow + string(row[x])
		}

		prettyGrid = prettyGrid + prettyRow + "\n"
	}

	return prettyGrid
}

// CountAll - return instance count of all
func (gR *Rune) CountAll() map[rune]int {
	result := map[rune]int{}

	for _, row := range *gR {
		for _, char := range row {
			result[char]++
		}
	}

	return result
}

// ForEachDirection - execute callback for each direction
func ForEachDirection(input Point, eachDirection func(direction Point, index int)) {
	directions := getDirections()

	for i, direction := range directions {
		eachDirection(direction, i)
	}
}

func getDirections() [8]Point {
	var directions [8]Point

	directions[0] = Point{-1, -1}
	directions[1] = Point{-1, 0}
	directions[2] = Point{-1, 1}
	directions[3] = Point{1, -1}
	directions[4] = Point{1, 0}
	directions[5] = Point{1, 1}
	directions[6] = Point{0, -1}
	directions[7] = Point{0, 1}

	return directions
}
