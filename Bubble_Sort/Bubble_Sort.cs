//  Bubble Sort in C#
using System;

public class Bubble_Sort
{
    // Bubble Sort function
    public static void BubbleSort(int[] arr)
    {
        int n = arr.Length;
        bool swapped;
        for (int i = 0; i < n - 1; i++)
        {
            swapped = false;
            for (int j = 0; j < n - i - 1; j++)
            {
                if (arr[j] > arr[j + 1])
                {
                    (arr[j], arr[j + 1]) = (arr[j + 1], arr[j]);
                    swapped = true;
                }
            }
            // If no two elements were swapped
            // by inner loop, then break
            if (!swapped)
                break;
        }
    }

    // Function to print an array
    public static void PrintArray(int[] arr)
    {
        foreach (int i in arr)
            Console.Write(" " + i);
    }

    public static void Main(string[] args)
    {
        int[] arr = { 64, 34, 25, 12, 22, 11, 90 };
        BubbleSort(arr);
        Console.WriteLine("Sorted array: ");
        PrintArray(arr);
    }
}

// This code is contributed by Ali Barzegari Dahaj