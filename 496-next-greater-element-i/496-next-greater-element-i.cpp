class Solution {
public:
  vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
    map<int, int> dict;
    stack<int> stack;
    for (int i = nums2.size()-1; i >= 0; i--) {
      auto v{nums2[i]};
      while (!stack.empty() && stack.top() <= v) {
        stack.pop();
      }
      if (stack.empty()) {
        dict.insert({v, -1});
      } else {
        dict.insert({v, stack.top()});        
      }
      stack.push(v);
    }
    vector<int> ans(nums1.size(), -1);
    for (int i = 0; i < nums1.size(); i++) {
      auto v{nums1[i]};
      if (dict.count(v)) {
        ans[i] = dict[v];
      }
    }
    return ans;
  }
};