# Bubble sort in Python

#bubble sort function
def bubbleSort(array,size):
    #loop to access each array element
    for step in range(size):
        for i in range(size-step-1):
            #compare two adjacent elements
            #change > to < to sort in descending order
            if (array[i]>array[i+1]):
                #swapping elements if elements 
                #are not in the intended order
                temp=array[i]
                array[i]=array[i+1]
                array[i+1]=temp
    return array

#Eample
data=[-2,45,0,11,-9]
#find array's length
size=len(data)

print("Sorted Array in Ascending Order: ")
print(bubbleSort(data,size))

# This code is contributed by Erfan Shafiee