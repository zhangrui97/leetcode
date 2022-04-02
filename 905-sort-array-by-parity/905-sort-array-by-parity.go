func sortArrayByParity(nums []int) []int {
  n := len(nums)
  fast, slow := 0, 0
  for slow < n && fast < n {
    for ; slow < n; slow ++ {
      if nums[slow] % 2 == 1 {
        break
      }
    }
    fast = slow + 1
    for ; fast < n; fast ++ {
      if nums[fast] % 2 == 0 {
        break
      }
    }
    if slow < n && fast < n {
      nums[slow], nums[fast] = nums[fast], nums[slow]
    }
    slow ++
  }
  return nums
}