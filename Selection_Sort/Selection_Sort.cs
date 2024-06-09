// Selection sort in C#
using System;

class Selection_Sort
{
    // function to swap the the position of two elements
    static void Swap(ref int a, ref int b)
    {
        int temp = a;
        a = b;
        b = temp;
    }

    // function to print an array
    static void PrintArray(int[] array)
    {
        foreach (int i in array)
        {
            Console.Write(i + " ");
        }
        Console.WriteLine();
    }

    static void SelectionSort(int[] array)
    {
        int size = array.Length;
        for (int step = 0; step < size - 1; step++)
        {
            int minIdx = step;
            for (int i = step + 1; i < size; i++)
            {
                // To sort in descending order, change > to < in this line.
                // Select the minimum element in each loop.
                if (array[i] < array[minIdx])
                    minIdx = i;
            }
            // put min at the correct position
            Swap(ref array[minIdx], ref array[step]);
        }
    }

    static void Main(string[] args)
    {
        int[] data = { 20, 12, 10, 15, 2 };
        SelectionSort(data);
        Console.WriteLine("Sorted array in Ascending Order:");
        PrintArray(data);
    }
}

// This code is contributed by Ali Barzegari Dahaj