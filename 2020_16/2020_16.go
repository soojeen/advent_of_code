package main

import "fmt"
import "log"
import "regexp"
import "strconv"
import "strings"
import "advent_of_code/utils"

type ticketFields map[string][4]int

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
			values := getValuesForIndex(validTickets, i)
			isValid := true

			for _, value := range values {
				isInRangeA := value >= fieldRanges[0] && value <= fieldRanges[1]
				isInRangeB := value >= fieldRanges[2] && value <= fieldRanges[3]

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

	rawTickets := strings.Split(parts[2], "\n")
	tickets := make([][]int, len(rawTickets)-1)

	for i, rawTicket := range rawTickets[1:] {
		ticket := parseTicket(rawTicket)
		tickets[i] = ticket
	}

	return puzzleInput{ticketFields, ticket, tickets}
}

func parseTicketFields(input string) ticketFields {
	labelRe := regexp.MustCompile(`(.*):`)
	rangesRe := regexp.MustCompile(`(\d*)-(\d*) or (\d*)-(\d*)`)

	fields := strings.Split(input, "\n")
	ticketFields := make(ticketFields, len(fields))

	for _, field := range fields {
		label := labelRe.FindStringSubmatch(field)[1]
		rawRanges := rangesRe.FindStringSubmatch(field)

		validRanges := [4]int{}
		for i, rawRange := range rawRanges[1:] {
			value, _ := strconv.Atoi(rawRange)
			validRanges[i] = value
		}

		ticketFields[label] = validRanges
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

		if value > 0 {
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
			isRangeA := value >= fieldRanges[0] && value <= fieldRanges[1]
			isRangeB := value >= fieldRanges[2] && value <= fieldRanges[3]

			if isRangeA || isRangeB {
				isValid = true
				break
			}
		}

		if !isValid {
			return value
		}
	}

	return -1
}

func getValuesForIndex(input [][]int, index int) []int {
	values := make([]int, len(input))

	for j, ticket := range input {
		values[j] = ticket[index]
	}

	return values
}
