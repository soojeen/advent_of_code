package main

import "fmt"

type publicKeys [2]int

func main() {
	publicKeys := publicKeys{11404017, 13768789}

	resultA := findEncryptionKey(publicKeys)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

const subject = 7
const initialValue = 1
const divisor = 20201227

func findEncryptionKey(input publicKeys) int {
	value := initialValue

	loopSize := 1

	for {
		value = value * subject
		value = value % divisor
		fmt.Println("value:", value, "loop:", loopSize)

		loopSize++

		if loopSize == 100 {
			break
		}
	}

	return 0
}
