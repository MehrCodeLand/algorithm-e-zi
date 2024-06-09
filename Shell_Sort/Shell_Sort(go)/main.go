// Shell sort in go
package main

import (
	"fmt"
)

// shellSort sorts an array of integers using the Shell Sort algorithm.
// It starts by sorting pairs of elements far apart from each other, then progressively reducing the gap between elements to be compared.
func shellSort(arr []int) {
	n := len(arr) // Get the length of the array

	// Start with a big gap, then reduce the gap
	for gap := n / 2; gap > 0; gap /= 2 {
		// Perform a gapped insertion sort for this gap size.
		for i := gap; i < n; i++ {
			// Save the current element and create a hole at position i
			temp := arr[i]

			// Shift earlier gap-sorted elements up until the correct location for arr[i] is found
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			// Put temp (the original arr[i]) in its correct location
			arr[j] = temp
		}
	}
}

func main() {
	// Example array
	arr := []int{12, 34, 54, 2, 3}

	fmt.Println("Unsorted array:", arr)

	// Call shellSort to sort the array
	shellSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", arr)
}

// This code is contributed by Erfan Shafiee