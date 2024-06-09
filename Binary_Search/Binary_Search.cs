// Binary Search in C#
using System;

public class Binary_Search
{
    // An iterative binary search function.
    public static int BinarySearch(int[] arr, int low, int high, int x)
    {
        while (low <= high)
        {
            int mid = low + (high - low) / 2;

            // Check if x is present at mid
            if (arr[mid] == x)
                return mid;

            // If x greater, ignore left half
            if (arr[mid] < x)
                low = mid + 1;

            // If x is smaller, ignore right half
            else
                high = mid - 1;
        }

        // If we reach here, then element was not present
        return -1;
    }

    public static void Main(void)
    {
        int[] arr = { 2, 3, 4, 10, 40 };
        int x = 10;
        int n = arr.Length;
        int result = BinarySearch(arr, 0, n - 1, x);
        Console.WriteLine(result == -1
            ? "Element is not present in array"
            : $"Element is present at index {result}");
    }
}

// This code is contributed by Ali Barzegari Dahaj
