func getConcatenation(nums []int) []int {
  n := len(nums)
  ans := make([]int, n*2)
  for i, v := range nums {
    ans[i] = v
    ans[i + n] = v
  }
  return ans
}