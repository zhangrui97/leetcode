func removeElement(nums []int, val int) int {
  l, r := 0, len(nums) - 1
  if r < 0 { return 0 }
  for l < r {
    for ; l < r && nums[l] != val; l++ {}
    for ; l < r && nums[r] == val; r-- {}
    if l < r {
      nums[l], nums[r] = nums[r], nums[l]
      l++
      r--
    }
  }
  if nums[l] == val { 
    return l
  } else { 
    return l + 1
  }
}