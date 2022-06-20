func min(a, b int)int {
  if a < b { return a }
  return b
}
func minFallingPathSum(matrix [][]int) int {
  n := len(matrix)
  for i := n - 2; i >= 0; i-- {
    rk := matrix[i+1]
    for j := 0; j < n-1; j++ {
      rk[j] = min(rk[j], rk[j+1])
    } 
    ri := matrix[i]
    ri[0] += rk[0]
    for j := 1; j < n-1; j++ {
      ri[j] += min(rk[j-1], rk[j])
    }
    ri[n-1] += rk[n-2]
  }
  r0 := matrix[0]
  ans := r0[0]
  for i := 1; i < n; i++ {
    if r0[i] < ans {
      ans = r0[i]
    }
  }
  return ans
}