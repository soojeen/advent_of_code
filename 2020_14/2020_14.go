package main

import "bytes"
import "fmt"
import "log"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := processA(input)
	// resultB := processB(input)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func processA(input []string) int {
	mask := ""
	memory := map[string]int64{}
	maskRE := regexp.MustCompile(`mask = (.*)`)
	memRE := regexp.MustCompile(`mem\[(\d*)\] = (\d*)`)

	for _, line := range input {
		maskMatch := maskRE.FindStringSubmatch(line)
		if maskMatch != nil {
			mask = maskMatch[1]
			continue
		}

		memMatch := memRE.FindStringSubmatch(line)
		key := memMatch[1]
		value, _ := strconv.Atoi(memMatch[2])
		binary := strconv.FormatInt(int64(value), 2)
		binary = padLeft(binary, "0", len(mask))
		maskedValue, _ := strconv.ParseInt(binary, 2, 0)
		memory[key] = maskedValue
	}

	return 0
}

func padLeft(input string, padding string, length int) string {
	var b strings.Builder

	padLength := length - len(input)

	for i := 0; i < padLength; i++ {
		b.WriteString(padding)
	}

	b.WriteString(input)

	return b.String()
}

func bitMask(value string, mask string) string {
	x := "X"[0]
	var b bytes.Buffer

	for i := 0; i < len(value); i++ {
		bitMask := mask[i]

		if bitMask != x {
			b.WriteByte(value[i])
		}

		b.WriteByte(bitMask)
	}

	return b.String()
}
