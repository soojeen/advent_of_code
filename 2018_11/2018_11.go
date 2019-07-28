package main

// import "bufio"
import "fmt"

// import "log"
// import "os"
// import "regexp"
// import "strconv"
// import "strings"
// import "advent_of_code/utils"

func main() {
	const gridSize = 300
	input := 1955
	// input := 18

	grid := generatePowerLevels(gridSize, input)

	resultA := findLargest(grid)
	fmt.Println(resultA)
}

func generatePowerLevels(gridSize int, input int) [][]int {
	grid := make([][]int, gridSize)

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			grid[x] = append(grid[x], getPowerLevel(x, y, input))
		}
	}

	return grid
}

func findLargest(grid [][]int) [2]int {
	type largest struct {
		sum int
		x   int
		y   int
	}

	const interfaceSize = 3
	var large = largest{-46, 0, 0}
	upperBound := len(grid) - interfaceSize

	for x := 0; x < upperBound; x++ {
		for y := 0; y < upperBound; y++ {
			sum := powerSum(grid, x, y)

			if sum > large.sum {
				large = largest{sum, x, y}
			}
		}
	}

	return [2]int{large.x, large.y}
}

func getPowerLevel(x int, y int, input int) int {
	rackID := x + 1 + 10
	powerLevel := rackID * (y + 1)
	powerLevel = powerLevel + input
	powerLevel = powerLevel * rackID
	powerLevel = hundredsDigit(powerLevel)
	powerLevel = powerLevel - 5

	return powerLevel
}

func hundredsDigit(number int) int {
	if number < 100 {
		return 0
	}

	number = number / 100

	return number % 10
}

func powerSum(grid [][]int, x int, y int) int {
	sum := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum = sum + grid[x+i][y+j]
		}
	}

	return sum
}
