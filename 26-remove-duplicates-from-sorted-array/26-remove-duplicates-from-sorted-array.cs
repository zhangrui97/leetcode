public class Solution {
  public int RemoveDuplicates(int[] nums) {
    if (nums.Length < 2) return nums.Length;
    int count = 0;
    for (int i = 1; i < nums.Length; i++) {
      if (nums[i] != nums[count]) {
        count++;
        nums[count] = nums[i];
      } 
    }
    return count+1;
  }
}