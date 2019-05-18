package main

import "fmt"

import "log"

// import "strconv"
// import "sort"
import "regexp"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// resultA := len(reactPolymers(rawInput))
	resultB := shortestPolymer(rawInput)
	// fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func reactPolymers(input string) string {
	polymers := strings.Split(input, "")
	polymerBuffer := []string{}
	bufferIdx := 0
	polymersIdx := 0

	for {
		if len(polymerBuffer) == 0 {
			polymerBuffer = append(polymerBuffer, polymers[polymersIdx])
			bufferIdx = 0
			polymersIdx++
		}

		remove := reactPolymer(polymers[polymersIdx], polymerBuffer[bufferIdx])

		if remove {
			newStackSize := len(polymerBuffer) - 1
			polymerBuffer[newStackSize] = ""
			polymerBuffer = polymerBuffer[:newStackSize]

			bufferIdx--
			polymersIdx++
		} else {
			polymerBuffer = append(polymerBuffer, polymers[polymersIdx])
			bufferIdx++
			polymersIdx++
		}

		if polymersIdx >= len(polymers) {
			break
		}
	}

	return strings.Join(polymerBuffer, "")
}

func shortestPolymer(input string) int {
	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	minPolymer := len(input)
	c := make(chan int, len(alphabet))

	for _, letter := range alphabet {
		letter := letter
		go func() {
			regexString := strings.Join([]string{letter, "|", strings.ToUpper(letter)}, "")
			regex := regexp.MustCompile(regexString)

			polymer := regex.ReplaceAllString(input, "")

			result := len(reactPolymers(polymer))

			c <- result
		}()
	}

	for i := 0; i < len(alphabet); i++ {
		result := <-c

		if result < minPolymer {
			minPolymer = result
		}
	}

	return minPolymer
}

func reactPolymer(letter1 string, letter2 string) bool {
	if strings.ToLower(letter1) == strings.ToLower(letter2) && letter1 != letter2 {
		return true
	}

	return false
}
