// Insertion sort in C#
using System;

class Insertion_Sort
{
    // Function to print an array
    static void PrintArray(int[] array, int size)
    {
        for (int i = 0; i < size; i++)
        {
            Console.Write(array[i] + " ");
        }
        Console.WriteLine();
    }

    static void InsertionSort(int[] array, int size)
    {
        for (int step = 1; step < size; step++)
        {
            int key = array[step];
            int j = step - 1;

            // Compare key with each element on the left of it until an element smaller than
            // it is found.
            // For descending order, change key < array[j] to key > array[j].
            while (key < array[j] && j >= 0)
            {
                array[j + 1] = array[j];
                --j;
            }
            array[j + 1] = key;
        }
    }

    static void Main(string[] args)
    {
        int[] data = { 9, 5, 1, 4, 3 };
        int size = data.Length;
        InsertionSort(data, size);
        Console.WriteLine("Sorted array in ascending order:");
        PrintArray(data, size);
    }
}

// This code is contributed by Ali Barzegari Dahaj