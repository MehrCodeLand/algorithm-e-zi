# Selection sort in python

def selection_sort(arr):
    n = len(arr)  # Length of the array
    for i in range(n):
        # Assume the current index as the index of the minimum element
        min_index = i
        # Iterate through the unsorted part of the array
        for j in range(i+1, n):
            # If we find an element smaller than the current minimum, update the minimum index
            if arr[j] < arr[min_index]:
                min_index = j
        # Swap the minimum element with the first element of the unsorted part
        arr[i], arr[min_index] = arr[min_index], arr[i]
    return arr


#Example
data=[-2,45,0,11,-9]
print("Sorted Array in Ascending Order: ")
print(selection_sort(data))



# This code is contributed by Erfan Shafiee