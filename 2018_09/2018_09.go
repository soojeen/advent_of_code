package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(rawInput)

	resultA := winningScore(input, 1)
	resultB := winningScore(input, 10)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

type gameInput struct {
	players int
	points  int
	special int
}

func parseInput(rawInput string) gameInput {
	values := strings.Split(rawInput, " ")
	players, _ := strconv.Atoi(values[0])
	points, _ := strconv.Atoi(values[6])

	return gameInput{players, points, 23}
}

type gameMarbles struct {
	currentIndex int
	marbles      []int
}

func (g *gameMarbles) place(value int) {
	lastIndex := len(g.marbles) - 1

	if g.currentIndex == lastIndex-1 || g.currentIndex == 0 {
		g.marbles = append(g.marbles, value)
		g.currentIndex = len(g.marbles) - 1
	} else {
		insertIndex := 1
		if g.currentIndex != lastIndex {
			insertIndex = g.currentIndex + 2
		}

		g.marbles = append(g.marbles, 0)
		copy(g.marbles[insertIndex+1:], g.marbles[insertIndex:])
		g.marbles[insertIndex] = value
		g.currentIndex = insertIndex
	}
}

func (g *gameMarbles) specialPlace() int {
	removeIndex := g.currentIndex - 7

	if removeIndex < 0 {
		removeIndex = len(g.marbles) + removeIndex

	}

	removeValue := g.marbles[removeIndex]

	lastIndex := len(g.marbles) - 1
	if removeIndex >= lastIndex {
		fmt.Println("a:", "warning")

	}
	copy(g.marbles[removeIndex:], g.marbles[removeIndex+1:])
	g.marbles[lastIndex] = 0
	g.marbles = g.marbles[:lastIndex]

	g.currentIndex = removeIndex

	return removeValue
}

type gamePlayers struct {
	current int
	scores  []int
}

func (g *gamePlayers) next() {
	if len(g.scores)-1 == g.current {
		g.current = 0
	} else {
		g.current++
	}
}

func (g *gamePlayers) score(points int) {
	g.scores[g.current] += points
}

func (g *gamePlayers) highScore() int {
	result := 0

	for _, score := range g.scores {
		if score > result {
			result = score
		}
	}

	return result
}

func winningScore(gameInput gameInput, multiplier int) int {
	gamePlayers := initializePlayers(gameInput.players)
	gameMarbles := gameMarbles{0, make([]int, 1)}

	for i := 1; i <= gameInput.points*multiplier; i++ {
		if i%gameInput.points == 0 {
			fmt.Println("current", i)
		}
		if i%gameInput.special == 0 {
			special := gameMarbles.specialPlace()
			gamePlayers.score(i + special)
		} else {
			gameMarbles.place(i)
		}

		gamePlayers.next()
	}

	return gamePlayers.highScore()
}

func initializePlayers(totalPlayers int) gamePlayers {
	return gamePlayers{0, make([]int, totalPlayers)}
}
