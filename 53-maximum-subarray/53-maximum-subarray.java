class Solution {
  public int maxSubArray(int[] nums) {
    int ans = -10001;
    int sum = 0;
    for (int n : nums) {
      sum = (sum <= 0) ? n : sum + n;
      ans = Math.max(ans, sum);      
    }
    return ans;
  }
}