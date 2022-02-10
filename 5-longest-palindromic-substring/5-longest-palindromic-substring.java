class Solution {
  public String longestPalindrome(String s) {
    String str = "^#" + String.join("#", s.split("")) + "#$";
    int len = str.length();
    int[] dp = new int[len];
    Arrays.fill(dp, 1);
    int index = 0;
    int center = 0;
    int right = 0;
    for (int i = 2; i < len - 2; i++) {
      if (i < right - 1) {
        dp[i] = Math.min(right - i, dp[2*center-i]);
      }
      while (str.charAt(i-dp[i]) == str.charAt(i+dp[i])) {
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
    int realIndex = (index-dp[index])/2;
    return s.substring(realIndex, realIndex + dp[index]-1);
  }
}