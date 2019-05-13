package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type claim struct {
	id      int
	cornerX int
	cornerY int
	width   int
	height  int
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input, err := parseInput(strings.Split(rawInput, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	resultA, resultB := resolveClaims(input)
	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func resolveClaims(claims []claim) (int, int) {
	var grid [1000][1000][]int
	claimOverlaps := make(map[int]bool)
	count := 0
	noOverlapID := 0

	for _, claim := range claims {
		for x := claim.cornerX; x < (claim.cornerX + claim.width); x++ {
			for y := claim.cornerY; y < (claim.cornerY + claim.height); y++ {
				claimants := grid[x][y]

				if len(claimants) == 1 {
					count++
				}

				claimants = append(claimants, claim.id)
				grid[x][y] = claimants

				if len(claimants) > 1 {
					for _, claimant := range claimants {
						claimOverlaps[claimant] = true
					}
				}
			}
		}
	}

	for _, claim := range claims {
		if !claimOverlaps[claim.id] {
			noOverlapID = claim.id
		}
	}

	return count, noOverlapID
}

func parseInput(input []string) ([]claim, error) {
	var err error
	claims := make([]claim, len(input))

	for i, rawClaim := range input {
		claimParts := strings.Split(rawClaim, " ")
		corner := strings.Split(claimParts[2], ",")

		cornerX, e := strconv.Atoi(corner[0])
		if e != nil {
			err = e
			break
		}

		cornerY, e := strconv.Atoi(strings.TrimSuffix(corner[1], ":"))
		if e != nil {
			err = e
			break
		}

		dimensions := strings.Split(claimParts[3], "x")

		width, e := strconv.Atoi(dimensions[0])
		if e != nil {
			err = e
			break
		}

		height, e := strconv.Atoi(dimensions[1])
		if e != nil {
			err = e
			break
		}

		claims[i] = claim{i + 1, cornerX, cornerY, width, height}
	}

	return claims, err
}
