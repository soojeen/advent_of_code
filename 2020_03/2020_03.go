package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type passwordPolicy struct {
	min      int
	max      int
	match    string
	password string
}

func (pp *passwordPolicy) isValid() bool {
	count := strings.Count(pp.password, pp.match)

	return count >= pp.min && count <= pp.max
}

func (pp *passwordPolicy) isValidReal() bool {
	x := string(pp.password[pp.min-1]) == pp.match
	y := string(pp.password[pp.max-1]) == pp.match

	return x != y
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input, parseError := parseInput(rawInput)
	if parseError != nil {
		log.Fatal(parseError)
	}

	resultA := findValid(input)
	resultB := findValidReal(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) ([]passwordPolicy, error) {
	lines := strings.Split(input, "\n")
	result := make([]passwordPolicy, len(lines))
	var err error

	for i, line := range lines {
		lineParts := strings.Split(line, ": ")
		policyParts := strings.Split(lineParts[0], " ")
		minMax := strings.Split(policyParts[0], "-")

		min, e := strconv.Atoi(minMax[0])
		if e != nil {
			err = e
			break
		}

		max, e := strconv.Atoi(minMax[1])
		if e != nil {
			err = e
			break
		}

		policy := passwordPolicy{min, max, policyParts[1], lineParts[1]}

		result[i] = policy
	}

	return result, err
}

func findValid(input []passwordPolicy) int {
	count := 0

	for _, password := range input {
		isValid := password.isValid()
		if isValid {
			count++
		}
	}

	return count
}

func findValidReal(input []passwordPolicy) int {
	count := 0

	for _, password := range input {
		isValid := password.isValidReal()
		if isValid {
			count++
		}
	}

	return count
}
