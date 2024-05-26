# Binary Search in Python

def binary_search(arr, target):
    left = 0  # Index of the leftmost element of the array.
    right = len(arr) - 1  # Index of the rightmost element of the array.

    while left <= right:
        mid = (left + right) // 2  # Calculate the midpoint of the array.

        if arr[mid] == target:
            return mid  # Return the index of the target element.
        elif arr[mid] < target:
            left = mid + 1  # Update the left pointer.
        else:
            right = mid - 1  # Update the right pointer.

    return -1  # Return -1 if the target element is not found.


#Example
arr = [1, 3, 5, 7, 9, 11, 13, 15]
target = 7
result = binary_search(arr, target)
if result != -1:
    print(f"Element {target} is present at index {result}.")
else:
    print(f"Element {target} is not present in the array.")


# This code is contributed by Erfan Shafiee