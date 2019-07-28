package main

import "fmt"

const gridSize = 300

type largest struct {
	sum  int
	x    int
	y    int
	size int
}

func main() {
	input := 1955

	grid := generatePowerLevels(input)

	resultA := findLargest(grid)
	fmt.Println(resultA)
}

func generatePowerLevels(input int) [][]int {
	grid := make([][]int, gridSize)

	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			grid[x] = append(grid[x], getPowerLevel(x, y, input))
		}
	}

	return grid
}

func findLargest(grid [][]int) largest {
	c := make(chan largest, gridSize)

	for i := 0; i <= gridSize; i++ {
		i := i + 1

		go func() {
			var large largest
			upperBound := len(grid) - i

			fmt.Println(i)
			for x := 0; x < upperBound; x++ {
				for y := 0; y < upperBound; y++ {
					sum := powerSum(grid, x, y, i)

					if sum > large.sum {
						large = largest{sum, x, y, i}
					}
				}
			}

			c <- large
		}()
	}

	var mostLargest largest
	for j := 0; j <= gridSize; j++ {
		large := <-c

		if large.sum > mostLargest.sum {
			mostLargest = large
		}
	}

	return mostLargest
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

func powerSum(grid [][]int, x int, y int, interfaceSize int) int {
	sum := 0

	for i := 0; i < interfaceSize; i++ {
		for j := 0; j < interfaceSize; j++ {
			sum = sum + grid[x+i][y+j]
		}
	}

	return sum
}
