func maxSlidingWindow(nums []int, k int) []int {
  ans := make([]int, len(nums) - k + 1)
  q := make([]int, 0, len(nums))
  for i, v := range nums {
    for l:= len(q); l > 0 && v > q[l-1]; l-- {
      q = q[:l-1]
    }
    q = append(q, v)
    if i+1 >= k {
      max := q[0]
      ans[i+1-k] = max
      if nums[i+1-k] == max {
        q = q[1:]
      }
    }
  }
  return ans
}