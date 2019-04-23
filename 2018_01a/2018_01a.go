package main

import "fmt"
import "advent_of_code/utils"

func main() {

  input, err := utils.ReadInput("input.txt")

  if err != nil {
    fmt.Println("File Read Error")
  }

  fmt.Println(input)
}
