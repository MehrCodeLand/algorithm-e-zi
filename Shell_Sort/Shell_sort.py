## Shell sort in python

def shell_sort(arr):
    n = len(arr)
    gap = n // 2  # Initialize gap as half of the array length

    while gap > 0:
        # Perform gapped insertion sort for this gap size.
        # The first gap elements a[0..gap-1] are already in gapped order
        # keep adding one more element until the entire array is gap sorted
        for i in range(gap, n):

            # add a[i] to the elements that have been gap sorted
            # save a[i] in temp and make a hole at position i
            temp = arr[i]

            # shift earlier gap-sorted elements up until the correct location for a[i] is found
            j = i
            while j >= gap and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap

            # put temp (the original a[i]) in its correct location
            arr[j] = temp

        # Reduce the gap for the next iteration
        gap //= 2

    return arr


#Example
data=[-2,45,0,11,-9]
print("Sorted Array in Ascending Order: ")
print(shell_sort(data))

# This code is contributed by Erfan Shafiee
