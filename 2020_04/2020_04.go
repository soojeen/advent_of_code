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

	resultA := countValid(input)
	resultB := countValidValid(input)

	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func parseInput(input string) []passport {
	passports := strings.Split(input, "\n\n")

	result := make([]passport, len(passports))

	for i, rawPassport := range passports {
		passportLine := strings.Replace(rawPassport, "\n", " ", -1)
		passportParts := strings.Split(passportLine, " ")

		passport := passport{}
		for _, passportPart := range passportParts {
			passportKeyValue := strings.Split(passportPart, ":")
			passport[passportKeyValue[0]] = passportKeyValue[1]
		}

		result[i] = passport
	}

	return result
}

// cid optional
func isValidPassport(input passport) bool {
	var requiredKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	isValid := true

	for _, key := range requiredKeys {
		value := input[key]

		if value == "" {
			isValid = false
			break
		}
	}

	return isValid
}

// cid optional
func isValidValidPassport(input passport) bool {
	isValid := true

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byrRegex := regexp.MustCompile(`^19[2-9]\d$|^200[0-2]$`)
	byrValid := byrRegex.MatchString(input["byr"])
	if !byrValid {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyrRegex := regexp.MustCompile(`^201\d$|^2020$`)
	iyrValid := iyrRegex.MatchString(input["iyr"])
	if !iyrValid {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyrRegex := regexp.MustCompile(`^202\d$|^2030$`)
	eyrValid := eyrRegex.MatchString(input["eyr"])
	if !eyrValid {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	// 		If cm, the number must be at least 150 and at most 193.
	// 		If in, the number must be at least 59 and at most 76.
	hgtRegex := regexp.MustCompile(`^1([5-8]\d|9[0-3])cm$|^(59|6\d|7[0-6])in$`)
	hgtValid := hgtRegex.MatchString(input["hgt"])
	if !hgtValid {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	hclValid := hclRegex.MatchString(input["hcl"])
	if !hclValid {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eclRegex := regexp.MustCompile(`^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$`)
	eclValid := eclRegex.MatchString(input["ecl"])
	if !eclValid {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pidRegex := regexp.MustCompile(`^\d{9}$`)
	pidValid := pidRegex.MatchString(input["pid"])
	if !pidValid {
		return false
	}

	return isValid
}

func countValid(input []passport) int {
	validCount := 0

	for _, passport := range input {
		isValid := isValidPassport(passport)

		if isValid {
			validCount++
		}
	}

	return validCount
}

func countValidValid(input []passport) int {
	validCount := 0

	for _, passport := range input {
		isValid := isValidValidPassport(passport)

		if isValid {
			validCount++
		}
	}

	return validCount
}
