package main

import "fmt"
import "log"
import "strings"
import "advent_of_code/utils"

type node struct {
	value    string
	children []*node
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	root := parseInput(rawInput)

	resultA := process(root)
	fmt.Println(resultA)
}

func parseInput(rawInput string) *node {
	lookup := make(map[string]*node)
	// to find root
	// parentLookup := make(map[string]*node)

	orbitPairs := strings.Split(rawInput, "\n")

	for _, orbitPair := range orbitPairs {
		orbitPair := strings.Split(orbitPair, ")")

		valueB := orbitPair[1]
		nodeB := lookup[valueB]

		if nodeB == nil {
			nodeB = &node{valueB, []*node{}}
		}

		valueA := orbitPair[0]
		nodeA := lookup[valueA]

		if nodeA == nil {
			nodeA = &node{valueA, []*node{nodeB}}
		} else {
			nodeA.children = append(nodeA.children, nodeB)
		}

		lookup[valueA] = nodeA
		lookup[valueB] = nodeB
	}

	fmt.Println(lookup["ZZZ"])
	fmt.Println(lookup["ZZZ"].children[0])

	rootNode := node{"a", make([]*node, 0)}

	return &rootNode
}

func process(root *node) int {
	return 1
}
