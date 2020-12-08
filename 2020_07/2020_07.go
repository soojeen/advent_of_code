package main

import "fmt"
import "log"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

type childBag struct {
	color string
	count int
}

type luggageRule struct {
	container string
	children  []childBag
}

type tracker map[string]bool

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input, parseError := parseInput(rawInput)
	if parseError != nil {
		log.Fatal(parseError)
	}
	resultA := countParents(input, "shiny gold")
	// resultB := countAll(input, countGroupB)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) ([]luggageRule, error) {
	var err error
	rules := strings.Split(input, "\n")
	luggageRules := make([]luggageRule, len(rules))
	containerRe := regexp.MustCompile(`(\w+ \w+) bags contain`)
	childrenRe := regexp.MustCompile(` (\d) (\w+ \w+) bags?[,.]`)

	for i, rule := range rules {
		containerMatch := containerRe.FindStringSubmatch(rule)
		childrenMatch := childrenRe.FindAllStringSubmatch(rule, -1)

		children := make([]childBag, len(childrenMatch))
		for i, childMatch := range childrenMatch {
			if childMatch == nil {
				continue
			}

			count, error := strconv.Atoi(childMatch[1])
			if error != nil {
				err = error
				break
			}
			children[i] = childBag{childMatch[2], count}
		}

		if err != nil {
			break
		}

		luggageRules[i] = luggageRule{containerMatch[1], children}
	}

	return luggageRules, err
}

func countParents(input []luggageRule, color string) int {
	childToParents := make(map[string]tracker)
	for _, rule := range input {
		for _, child := range rule.children {
			if childToParents[child.color] == nil {
				childToParents[child.color] = tracker{rule.container: true}
			}

			childToParents[child.color][rule.container] = true
		}
	}

	parents := getParents(childToParents, color)

	return len(parents)
}

func getParents(input map[string]tracker, color string) tracker {
	parents := input[color]
	moreParents := make(tracker)

	if parents == nil || len(parents) == 0 {
		return moreParents
	}

	for parent := range parents {
		moreParents = merge(moreParents, getParents(input, parent))
	}

	return merge(parents, moreParents)
}

func merge(inputA, inputB tracker) tracker {
	result := make(tracker)

	for key, value := range inputA {
		result[key] = value
	}

	for key, value := range inputB {
		result[key] = value
	}

	return result
}
