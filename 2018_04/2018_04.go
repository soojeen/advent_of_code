package main

import "fmt"
import "log"

import "strconv"
import "sort"
import "regexp"
import "strings"
import "advent_of_code/utils"

type guardSleep struct {
	id         int
	sleepStart int
	wakeStart  int
	sleepTime  int
}

func main() {
	rawInput, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := parseGuardSleeps(strings.Split(rawInput, "\n"))

	resultA := mostSleepTimeGuard(input)
	resultB := mostFrequentSleepMinute(input)
	fmt.Println("a:", resultA)
	fmt.Println("b:", resultB)
}

func mostSleepTimeGuard(input []guardSleep) int {
	var totalSleepMax int
	var maxGuardID int
	totalGuardSleepTime := make(map[int]int)

	for _, guardSleep := range input {
		totalGuardSleepTime[guardSleep.id] += guardSleep.sleepTime

		currentSleepTotal := totalGuardSleepTime[guardSleep.id]
		if currentSleepTotal > totalSleepMax {
			totalSleepMax = currentSleepTotal
			maxGuardID = guardSleep.id
		}
	}

	var minuteFrequency [60]int
	var minuteCountMax int
	var minuteMax int

	for _, guardSleep := range input {
		if guardSleep.id == maxGuardID {
			for i := guardSleep.sleepStart; i < guardSleep.wakeStart; i++ {
				minuteFrequency[i]++

				if minuteFrequency[i] > minuteCountMax {
					minuteCountMax = minuteFrequency[i]
					minuteMax = i
				}
			}
		}
	}

	return maxGuardID * minuteMax
}

func mostFrequentSleepMinute(input []guardSleep) int {
	var minuteCountMax int
	var minuteMax int
	var guardID int
	guardMinuteFrequency := make(map[int][60]int)

	for _, guardSleep := range input {
		minuteFrequency := guardMinuteFrequency[guardSleep.id]

		for i := guardSleep.sleepStart; i < guardSleep.wakeStart; i++ {
			minuteFrequency[i]++

			if minuteFrequency[i] > minuteCountMax {
				minuteCountMax = minuteFrequency[i]
				minuteMax = i
				guardID = guardSleep.id
			}
		}

		guardMinuteFrequency[guardSleep.id] = minuteFrequency
	}

	return minuteMax * guardID
}

func parseGuardSleeps(input []string) []guardSleep {
	var guardSleeps []guardSleep
	var currentGuardID int
	var currentSleepStart int

	sort.Strings(input)

	for _, inputLine := range input {
		guardIDMatch, guardID := parseGuardID(inputLine)
		if guardIDMatch {
			currentGuardID = guardID
			continue
		}

		sleepMatch, sleepStart := parseTime(inputLine, `falls`)
		if sleepMatch {
			currentSleepStart = sleepStart
			continue
		}

		wakeMatch, wakeStart := parseTime(inputLine, `wakes`)
		if wakeMatch {
			sleepTime := wakeStart - currentSleepStart
			newGuardSleep := guardSleep{currentGuardID, currentSleepStart, wakeStart, sleepTime}
			guardSleeps = append(guardSleeps, newGuardSleep)
		}
	}

	return guardSleeps
}

func parseGuardID(input string) (bool, int) {
	guardIDRegexp := regexp.MustCompile(`#\d*`)

	if !guardIDRegexp.MatchString(input) {
		return false, 0
	}

	guardID, _ := strconv.Atoi(strings.Trim(guardIDRegexp.FindString(input), "#"))

	return true, guardID
}

func parseTime(input string, regExpString string) (bool, int) {
	regex := regexp.MustCompile(regExpString)
	timeRegexp := regexp.MustCompile(`:\d\d`)

	if !regex.MatchString(input) {
		return false, 0
	}

	time, _ := strconv.Atoi(strings.Trim(timeRegexp.FindString(input), ":"))

	return true, time
}
