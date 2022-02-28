public class Solution {
  public int RemoveElement(int[] nums, int val) {
    int fast = 0, slow = 0;
    while (fast < nums.Length) {
        if (nums[fast] != val) {
            nums[slow] = nums[fast];
            slow++;
        }
        fast++;
    }
    return slow;
  }
}