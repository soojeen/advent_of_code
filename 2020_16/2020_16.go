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
type fieldPossibles struct {
	label     string
	possibles []int
}

func (p *puzzleInput) getValidTickets() [][]int {
	validTickets := [][]int{}

	for _, ticket := range p.tickets {
		invalid := invalidTicket(ticket, p.ticketFields)

		if !invalid {
			validTickets = append(validTickets, ticket)
		}
	}

	return validTickets
}

func (p *puzzleInput) getFieldPossibles() []fieldPossibles {
	result := make([]fieldPossibles, len(p.ticketFields))

	validTickets := p.getValidTickets()

	for label, fieldRanges := range p.ticketFields {
		valids := []int{}
		count := 0

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

		result[len(valids)-1] = fieldPossibles{label, valids}
		count++
	}

	return result
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
		for _, value := range ticket {
			invalid := invalidValue(value, input.ticketFields)

			if invalid > 0 {
				result += invalid
			}
		}
	}

	return result
}

func processB(input puzzleInput) int {
	labelRe := regexp.MustCompile(`departure`)
	used := map[int]bool{}
	result := 1

	orderedPossibles := input.getFieldPossibles()

	for _, fieldPossibles := range orderedPossibles {
		for _, possible := range fieldPossibles.possibles {
			if used[possible] {
				continue
			}

			if labelRe.MatchString(fieldPossibles.label) {
				result *= input.ticket[possible]
			}

			used[possible] = true
		}
	}

	return result
}

func invalidValue(input int, ticketFields ticketFields) int {
	for _, fieldRanges := range ticketFields {
		isRangeA := input >= fieldRanges[0] && input <= fieldRanges[1]
		isRangeB := input >= fieldRanges[2] && input <= fieldRanges[3]

		if isRangeA || isRangeB {
			return -1
		}
	}

	return input
}

func invalidTicket(input []int, ticketFields ticketFields) bool {
	for _, value := range input {
		invalid := invalidValue(value, ticketFields)

		if invalid != -1 {
			return true
		}
	}

	return false
}

func getValuesForIndex(input [][]int, index int) []int {
	values := make([]int, len(input))

	for j, ticket := range input {
		values[j] = ticket[index]
	}

	return values
}
