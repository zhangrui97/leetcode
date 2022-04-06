class Solution {
public:
  vector<int> searchRange(vector<int>& nums, int target) {
    auto begin = nums.begin();
    auto end = nums.end();
    auto first = std::lower_bound(begin, end, target);
    if (!(first == end) && !(target < *first)) {
      auto last = std::upper_bound(begin, end, target);
      return {(int)std::distance(begin, first), (int)std::distance(begin, last)-1};
    } else {
      return {-1, -1};
    }
  }
};