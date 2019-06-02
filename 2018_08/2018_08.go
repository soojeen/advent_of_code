package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

type treePartial struct {
	metaDataSum int
	value       int
	subTree     []int
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseInput(rawInput)
	result := parseTreeSum(input)

	fmt.Println("a:", result.metaDataSum)
	fmt.Println("b:", result.value)
}

func parseInput(input string) []int {
	values := strings.Split(input, " ")
	result := make([]int, len(values))

	for i, value := range values {
		iValue, _ := strconv.Atoi(value)
		result[i] = iValue
	}

	return result
}

func parseTreeSum(tree []int) treePartial {
	childrenLength, metaDataLength, subTree := tree[0], tree[1], tree[2:]
	metaDataSum := 0
	values := []int{}

	for i := 0; i < childrenLength; i++ {
		childResult := parseTreeSum(subTree)
		metaDataSum += childResult.metaDataSum
		subTree = childResult.subTree

		values = append(values, childResult.value)
	}

	metaDatas := subTree[:metaDataLength]
	value := sum(metaDatas)
	metaDataSum += value

	if childrenLength == 0 {
		return treePartial{metaDataSum, value, subTree[metaDataLength:]}
	}

	valueSubtotal := 0

	for _, metaDataValue := range metaDatas {
		valuesIndex := metaDataValue - 1

		if valuesIndex < len(values) {
			valueSubtotal += values[valuesIndex]
		}
	}

	return treePartial{metaDataSum, valueSubtotal, subTree[metaDataLength:]}
}

func sum(addends []int) int {
	result := 0

	for _, addend := range addends {
		result += addend
	}

	return result
}
