package main

import "fmt"
import "log"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

type ticketFields map[string][2][2]int

type puzzleInput struct {
	ticketFields ticketFields
	ticket       []int
	tickets      [][]int
}

func (p *puzzleInput) getValidTickets() [][]int {
	validTickets := [][]int{}

	for _, ticket := range p.tickets {
		invalid := invalidValue(ticket, p.ticketFields)

		if invalid == -1 {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func (p *puzzleInput) getFieldPossibles() map[string][]int {
	tracker := map[string][]int{}
	validTickets := p.getValidTickets()

	for label, fieldRanges := range p.ticketFields {
		valids := []int{}

		for i := 0; i < len(p.ticket); i++ {
			values := make([]int, len(validTickets))

			for j, ticket := range validTickets {
				values[j] = ticket[i]
			}

			isValid := true

			for _, value := range values {
				isInRangeA := value >= fieldRanges[0][0] && value <= fieldRanges[0][1]
				isInRangeB := value >= fieldRanges[1][0] && value <= fieldRanges[1][1]

				if !isInRangeA && !isInRangeB {
					isValid = false
					break
				}
			}

			if isValid {
				valids = append(valids, i)
			}
		}

		tracker[label] = valids
	}

	return tracker
}

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

func parseInput(input string) puzzleInput {
	parts := strings.Split(input, "\n\n")

	ticketFields := parseTicketFields(parts[0])

	ticketParts := strings.Split(parts[1], "\n")
	ticket := parseTicket(ticketParts[1])

	ticketsParts := strings.Split(parts[2], "\n")
	tickets := make([][]int, len(ticketsParts)-1)
	for i, ticket := range ticketsParts[1:] {
		ticket := parseTicket(ticket)
		tickets[i] = ticket
	}

	return puzzleInput{ticketFields, ticket, tickets}
}

func parseTicketFields(input string) ticketFields {
	labelRe := regexp.MustCompile(`(.*):`)
	rangesRe := regexp.MustCompile(`\d*-\d*`)

	fields := strings.Split(input, "\n")
	ticketFields := make(ticketFields, len(fields))

	for _, field := range fields {
		label := labelRe.FindStringSubmatch(field)
		rawRanges := rangesRe.FindAllString(field, -1)

		validRanges := [2][2]int{}
		for i, rawRange := range rawRanges {
			rangeParts := strings.Split(rawRange, "-")
			minMaxRange := [2]int{}

			for j, rangePart := range rangeParts {
				value, _ := strconv.Atoi(rangePart)
				minMaxRange[j] = value
			}

			validRanges[i] = minMaxRange
		}

		ticketFields[label[1]] = validRanges
	}

	return ticketFields
}

func parseTicket(input string) []int {
	ticketValues := strings.Split(input, ",")
	ticket := make([]int, len(ticketValues))

	for i, ticketValue := range ticketValues {
		value, _ := strconv.Atoi(ticketValue)
		ticket[i] = value
	}

	return ticket
}

func processA(input puzzleInput) int {
	result := 0

	for _, ticket := range input.tickets {
		value := invalidValue(ticket, input.ticketFields)

		if value >= 0 {
			result += value
		}
	}

	return result
}

func processB(input puzzleInput) int {
	fieldPossibles := input.getFieldPossibles()
	final := map[string]int{}
	used := map[int]bool{}

	for i := 0; ; i++ {
		for label, possibles := range fieldPossibles {
			if len(possibles) == i+1 {
				for _, possible := range possibles {
					if !used[possible] {
						final[label] = possible
						used[possible] = true
					}
				}
			}
		}

		if i == 20 {
			break
		}
	}

	result := 1
	labelRe := regexp.MustCompile(`departure`)

	for label, index := range final {
		if labelRe.MatchString(label) {
			result *= input.ticket[index]
		}
	}

	return result
}

func invalidValue(input []int, ticketFields ticketFields) int {
	for _, value := range input {
		isValid := false

		for _, fieldRanges := range ticketFields {
			for _, validRange := range fieldRanges {
				if value >= validRange[0] && value <= validRange[1] {
					isValid = true
					break
				}
			}
		}

		if !isValid {
			return value
		}
	}

	return -1
}
