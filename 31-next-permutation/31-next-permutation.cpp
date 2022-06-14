class Solution {
public:
  void nextPermutation(vector<int>& nums) {
    int n = nums.size();
    int k = n - 2;
    for (; k >= 0; k--) {
      if (nums[k+1] > nums[k]) break;
    }
    if (k >= 0) {
      int l = n - 1;
      int v = nums[k];
      for (; l > k+1; l--) {
        if (nums[l] > v) break;
      }
      swap(nums[k], nums[l]);
    }
    reverse(nums.begin()+k+1, nums.end());
  }
};