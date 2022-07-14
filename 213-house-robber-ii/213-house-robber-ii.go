func max(a, b int)int { if a > b { return a } else { return b } }

func rob(nums []int) int {
  n := len(nums)
  if n == 1 { return nums[0] }
  if n == 2 { return max(nums[0], nums[1]) }
  dpa0, dpa1, dpb0, dpb1 := nums[0], max(nums[0], nums[1]), nums[1], max(nums[1], nums[2])
  for i, num := range nums {
    if i > 1 && i < n-1 {
      dpa0, dpa1 = dpa1, max(dpa1, dpa0 + num)
    }
    if i > 2 {
      dpb0, dpb1 = dpb1, max(dpb1, dpb0 + num)
    }
  }
  return max(dpa1, dpb1)
}