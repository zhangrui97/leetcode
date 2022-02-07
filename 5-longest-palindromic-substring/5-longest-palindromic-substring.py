class Solution:
  def longestPalindrome(self, s: str) -> str:
    if len(s) == 1: return s
    preprocess = '^#' + '#'.join(s) + '#$'
    dp = [1] * len(preprocess)
    index = 1
    center = 0
    right = 0
    for i in range(1, len(preprocess)-3):
      if i < (right-1):
        dp[i] = min(right-i, dp[2*center-i])
      while preprocess[i-dp[i]] == preprocess[i+dp[i]]:
        dp[i] += 1
      if i + dp[i] > right:
        center = i
        right = center + dp[i]
      if dp[i] > dp[index]:
        index = i
    
    return s[(index-dp[index])//2:(index-dp[index])//2+dp[index]-1]
    