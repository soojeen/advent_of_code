package main

import "fmt"
import "log"
import "strconv"
import "strings"
import "advent_of_code/utils"

func main() {
  rawInput, err := utils.ReadInput("rawInput.txt")
  if err != nil {
    log.Fatal(err)
  }

  input, err := parseInput(rawInput)
  if err != nil {
    log.Fatal(err)
  }

  result := sumChanges(input)

  fmt.Println(result)
}

func parseInput(input string) ([]int, error) {
  changes := strings.Split(input, "\n")

  new := make([]int, len(changes))
  var err error

  for i, changeStr := range changes {
    changeInt, e := strconv.Atoi(changeStr)
    if e != nil {
      err = e
      break
    }

    new[i] = changeInt
  }

  return new, err
}

func sumChanges(changes []int) int {
  sum := 0
  for _, change := range changes {
    sum += change
  }

  return sum
}
