class Solution {
private:
  static string prepareString(string s) {
    vector<char> buffer;
    buffer.push_back('^');
    for (char ch: s) {
      buffer.push_back('#');
      buffer.push_back(ch);
    }
    buffer.push_back('#');
    buffer.push_back('$');
    string result(buffer.begin(), buffer.end());
    return result;
  }
public:
  string longestPalindrome(string s) {
    string str = prepareString(s);
    auto len = str.length();
    int index = 0;
    int center = 0;
    int right = 0;
    vector<int> dp(len, 1);
    for (int i = 1; i < len - 1 - dp[index]; i++) {
      if (i < right - 1) {
        dp[i] = min(right - i, dp[2*center - i]);
      }
      while (str[i - dp[i]] == str[i + dp[i]]) {
        dp[i] += 1;
      }
      if (i + dp[i] > right) {
        center = i;
        right = i + dp[i];
      }
      if (dp[i] > dp[index]) {
        index = i;
      }
    }
    return s.substr((index-dp[index])/2, dp[index] - 1);
  }
};