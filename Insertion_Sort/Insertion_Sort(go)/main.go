//Insertion Sort in go
package main

import (
	"fmt"
)

// insertionSort sorts an array of integers using the Insertion Sort algorithm.
// It builds the sorted array one item at a time, with the array elements being
// rearranged in place.
func insertionSort(arr []int) {
	// Iterate over each element in the array starting from the second element
	for i := 1; i < len(arr); i++ {
		key := arr[i] // Current element to be inserted in the sorted portion
		j := i - 1    // Index of the last element in the sorted portion

		// Move elements of the sorted portion that are greater than key one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key // Insert the key at the correct position
	}
}

func main() {
	// Example array
	arr := []int{12, 11, 13, 5, 6}

	fmt.Println("Unsorted array:", arr)

	// Call insertionSort to sort the array
	insertionSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", arr)
}

// This code is contributed by Erfan Shafiee