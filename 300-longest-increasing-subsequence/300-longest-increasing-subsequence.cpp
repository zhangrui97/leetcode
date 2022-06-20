class Solution {
public:
  int lengthOfLIS(vector<int>& nums) {
    vector<int> dp(nums.size());
    int ans = 1;
    dp[0] = 1;
    for (int i = 1; i < nums.size(); i++) {
      int v = nums[i];
      int res = 1;
      for (int j = 0; j < i; j++) {
        if (v > nums[j]) {
          res = max(dp[j]+1, res);
        }
      }
      dp[i] = res;
      if (res > ans) {
        ans = res;
      }
    }
    return ans;
  }
};