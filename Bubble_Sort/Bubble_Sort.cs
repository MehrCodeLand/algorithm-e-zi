public class Algo_BubbleSort
{


    public static void Main()
    {
        int[] nums = { 64, 34, 25, 12, 22, 11, 90 };
        Sort(nums);
        Print(nums);
    }

    public static void Sort(int[] nums)
    {
        for(var i = 0; i < nums.Length; i++)
        {
            for(var j = 0; j < nums.Length - i - 1; ++j)
            {
                if (nums[j] > nums[j + 1])
                {
                    // swap 
                    var temp = nums[j];
                    nums[j] = nums[j + 1];
                    nums[j + 1] = temp;
                }
            }
        }
    }

    public static void Print(int[] nums)
    {
        foreach(var i in nums)
        {
            Console.WriteLine(i);
        }
    }
}
