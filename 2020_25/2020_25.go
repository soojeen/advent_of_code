package main

import "fmt"

type publicKeys [2]int

func main() {
	publicKeys := publicKeys{11404017, 13768789}

	resultA := findEncryptionKey(publicKeys)

	fmt.Println("a:", resultA)
}

const subject = 7
const initialValue = 1
const divisor = 20201227

func findEncryptionKey(input publicKeys) int {
	loopSizeA := findLoopSize(input[0])
	loopSizeB := findLoopSize(input[1])

	keyA := transform(input[0], loopSizeB)
	keyB := transform(input[1], loopSizeA)

	if keyA == keyB {
		return keyA
	}

	return 0
}

func transform(subject int, loopSize int) int {
	value := initialValue

	for i := 0; i < loopSize; i++ {
		value = (value * subject) % divisor
	}

	return value
}

func findLoopSize(input int) int {
	value := initialValue
	loopSize := 1

	for {
		value = (value * subject) % divisor

		if value == input {
			return loopSize
		}

		loopSize++
	}
}
