class Solution {
public:
  int lengthOfLongestSubstring(string s) {
    map<char, int> dp;
    int last = 0;
    int result = 0;
    for (string::size_type i = 0; i < s.size(); i++) {
      char ch = s[i];
      if (dp.count(ch) && dp[ch] >= last) {
        last = dp[ch] + 1;
      } else {
        if (i + 1 - last > result) {
          result = i + 1 - last;
        }
      }
      dp[ch] = i;
    }
    return result;
  }
};