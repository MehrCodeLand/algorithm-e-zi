// Selection sort in go
package main

import (
	"fmt"
)

// selectionSort sorts an array of integers using the Selection Sort algorithm.
// It repeatedly finds the minimum element from the unsorted part and moves it to the sorted part.
func selectionSort(arr []int) {
	n := len(arr) // Get the length of the array

	// Loop over each element in the array
	for i := 0; i < n-1; i++ {
		// Assume the minimum is the first element
		minIdx := i

		// Loop through the unsorted part of the array
		for j := i + 1; j < n; j++ {
			// If we find a smaller element, update minIdx
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		// Swap the found minimum element with the first element of the unsorted part
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	// Example array
	arr := []int{64, 25, 12, 22, 11}

	fmt.Println("Unsorted array:", arr)

	// Call selectionSort to sort the array
	selectionSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", arr)
}

// This code is contributed by Erfan Shafiee