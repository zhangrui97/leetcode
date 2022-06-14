class Solution {
  public void nextPermutation(int[] nums) {
    int n = nums.length;
    int k = n - 2;
    for (;k >= 0; k--) {
      if (nums[k] < nums[k+1]) break;
    }
    if (k >= 0) {
      int v = nums[k];
      int l = n - 1;
      for (;l > k; l--) {
        if (nums[l] > v) break;
      }
      nums[k] = nums[l];
      nums[l] = v;
    }
    for (int i = k+1, j = n-1; i < j; i++, j--) {
      int temp = nums[i];
      nums[i] = nums[j];
      nums[j] = temp;
    }
  }
}