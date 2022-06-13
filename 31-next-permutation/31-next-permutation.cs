public class Solution {
  public void NextPermutation(int[] nums) {
    int k = nums.Length - 2;
    for (; k >= 0; k--) {
      if (nums[k+1] > nums[k]) break;
    }
    if (k >= 0) {
      int l = nums.Length - 1;
      int v = nums[k];
      for (; l > k + 1; l--) {
        if (nums[l] > v) break;
      }
      nums[k] = nums[l];
      nums[l] = v;
    }
    for (int i = k+1, j = nums.Length-1; i < j; i++, j--) {
      (nums[i], nums[j]) = (nums[j], nums[i]);
    }
  }
}