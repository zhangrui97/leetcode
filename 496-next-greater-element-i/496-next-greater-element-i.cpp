class Solution {
public:
  vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
    unordered_map<int, int> dict;
    stack<int> stack;
    for (int i = nums2.size()-1; i >= 0; i--) {
      auto v{nums2[i]};
      while (!stack.empty() && stack.top() <= v) {
        stack.pop();
      }
      dict.insert({v, stack.empty() ? -1 : stack.top()});
      stack.push(v);
    }
    vector<int> ans(nums1.size());
    for (int i = 0; i < nums1.size(); i++) {
      ans[i] = dict[nums1[i]];
    }
    return ans;
  }
};