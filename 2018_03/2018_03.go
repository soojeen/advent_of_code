package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type Claim struct {
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

	resultA, resultB := claimOverlap(input)
	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func claimOverlap(claims []Claim) (int, int) {
	type ClaimStatus struct {
		claimed      bool
		multiClaimed bool
		claimants    []int
	}

	var grid [1000][1000]ClaimStatus
	count := 0
	noOverlapId := 0

	for _, claim := range claims {
		for x := claim.cornerX; x < (claim.cornerX + claim.width); x++ {
			for y := claim.cornerY; y < (claim.cornerY + claim.height); y++ {
				claimStatus := &grid[x][y]
				if claimStatus.claimed && !claimStatus.multiClaimed {
					claimStatus.multiClaimed = true
					claimStatus.claimants = append(claimStatus.claimants, claim.id)
					count++
				} else if !claimStatus.claimed {
					claimStatus.claimed = true
				}
			}
		}
	}

	return count, noOverlapId
}

func parseInput(input []string) ([]Claim, error) {
	var err error
	claims := make([]Claim, len(input))

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

		claims[i] = Claim{i + 1, cornerX, cornerY, width, height}
	}

	return claims, err
}
