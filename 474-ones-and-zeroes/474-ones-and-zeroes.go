func max(a, b int)int {
  if a < b {
    return b
  } else {
    return a
  }
}

func findMaxForm(strs []string, m int, n int) int {
  dp := make([][]int, m + 1)
  for i := range dp {
    dp[i] = make([]int, n + 1)
  }
  for _, s := range strs {
    ones, zeroes := 0, 0
    for _, ch := range s {
      if ch == '1' { ones++ }
      if ch == '0' { zeroes++ }
    }
    for i := m; i >= zeroes; i-- {
      for j:= n; j >= ones; j-- {
        dp[i][j] = max(dp[i][j], dp[i-zeroes][j-ones] + 1)
      }
    }
  }  
  return dp[m][n]
}