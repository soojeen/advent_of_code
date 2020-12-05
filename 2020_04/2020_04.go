package main

import "fmt"
import "log"
import "regexp"
import "strings"
import "advent_of_code/utils"

type passport string
type validator []*regexp.Regexp

func (pp passport) isValid(validator validator) bool {
	isValidPassport := true

	for _, re := range validator {
		if !re.MatchString(string(pp)) {
			isValidPassport = false
			break
		}
	}

	return isValidPassport
}

func main() {
	rawInput, readError := utils.ReadInput("input.txt")
	if readError != nil {
		log.Fatal(readError)
	}

	input := parseInput(rawInput)

	resultA := countValid(input, requiredKeysRegex)
	resultB := countValid(input, passportRegex)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []passport {
	passports := strings.Split(input, "\n\n")

	result := make([]passport, len(passports))

	for i, passportLine := range passports {
		// key/value pairs in passport can be separated by space or new line.
		// format consistently with space.
		result[i] = passport(strings.Replace(passportLine, "\n", " ", -1))
	}

	return result
}

func countValid(input []passport, validator validator) int {
	validCount := 0

	for _, passport := range input {
		isValid := passport.isValid(validator)

		if isValid {
			validCount++
		}
	}

	return validCount
}

var requiredKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var requiredKeysRegex = createRequiredKeysRegex()

func createRequiredKeysRegex() validator {
	result := make(validator, len(requiredKeys))

	for i, key := range requiredKeys {
		result[i] = regexp.MustCompile(key)
	}

	return result
}

var passportRegex = validator{
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	regexp.MustCompile(`\bbyr:(19[2-9]\d|200[0-2])\b`),

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	regexp.MustCompile(`\biyr:(201\d|2020)\b`),

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	regexp.MustCompile(`\beyr:(202\d|2030)\b`),

	// hgt (Height) - a number followed by either cm or in:
	// 		If cm, the number must be at least 150 and at most 193.
	// 		If in, the number must be at least 59 and at most 76.
	regexp.MustCompile(`\bhgt:(1([5-8]\d|9[0-3])cm|(59|6\d|7[0-6])in)\b`),

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	regexp.MustCompile(`\bhcl:#[0-9a-f]{6}\b`),

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	regexp.MustCompile(`\becl:(amb|blu|brn|gry|grn|hzl|oth)\b`),

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	regexp.MustCompile(`\b\d{9}\b`),
}
