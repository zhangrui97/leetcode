func countBits(n int) []int {
  ans := make([]int, n+1)
  ans[0] = 0
  if n == 0 {
    return ans
  }
  ans[1] = 1
  for i,j := 2, 1; i <= n; i++ {
    if i == j << 1 {
      j = i
      ans[i] = 1
    } else {
      ans[i] = ans[i - j] + 1
    }
  }
  return ans  
}