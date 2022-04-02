func sumZero(n int) []int {
  ans := make([]int, n)
  for i := 0; i < n/2; i++ {
    ans[i] = -i-1
    ans[n-1-i] = i+1
  }
  if n % 2 == 1 {
    ans[n/2] = 0
  }
  return ans
}