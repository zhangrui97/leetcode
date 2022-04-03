func minSubsequence(nums []int) []int {
  n := len(nums)
  sort.Ints(nums)
  sum := 0
  for _, v := range nums {
    sum += v
  }
  max := 0
  for i := n - 1; i >= 0; i-- {
    max += nums[i]
    if (max<<1) > sum {
      
      for j := i; j < (n+i)/2; j ++ {
        nums[j], nums[n-1-j+i] = nums[n-1-j+i], nums[j]
      }
      return nums[i:]
    }
  }
  return []int{}
}