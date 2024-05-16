# Quick sort in python

def quick_sort(arr):
    # Base case: if the list has 1 or fewer elements, it's already sorted
    if len(arr) <= 1:
        return arr
    
    # Selecting the pivot element (here, we choose the middle element)
    pivot = arr[len(arr) // 2]
    
    # Dividing the list into three parts: elements less than pivot, equal to pivot, and greater than pivot
    left = [x for x in arr if x < pivot]      # Elements smaller than the pivot
    middle = [x for x in arr if x == pivot]   # Elements equal to the pivot
    right = [x for x in arr if x > pivot]     # Elements greater than the pivot

    # Recursively sorting the left and right partitions
    return quick_sort(left) + middle + quick_sort(right)



data=[-2,45,0,11,-9]
print("Sorted Array in Ascending Order: ")
print(quick_sort(data))


# This code is contributed by Erfan Shafiee