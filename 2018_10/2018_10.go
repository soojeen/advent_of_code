package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
import "sort"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(rawInput)

	mainLights(input)
}

type light struct {
	position [2]int
	velocity [2]int
}

type lights struct {
	lights     []light
	positionsX []int
	positionsY []int
}

func (l *lights) move() {
	var positionsX []int
	var positionsY []int

	for i := range l.lights {
		l.lights[i].position[0] += l.lights[i].velocity[0]
		l.lights[i].position[1] += l.lights[i].velocity[1]

		positionsX = append(positionsX, l.lights[i].position[0])
		positionsY = append(positionsY, l.lights[i].position[1])
	}

	sort.Ints(positionsX)
	sort.Ints(positionsY)

	l.positionsX = positionsX
	l.positionsY = positionsY
}

func (l *lights) draw() {
	const margin = 10
	const inputLen = 391

	mapping := [inputLen][inputLen]string{}

	for y, row := range mapping {
		for x := range row {
			mapping[y][x] = "."
		}
	}

	for _, light := range l.lights {
		mapping[light.position[1]][light.position[0]] = "#"
	}

	for _, row := range mapping {
		fmt.Println(strings.Join(row[:], ""))
	}
}

func parseInput(rawInput string) lights {
	r, _ := regexp.Compile("position=<(.{6}), (.{6})> velocity=<(.{2}), (.{2})>")
	lines := strings.Split(rawInput, "\n")

	lightsList := make([]light, len(lines))
	var positionsX []int
	var positionsY []int

	for i, line := range lines {
		match := r.FindStringSubmatch(line)

		positionX, _ := strconv.Atoi(strings.TrimSpace(match[1]))
		positionY, _ := strconv.Atoi(strings.TrimSpace(match[2]))
		velocityX, _ := strconv.Atoi(strings.TrimSpace(match[3]))
		velocityY, _ := strconv.Atoi(strings.TrimSpace(match[4]))

		lightsList[i] = light{[2]int{positionX, positionY}, [2]int{velocityX, velocityY}}

		positionsX = append(positionsX, positionX)
		positionsY = append(positionsY, positionY)
	}

	sort.Ints(positionsX)
	sort.Ints(positionsY)

	return lights{lightsList, positionsX, positionsY}
}

func mainLights(lights lights) {
	const diffThreshold = 100
	lightsLen := len(lights.lights)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; ; i++ {
		xDiff := lights.positionsX[0] - lights.positionsX[lightsLen-1]
		xDiffStop := (xDiff < 0 && xDiff > -diffThreshold) || (xDiff > 0 && xDiff < diffThreshold)

		yDiff := lights.positionsY[0] - lights.positionsY[lightsLen-1]
		yDiffStop := (yDiff < 0 && yDiff > -diffThreshold) || (yDiff > 0 && yDiff < diffThreshold)

		if xDiffStop && yDiffStop {
			for j := i; ; j++ {
				lights.draw()
				fmt.Println("second: ", j)
				scanner.Scan()
				lights.move()
			}
		}

		lights.move()
	}
}
