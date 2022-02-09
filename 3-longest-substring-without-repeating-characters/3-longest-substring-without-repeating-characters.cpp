class Solution {
public:
  int lengthOfLongestSubstring(string s) {
    vector<int> dp(256, -1);
    int last = -1;
    int result = 0;
    for (string::size_type i = 0; i < s.size(); i++) {
      char ch = s[i];
      if (dp[ch] > last) {
        last = dp[ch];
      } else {
        if (i - last > result) {
          result = i - last;
        }
      }
      dp[ch] = i;
    }
    return result;
  }
};