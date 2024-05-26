// Binary Search in go
package Binary_Search

import (
	"fmt"
)

// binarySearch performs binary search on a sorted array.
// It returns the index of the target value if found, otherwise returns -1.
func binarySearch(arr []int, target int) int {
	// Initialize the start and end pointers
	start, end := 0, len(arr)-1

	// Loop until the pointers cross
	for start <= end {
		// Calculate the middle index
		mid := start + (end-start)/2

		// Check if the target is present at mid
		if arr[mid] == target {
			return mid // Target found
		}

		// If target is greater, ignore the left half
		if arr[mid] < target {
			start = mid + 1
		} else {
			// If target is smaller, ignore the right half
			end = mid - 1
		}
	}

	// Target not found
	return -1
}

func main() {
	// Example sorted array
	arr := []int{2, 3, 4, 10, 40}
	target := 10

	// Call binarySearch and store the result
	result := binarySearch(arr, target)

	// Print the result
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found in the array")
	}
}

// This code is contributed by Erfan Shafiee