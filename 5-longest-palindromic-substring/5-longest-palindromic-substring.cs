public class Solution {
  public string LongestPalindrome(string s) {
    string str = "^#" + String.Join('#', s as IEnumerable<char>) + "#$";
    int index = 0;
    int center = 0;
    int right = 0;
    int n = str.Length;
    int[] dp = Enumerable.Repeat<int>(1, n).ToArray();
    
    for (int i = 1; i < n - 2; i++) {
      if (i < right - 1) {
        dp[i] = Math.Min(right - i, dp[2*center - i]);
      }
      while (str[i-dp[i]] == str[i+dp[i]]) {
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
    return s.Substring((index - dp[index])/2, dp[index]-1);
  }
}