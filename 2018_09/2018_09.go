package main

import "fmt"

// import "log"
import "strconv"
import "strings"

// import "advent_of_code/utils"

func main() {
	// rawInput, err := utils.ReadInput("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// rawInput := "9 players; last marble is worth 25 points"
	rawInput := "10 players; last marble is worth 1618 points"
	// rawInput := "13 players; last marble is worth 7999 points"
	// rawInput := "17 players; last marble is worth 1104 points"
	// rawInput := "21 players; last marble is worth 6111 points"
	// rawInput := "30 players; last marble is worth 5807 points"

	input := parseInput(rawInput)

	resultA := winningScore(input)
	// result := parseTreeSum(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", result.value)
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

	if g.currentIndex == lastIndex-1 || len(g.marbles) == 1 {
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
		removeIndex = len(g.marbles) - 1 + removeIndex
	}

	removeValue := g.marbles[removeIndex]

	lastIndex := len(g.marbles) - 1
	copy(g.marbles[removeIndex:], g.marbles[removeIndex+1:])
	g.marbles[lastIndex] = 0
	g.marbles = g.marbles[:lastIndex]

	g.currentIndex = removeIndex

	return removeValue
}

type player struct {
	id    int
	score int
}
type gamePlayers struct {
	current int
	players []player
}

func (g *gamePlayers) next() {
	if len(g.players) == g.current {
		g.current = 1
	}

	g.current++
}

func (g *gamePlayers) score(points int) {
	g.players[g.current-1].score += points
}

func (g *gamePlayers) highScore() int {
	result := 0

	for _, player := range g.players {
		if player.score > result {
			result = player.score
		}
	}

	return result
}

func winningScore(gameInput gameInput) int {
	gamePlayers := initializePlayers(gameInput.players)
	gameMarbles := gameMarbles{0, make([]int, 1)}

	for i := 1; i <= gameInput.points; i++ {
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
	players := make([]player, totalPlayers)

	for i := range players {
		players[i].id = i + 1
	}

	return gamePlayers{1, players}
}
