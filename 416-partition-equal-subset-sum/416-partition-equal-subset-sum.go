func canPartition(nums []int) bool {
  sum := 0
  mx := 0
  for _, v := range nums {
    sum += v
    if v > mx { mx = v }
  }  
  if sum % 2 == 1 || mx + mx > sum { return false }
  sum /= 2
  if mx == sum { return true }
  dp := make([]bool, sum+1)
  dp[0] = true
  for _, v := range nums {
    for j := sum; j >= v; j-- {
      if dp[j-v] {
        dp[j] = true
      }
    }
  }
  return dp[sum]
}