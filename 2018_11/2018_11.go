package main

// import "bufio"
import "fmt"

// import "log"
// import "os"
// import "regexp"
// import "strconv"
// import "strings"
// import "advent_of_code/utils"

const gridSize = 300

type grid [gridSize][gridSize]int

func main() {
	input := 1955

	grid := generatePowerLevels(input)
	fmt.Println(input)
	fmt.Println(grid[0])
}

func generatePowerLevels(input int) grid {
	grid := [gridSize][gridSize]int{}

	return grid
}
