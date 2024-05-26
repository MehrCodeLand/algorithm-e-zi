// Quick sort in go
package main

import (
	"fmt"
)

// quickSort sorts an array of integers using the Quick Sort algorithm.
// It selects a 'pivot' element and partitions the array into two sub-arrays,
// according to whether the elements are less than or greater than the pivot.
func quickSort(arr []int) []int {
	// Base case: arrays with 0 or 1 element are already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Select the pivot element (here we choose the last element)
	pivot := arr[len(arr)-1]

	// Partition the array into two sub-arrays: less than and greater than the pivot
	var less, greater []int
	for _, v := range arr[:len(arr)-1] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	// Recursively sort the sub-arrays and concatenate them with the pivot
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	// Example array
	arr := []int{10, 7, 8, 9, 1, 5}

	fmt.Println("Unsorted array:", arr)

	// Call quickSort to sort the array
	sortedArr := quickSort(arr)

	// Print the sorted array
	fmt.Println("Sorted array:", sortedArr)
}

// This code is contributed by Erfan Shafiee