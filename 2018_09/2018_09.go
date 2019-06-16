package main

import "container/list"
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
	resultB := winningScore(input, 100)

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
	current *list.Element
	marbles *list.List
}

func (g *gameMarbles) place(value int) {
	insertAfter := g.current.Next()

	if insertAfter == nil {
		insertAfter = g.marbles.Front()
	}

	g.current = g.marbles.InsertAfter(value, insertAfter)
}

func (g *gameMarbles) specialPlace() int {
	const special = 7
	specialMarble := g.current

	for i := 0; i < special; i++ {
		specialMarble = specialMarble.Prev()

		if specialMarble == nil {
			specialMarble = g.marbles.Back()
		}
	}

	// track new current before Remove
	g.current = specialMarble.Next()
	if g.current == nil {
		g.current = g.marbles.Front()
	}

	removeValue := g.marbles.Remove(specialMarble)

	return removeValue.(int)
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
	gameMarbles := initializeMarbles()

	for i := 1; i <= gameInput.points*multiplier; i++ {
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

func initializeMarbles() gameMarbles {
	marbles := list.New()
	current := marbles.PushBack(0)

	return gameMarbles{current, marbles}
}
