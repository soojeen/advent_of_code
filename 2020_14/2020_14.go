package main

import "bytes"
import "fmt"
import "log"
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

		masked := bitMask(decoder.valueB, mask)
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

func bitMask(value string, mask string) string {
	x := "X"[0]
	var b bytes.Buffer

	for i := 0; i < len(value); i++ {
		bitMask := mask[i]

		if bitMask == x {
			b.WriteByte(value[i])
		} else {
			b.WriteByte(bitMask)
		}
	}

	return b.String()
}

func bitMaskB(value string, mask string) []string {
	result := []string{}
	string := "X01"
	x := string[0]
	z := string[1]
	o := string[2]
	var b bytes.Buffer

	for i := 0; i < len(value); i++ {
		bitMask := mask[i]

		if bitMask == o {
			b.WriteByte(o)
		} else if bitMask == z {
			b.WriteByte(value[i])
		} else if bitMask == x {
			// TODO branch
		}
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
