class Solution {
public:
  vector<int> twoSum(vector<int>& nums, int target) {
    map<int, int> cache;
    vector<int> result;
    for (vector<int>::size_type i = 0; i < nums.size(); i++) {
      if (cache.count(nums[i]) > 0) {
        result.push_back(cache[nums[i]]);
        result.push_back(i);
        return result;
      }
      cache[target - nums[i]] = i;
    }
    return result;
  }
};