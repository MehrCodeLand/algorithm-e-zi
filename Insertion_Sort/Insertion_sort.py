# Insertion sort in python

def insertion_sort(arr):
    # Traverse through 1 to len(arr)
    for i in range(1, len(arr)):
        # Select the element to be inserted
        key = arr[i]
        
        # Initialize j to the index before i
        j = i - 1
        
        # Move elements of arr[0..i-1], that are greater than key,
        # to one position ahead of their current position
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        
        # Insert the key element after the element just smaller than it
        arr[j + 1] = key
    
    return arr


#Example
data=[-2,45,0,11,-9]
print("Sorted Array in Ascending Order: ")
print(insertion_sort(data))


# This code is contributed by Erfan Shafiee
