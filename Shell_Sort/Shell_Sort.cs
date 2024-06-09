// Shell Sort in C#
using System;

public class Shell_Sort
{
    // Shell sort
    public static void ShellSort(int[] array, int n)
    {
        // Rearrange elements at each n/2, n/4, n/8, ... intervals
        for (int interval = n / 2; interval > 0; interval /= 2)
        {
            for (int i = interval; i < n; i += 1)
            {
                int temp = array[i];
                int j;
                for (j = i; j >= interval && array[j - interval] > temp; j -= interval)
                {
                    array[j] = array[j - interval];
                }
                array[j] = temp;
            }
        }
    }

    // Print array function
    public static void PrintArray(int[] array, int size)
    {
        for (int i = 0; i < size; i++)
            Console.Write(array[i] + " ");
        Console.WriteLine();
    }

    // Main
    public static void Main()
    {
        int[] data = { 9, 8, 3, 7, 5, 6, 4, 1 };
        int size = data.Length;
        ShellSort(data, size);
        Console.WriteLine("Sorted array:");
        PrintArray(data, size);
    }
}

// This code is contributed by Ali Barzegari Dahaj
