package main

// FindMax returns the maximum value in a slice of integers.
// If the slice is empty, it returns 0.
func FindMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, elem := range numbers[1:] {
		if elem > max {
			max = elem
		}
	}
	return max
}

// RemoveDuplicates returns a new slice with duplicate values removed,
// preserving the original order of elements.
func RemoveDuplicates(numbers []int) []int {
	memory := map[int]bool{}
	sol := []int{}

	for _, elem := range numbers {
		if memory[elem] {
			continue
		}

		memory[elem] = true
		sol = append(sol, elem)
	}

	return sol
}

// ReverseSlice returns a new slice with elements in reverse order.
func ReverseSlice(slice []int) []int {
	sol := []int{}

	for idx := len(slice) - 1; idx >= 0; idx-- {
		sol = append(sol, slice[idx])
	}

	return sol
}

// FilterEven returns a new slice containing only the even numbers
// from the original slice.
func FilterEven(numbers []int) []int {
	sol := []int{}

	for _, elem := range numbers {
		if elem%2 == 0 {
			sol = append(sol, elem)
		}
	}

	return sol
}
