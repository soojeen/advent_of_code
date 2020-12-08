package main

import "fmt"
import "log"
import "regexp"
import "strings"
import "advent_of_code/utils"

type luggageRule struct {
	container string
	children  []string
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := countParents(input, "shiny gold")
	// resultB := countAll(input, countGroupB)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) []luggageRule {
	rules := strings.Split(input, "\n")
	luggageRules := make([]luggageRule, len(rules))
	containerRe := regexp.MustCompile(`(\w+ \w+) bags contain`)
	childrenRe := regexp.MustCompile(` \d (\w+ \w+) bags?[,.]`)

	for i, rule := range rules {
		containerMatch := containerRe.FindStringSubmatch(rule)
		childrenMatch := childrenRe.FindAllStringSubmatch(rule, -1)

		children := make([]string, len(childrenMatch))
		for i, childMatch := range childrenMatch {
			if childMatch == nil {
				continue
			}

			children[i] = childMatch[1]
		}

		luggageRules[i] = luggageRule{containerMatch[1], children}
	}

	return luggageRules
}

func countParents(input []luggageRule, color string) int {
	childToParents := make(map[string]map[string]bool)
	for _, rule := range input {
		for _, child := range rule.children {
			if childToParents[child] == nil {
				childToParents[child] = map[string]bool{rule.container: true}
			}

			childToParents[child][rule.container] = true
		}
	}

	parents := getParents(childToParents, color)

	return len(parents)
}

func getParents(input map[string]map[string]bool, color string) map[string]bool {
	parents := input[color]

	if parents == nil || len(parents) == 0 {
		input[color] = nil
		return make(map[string]bool)
	}

	moreParents := make(map[string]bool)
	for parent := range parents {
		moreParents = merge(moreParents, getParents(input, parent))
		input[parent] = nil
	}

	return merge(parents, moreParents)
}

func merge(inputA, inputB map[string]bool) map[string]bool {
	result := make(map[string]bool)

	for key, value := range inputA {
		result[key] = value
	}

	for key, value := range inputB {
		result[key] = value
	}

	return result
}
