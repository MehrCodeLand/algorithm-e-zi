// Bubble sort in go
package main

import (
	"fmt"
)

// bubbleSort sorts an array of integers using the Bubble Sort algorithm.
// It iterates over the list, repeatedly swapping adjacent elements if they are in the wrong order.
func bubbleSort(arr []int) {
	n := len(arr) // Get the length of the array

	// Perform n-1 passes over the array
	for i := 0; i < n-1; i++ {
		// Flag to detect any swap in this pass
		swapped := false

		// Perform the inner loop, comparing adjacent elements
		for j := 0; j < n-i-1; j++ {
			// If the current element is greater than the next element, swap them
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true // Mark that a swap occurred
			}
		}

		// If no elements were swapped in this pass, the array is sorted
		if !swapped {
			break
		}
	}
}

func main() {
	// Example array
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Unsorted array:", arr)

	// Call bubbleSort to sort the array
	bubbleSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", arr)
}


// This code is contributed by Erfan Shafiee