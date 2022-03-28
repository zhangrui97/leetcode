func runningSum(nums []int) []int {
  ans := make([]int, len(nums))
  sum := 0
  for i, v := range nums {
    sum += v
    ans[i] = sum
  }
  return ans
}