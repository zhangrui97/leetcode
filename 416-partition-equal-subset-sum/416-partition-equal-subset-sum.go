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
  if mx == sum { return true }
  dpPre, dpCur := make([]int, sum+1), make([]int, sum+1)
  for i := 1; i <= n; i++ {
    dpCur, dpPre = dpPre, dpCur
    v := nums[i-1]
    for j := 1; j <= sum; j++ {
      if j < v {
        dpCur[j] = dpPre[j]
      } else if j == v {
        dpCur[j] = v
      } else {
        dpCur[j] = max(dpPre[j], dpPre[j-v] + v)
        if dpCur[j] == sum { return true }
      }
    }
  }
  return dpCur[sum] == sum
}