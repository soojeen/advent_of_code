package main

import "fmt"
import "log"
import "regexp"
import "strings"
import "advent_of_code/utils"

// import "bufio"
// import "os"
// import "strconv"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	initial, patterns := parseInput(rawInput)

	generations(initial, patterns)
}

type pattern struct {
	match   string
	replace string
}

func parseInput(rawInput string) (string, []pattern) {
	input := strings.Split(rawInput, "\n")

	r, _ := regexp.Compile(`(state:\s)(.*)`)
	initial := r.FindStringSubmatch(input[0])[2]

	patterns := make([]pattern, len(input)-2)
	for i := range patterns {
		rawPattern := strings.Split(input[i+2], " => ")
		patterns[i] = pattern{rawPattern[0], rawPattern[1]}
	}

	return initial, patterns
}

func generations(initial string, patterns []pattern) {
	center := 0
	current := initial

	for i := 0; i < 20; i++ {
		switch {
		case current[:2] == ".#":
			current = "." + current
			center++
		case current[:1] == "#":
			current = ".." + current
			center += 2
		case current[len(current)-2:] == "#.":
			current = current + "."
		case current[len(current)-1:] == "#":
			current = current + ".."
		}

		// for j := range current {

		// }
	}

	fmt.Println("a:", current)
	fmt.Println("a:", center)
}
