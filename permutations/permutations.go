package permutations

func nextPermutation(offsets []int) {
	for i := len(offsets) - 1; i >= 0; i-- {
		if i == 0 || offsets[i] < len(offsets)-i-1 {
			offsets[i]++
			return
		}

		offsets[i] = 0
	}
}

func getPermutation(base, offsets []int) []int {
	result := make([]int, len(base))
	copy(result, base)

	for i, v := range offsets {
		result[i], result[i+v] = result[i+v], result[i]
	}

	return result
}

// IteratePermutation - iterate through all permutations
func IteratePermutation(base []int, f func([]int)) {
	offsets := make([]int, len(base))

	for offsets[0] < len(offsets) {
		permutation := getPermutation(base, offsets)
		f(permutation)

		nextPermutation(offsets)
	}
}
