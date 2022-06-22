func maxSubArray(nums []int) int {
  ans := -10001
  sum := 0
  for _, v := range nums {
    if sum > 0 {
      sum += v
    } else {
      sum = v
    }
    if sum > ans {
      ans = sum
    }
  }
  return ans
}