// Quick sort in C#
using System;

// function to swap elements
static void Swap(ref int a, ref int b)
{
    int t = a;
    a = b;
    b = t;
}

// function to print the array
static void PrintArray(int[] array, int size)
{
    for (int i = 0; i < size; i++)
        Console.Write(array[i] + " ");
    Console.WriteLine();
}

// function to rearrange array (find the partition point)
static int Partition(int[] array, int low, int high)
{
    // select the rightmost element as pivot
    int pivot = array[high];

    // pointer for greater element
    int i = (low - 1);

    // traverse each element of the array
    // compare them with the pivot
    for (int j = low; j < high; j++)
    {
        if (array[j] <= pivot)
        {
            // if element smaller than pivot is found
            // swap it with the greater element pointed by i
            i++;

            // swap element at i with element at j
            Swap(ref array[i], ref array[j]);
        }
    }

    // swap pivot with the greater element at i
    Swap(ref array[i + 1], ref array[high]);

    // return the partition point
    return (i + 1);
}

static void QuickSort(int[] array, int low, int high)
{
    if (low < high)
    {
        // find the pivot element such that
        // elements smaller than pivot are on left of pivot
        // elements greater than pivot are on righ of pivot
        int pi = Partition(array, low, high);

        // recursive call on the left of pivot
        QuickSort(array, low, pi - 1);

        // recursive call on the right of pivot
        QuickSort(array, pi + 1, high);
    }
}

static void Main(string[] args)
{
    int[] data = { 8, 7, 6, 1, 0, 9, 2 };
    int n = data.Length;

    Console.WriteLine("Unsorted Array: ");
    PrintArray(data, n);

    // perform quicksort on data
    QuickSort(data, 0, n - 1);

    Console.WriteLine("Sorted array in ascending order: ");
    PrintArray(data, n);
}

// This code is contributed by Ali Barzegari Dahaj