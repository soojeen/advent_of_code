package main

import "fmt"
import "log"
import "sort"
import "strings"
import "advent_of_code/utils"

type reqsGraph map[string][]string
type completeMap map[string]bool
type workOrder struct {
	value    string
	timeLeft int
}

const maxWorkerCount = 5
const extraTime = 60

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reqs := parseReqs(rawInput)

	resultA := correctOrder(reqs)
	resultB := multiWorkers(reqs)
	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
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

func correctOrder(input reqsGraph) string {
	complete := make(completeMap)
	reqQueue := []string{}
	result := make([]string, len(input))

	for i := range result {
		for nodeValue, reqs := range input {
			valid := false
			if i == 0 {
				valid = len(reqs) == 0
			} else if i > 0 {
				valid = containsValue(reqs, result[i-1]) && allReqsComplete(reqs, complete)
			}

			if valid {
				reqQueue = append(reqQueue, nodeValue)
			}
		}

		sort.Strings(reqQueue)

		result[i] = reqQueue[0]
		reqQueue = reqQueue[1:]

		complete[result[i]] = true
	}

	return strings.Join(result, "")
}

func multiWorkers(input reqsGraph) int {
	complete := make(completeMap)
	result := make([]string, len(input))
	workQueue := []workOrder{}
	alphabetMap := createAlphabetMap(extraTime)
	time := 0

	for i := range result {
		for nodeValue, reqs := range input {
			valid := false
			if i == 0 {
				valid = len(reqs) == 0
			} else if i > 0 {
				valid = containsValue(reqs, result[i-1]) && allReqsComplete(reqs, complete)
			}

			if valid {
				workQueue = append(workQueue, workOrder{nodeValue, alphabetMap[nodeValue]})
			}
		}

		sort.Slice(workQueue, func(i, j int) bool {
			return workQueue[i].timeLeft < workQueue[j].timeLeft
		})

		workOrderValue, timeElapsed, newWorkQueue := popWorkOrder(workQueue)
		workQueue = newWorkQueue

		time += timeElapsed

		result[i] = workOrderValue
		complete[result[i]] = true
	}

	return time
}

func popWorkOrder(workQueue []workOrder) (string, int, []workOrder) {
	timeElapsed := workQueue[0].timeLeft

	for i, workOrder := range workQueue {
		if i < maxWorkerCount {
			workQueue[i].timeLeft = workOrder.timeLeft - timeElapsed
		}
	}

	return workQueue[0].value, timeElapsed, workQueue[1:]
}

func createAlphabetMap(extra int) map[string]int {
	rawLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters := strings.Split(rawLetters, "")
	result := make(map[string]int)

	for i, letter := range letters {
		result[letter] = i + 1 + extra
	}

	return result
}

func containsValue(list []string, value string) bool {
	if len(list) == 0 || value == "" {
		return false
	}

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
