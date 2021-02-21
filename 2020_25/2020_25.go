package main

import "fmt"

type publicKeys [2]int

func main() {
	publicKeys := publicKeys{11404017, 13768789}

	resultA := findEncryptionKey(publicKeys)

	fmt.Println("a:", resultA)
	// fmt.Println("b:", resultB)
}

func findEncryptionKey(input publicKeys) int {
	return 0
}
