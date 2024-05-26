// Merge sort in go
package main

import (
	"fmt"
)

// mergeSort sorts an array of integers using the Merge Sort algorithm.
// It recursively divides the array into halves, sorts each half,
// and then merges the sorted halves.
func mergeSort(arr []int) []int {
	// If the array has one or zero elements, it's already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Find the middle point to divide the array into two halves
	mid := len(arr) / 2

	// Recursively sort the first and second halves
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

// merge combines two sorted arrays into one sorted array.
func merge(left, right []int) []int {
	// Create a result array to hold the merged sorted elements
	result := make([]int, 0, len(left)+len(right))

	// Indices to track the current position in left and right arrays
	i, j := 0, 0

	// Loop until we reach the end of either left or right
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			// If the current element in left is smaller, add it to the result
			result = append(result, left[i])
			i++
		} else {
			// If the current element in right is smaller or equal, add it to the result
			result = append(result, right[j])
			j++
		}
	}

	// Append any remaining elements in left (if any)
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append any remaining elements in right (if any)
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	// Return the merged sorted array
	return result
}

func main() {
	// Example array
	arr := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println("Unsorted array:", arr)

	// Call mergeSort to sort the array
	sortedArr := mergeSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", sortedArr)
}

// This code is contributed by Erfan Shafiee