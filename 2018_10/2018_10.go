package main

import "bufio"
import "fmt"
import "log"
import "os"
import "regexp"
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
	lights []light
}

func (l *lights) minMax() (int, int, int, int) {
	initialLight := l.lights[0]
	maxX := initialLight.position[0]
	maxY := initialLight.position[1]
	minX := maxX
	minY := maxY

	for _, light := range l.lights {
		xPos := light.position[0]
		yPos := light.position[1]

		if xPos > maxX {
			maxX = xPos
		} else if xPos < minX {
			minX = xPos
		}

		if yPos > maxY {
			maxY = yPos
		} else if yPos < minY {
			minY = yPos
		}
	}

	return minX, minY, maxX, maxY
}

func (l *lights) move() {
	for i, light := range l.lights {
		l.lights[i].position[0] += light.velocity[0]
		l.lights[i].position[1] += light.velocity[1]
	}
}

func (l *lights) draw() {
	minX, minY, maxX, maxY := l.minMax()

	mapping := make([][]string, absolute(minY, maxY)+2)
	for y := range mapping {
		row := make([]string, absolute(minX, maxX)+2)

		for x := range row {
			row[x] = "."
		}

		mapping[y] = row
	}

	for _, light := range l.lights {
		xPos := light.position[0] - minX
		yPos := light.position[1] - minY

		mapping[yPos][xPos] = "#"
	}

	for _, row := range mapping {
		fmt.Println(strings.Join(row[:], ""))
	}
}

func parseInput(rawInput string) lights {
	r, _ := regexp.Compile("position=<(.{6}), (.{6})> velocity=<(.{2}), (.{2})>")
	lines := strings.Split(rawInput, "\n")

	lightsList := make([]light, len(lines))

	for i, line := range lines {
		match := r.FindStringSubmatch(line)

		positionX, _ := strconv.Atoi(strings.TrimSpace(match[1]))
		positionY, _ := strconv.Atoi(strings.TrimSpace(match[2]))
		velocityX, _ := strconv.Atoi(strings.TrimSpace(match[3]))
		velocityY, _ := strconv.Atoi(strings.TrimSpace(match[4]))

		lightsList[i] = light{[2]int{positionX, positionY}, [2]int{velocityX, velocityY}}
	}

	return lights{lightsList}
}

func mainLights(lights lights) {
	const diffThreshold = 100
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; ; i++ {
		minX, minY, maxX, maxY := lights.minMax()

		xDiffStop := absolute(minX, maxX) < diffThreshold
		yDiffStop := absolute(minY, maxY) < diffThreshold

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

func absolute(valA int, valB int) int {
	diff := valA - valB

	if diff < 0 {
		return -diff
	}

	return diff
}
