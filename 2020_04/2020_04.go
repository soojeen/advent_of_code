package main

import "fmt"
import "log"
import "regexp"
import "strings"
import "advent_of_code/utils"

type passport = map[string]string

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := countValid(input, isValidPassport)
	resultB := countValid(input, isValidValidPassport)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []string {
	passports := strings.Split(input, "\n\n")

	result := make([]string, len(passports))

	for i, rawPassport := range passports {
		passport := strings.Replace(rawPassport, "\n", " ", -1)

		result[i] = passport
	}

	return result
}

func countValid(input []string, validator func(string) bool) int {
	validCount := 0

	for _, passport := range input {
		isValid := validator(passport)

		if isValid {
			validCount++
		}
	}

	return validCount
}

func isValidPassport(input string) bool {
	var requiredKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	isValid := true

	for _, key := range requiredKeys {
		re := regexp.MustCompile(key)
		match := re.MatchString(input)

		if !match {
			isValid = false
			break
		}
	}

	return isValid
}

func isValidValidPassport(input string) bool {
	isValid := true

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byrRegex := regexp.MustCompile(`\bbyr:(19[2-9]\d|200[0-2])\b`)
	byrValid := byrRegex.MatchString(input)
	if !byrValid {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyrRegex := regexp.MustCompile(`\biyr:(201\d|2020)\b`)
	iyrValid := iyrRegex.MatchString(input)
	if !iyrValid {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyrRegex := regexp.MustCompile(`\beyr:(202\d|2030)\b`)
	eyrValid := eyrRegex.MatchString(input)
	if !eyrValid {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// 		If cm, the number must be at least 150 and at most 193.
	// 		If in, the number must be at least 59 and at most 76.
	hgtRegex := regexp.MustCompile(`\bhgt:(1([5-8]\d|9[0-3])cm|(59|6\d|7[0-6])in)\b`)
	hgtValid := hgtRegex.MatchString(input)
	if !hgtValid {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hclRegex := regexp.MustCompile(`\bhcl:#[0-9a-f]{6}\b`)
	hclValid := hclRegex.MatchString(input)
	if !hclValid {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eclRegex := regexp.MustCompile(`\becl:(amb|blu|brn|gry|grn|hzl|oth)\b`)
	eclValid := eclRegex.MatchString(input)
	if !eclValid {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pidRegex := regexp.MustCompile(`\b\d{9}\b`)
	pidValid := pidRegex.MatchString(input)
	if !pidValid {
		return false
	}

	return isValid
}
