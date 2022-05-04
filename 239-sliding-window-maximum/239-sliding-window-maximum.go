func maxSlidingWindow(nums []int, k int) []int {
  ans := make([]int, len(nums) - k + 1)
  q := make([]int, 0, len(nums))
  for i, v := range nums {
    for l:= len(q); l > 0 && v > q[l-1]; l-- {
      q = q[:l-1]
    }
    q = append(q, v)
    if i >= k && nums[i-k] == q[0] {
      q = q[1:]
    }
    if i >= k - 1 {
      ans[i+1-k] = q[0]
    }
  }
  return ans
}