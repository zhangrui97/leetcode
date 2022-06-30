func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}

func canPartition(nums []int) bool {
  n := len(nums)
  sum := 0
  mx := 0
  for _, v := range nums {
    sum += v
    if v > mx { mx = v }
  }  
  if sum % 2 == 1 || mx + mx > sum { return false }
  sum /= 2
  dp := make([][]int, n+1)
  dp[0] = make([]int, sum+1)
  for i := 1; i <= n; i++ {
    v := nums[i-1]
    dp[i] = make([]int, sum+1)
    for j := 1; j <= sum; j++ {
      if v > j {
        dp[i][j] = dp[i-1][j]
      } else {
        dp[i][j] = max(dp[i-1][j], dp[i-1][j-v] + v)
      }
    }
  }
  return dp[n][sum] == sum
}