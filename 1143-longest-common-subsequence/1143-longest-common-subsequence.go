func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}

func longestCommonSubsequence(text1 string, text2 string) int {
  n1, n2 := len(text1), len(text2)
  result := 0
  dp := make([][]int, n1+1)
  for i := 0; i <= n1; i++ {
    dp[i] = make([]int, n2+1)
    for j := 0; j <= n2; j++ {
      if i == 0 || j == 0 {
        dp[i][j] = 0
      } else if text1[i-1] == text2[j-1] {
        dp[i][j] = dp[i-1][j-1] + 1
        result = max(dp[i][j], result)
      } else {
        dp[i][j] = max(dp[i-1][j], dp[i][j-1])
      }
    }
  }
  return result
}