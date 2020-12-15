package main

import "bytes"
import "fmt"
import "log"
import "math"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

type decoder struct {
	mask     string
	address  int
	addressB string
	value    int
	valueB   string
}

const bitLength = 36

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := processA(input)
	resultB := processB(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []decoder {
	lines := strings.Split(input, "\n")
	result := make([]decoder, len(lines))

	maskRE := regexp.MustCompile(`mask = (.*)`)
	memRE := regexp.MustCompile(`mem\[(\d*)\] = (\d*)`)

	for i, line := range lines {
		maskMatch := maskRE.FindStringSubmatch(line)
		if maskMatch != nil {
			result[i] = decoder{maskMatch[1], 0, "", 0, ""}
			continue
		}

		memMatch := memRE.FindStringSubmatch(line)

		address, _ := strconv.Atoi(memMatch[1])
		addressB := strconv.FormatInt(int64(address), 2)
		padAddress := padLeft(addressB, "0", bitLength)

		value, _ := strconv.Atoi(memMatch[2])
		valueB := strconv.FormatInt(int64(value), 2)
		padValue := padLeft(valueB, "0", bitLength)

		result[i] = decoder{"", address, padAddress, value, padValue}
	}

	return result
}

func processA(input []decoder) int {
	mask := ""
	memory := map[int]int{}

	for _, decoder := range input {
		if decoder.mask != "" {
			mask = decoder.mask
			continue
		}

		masked := bitMaskA(decoder.valueB, mask)
		maskedValue, _ := strconv.ParseInt(masked, 2, 0)
		memory[decoder.address] = int(maskedValue)
	}

	return sumAll(memory)
}

func processB(input []decoder) int {
	mask := ""
	memory := map[int]int{}

	for _, decoder := range input {
		if decoder.mask != "" {
			mask = decoder.mask
			continue
		}

		addresses := bitMaskB(decoder.addressB, mask)
		for _, maskedAddress := range addresses {
			address, _ := strconv.ParseInt(maskedAddress, 2, 0)
			memory[int(address)] = decoder.value
		}
	}

	return sumAll(memory)
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

func bitMaskA(value string, mask string) string {
	x := "X"
	var b bytes.Buffer

	for i := 0; i < len(value); i++ {
		bit := mask[i]

		if bit == x[0] {
			b.WriteByte(value[i])
		} else {
			b.WriteByte(bit)
		}
	}

	return b.String()
}

func bitMaskB(value string, mask string) []string {
	zero := "0"
	one := "1"
	x := "X"

	var b bytes.Buffer
	result := []string{}
	xCount := 0

	for i := 0; i < len(mask); i++ {
		bit := mask[i]

		if bit == one[0] {
			b.WriteByte(one[0])
		} else if bit == zero[0] {
			b.WriteByte(value[i])
		} else if bit == x[0] {
			xCount++
			b.WriteByte(x[0])
		}
	}

	counts := math.Pow(2, float64(xCount))

	for i := 0; i < int(counts); i++ {
		varMask := b.String()
		iB := strconv.FormatInt(int64(i), 2)
		iB = padLeft(iB, zero, xCount)

		for _, char := range iB {
			varMask = strings.Replace(varMask, x, string(char), 1)
		}

		result = append(result, varMask)
	}

	return result
}

func sumAll(input map[int]int) int {
	result := 0

	for _, value := range input {
		result += value
	}

	return result
}
