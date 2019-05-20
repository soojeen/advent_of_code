package main

import "fmt"
import "log"
import "sort"
import "strings"
import "advent_of_code/utils"

type reqsGraph map[string][]string

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reqs := parseReqs(rawInput)

	resultA := correctOrder(reqs)
	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseReqs(input string) reqsGraph {
	rawReqs := strings.Split(input, "\n")
	reqs := make(reqsGraph)

	for _, rawRequirement := range rawReqs {
		var nodeValue string
		var requirement string
		scanString := "Step %s must be finished before step %s can begin."

		_, error := fmt.Sscanf(rawRequirement, scanString, &requirement, &nodeValue)
		if error != nil {
			panic(error)
		}

		reqs[nodeValue] = append(reqs[nodeValue], requirement)

		if len(reqs[requirement]) <= 0 {
			reqs[requirement] = []string{}
		}
	}

	return reqs
}

type completeMap map[string]bool

func correctOrder(input reqsGraph) string {
	complete := make(completeMap)

	currentNodeValue := findRoot(input)
	complete[currentNodeValue] = true

	for i := 0; i < len(input); i++ {
		possibleNextNodeValues := []string{}

		for nodeValue, reqs := range input {
			if containsValue(reqs, currentNodeValue) && allReqsComplete(reqs, complete) {
				possibleNextNodeValues = append(possibleNextNodeValues, nodeValue)
			}
		}

		sort.Strings(possibleNextNodeValues)

		if len(possibleNextNodeValues) == 0 {
			fmt.Println("poss:", possibleNextNodeValues)
			fmt.Println("curr:", currentNodeValue)
			fmt.Println("reqs:", input)
			fmt.Println("comp:", complete)
		}
		currentNodeValue = possibleNextNodeValues[0]
		complete[currentNodeValue] = true
	}

	fmt.Println("root:", currentNodeValue)

	return "a"
}

func findRoot(input reqsGraph) string {
	var possibleRoots []string

	for nodeValue, reqs := range input {
		if len(reqs) == 0 {
			possibleRoots = append(possibleRoots, nodeValue)
		}
	}

	sort.Strings(possibleRoots)

	return possibleRoots[0]
}

func containsValue(list []string, value string) bool {
	for _, listValue := range list {
		if listValue == value {
			return true
		}
	}

	return false
}

func allReqsComplete(reqs []string, complete completeMap) bool {
	for _, req := range reqs {
		if !complete[req] {
			return false
		}
	}

	return true
}
