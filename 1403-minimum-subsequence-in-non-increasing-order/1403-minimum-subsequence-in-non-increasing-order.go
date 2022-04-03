func minSubsequence(nums []int) []int {
  sort.Ints(nums)
  sum := 0
  for _, v := range nums {
    sum += v
  }
  ans := make([]int, 0, len(nums))
  max := 0
  for i := len(nums) - 1; i >= 0; i-- {
    ans = append(ans, nums[i])
    max += nums[i]
    if (max<<1) > sum {
      return ans
    }
  }
  return []int{}
}