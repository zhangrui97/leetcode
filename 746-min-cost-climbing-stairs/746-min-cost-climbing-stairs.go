func min(a, b int) int {
  if a <= b {
    return a
  }
  return b
}

func minCostClimbingStairs(cost []int) int {
  dp0, dp1 := cost[0], cost[1]
  for i := 2; i < len(cost); i ++ {
    dp0, dp1 = dp1, min(dp0, dp1) + cost[i]
  }
  return min(dp0, dp1)
}