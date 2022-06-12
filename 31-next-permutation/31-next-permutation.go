func nextPermutation(nums []int)  {
  n := len(nums)
  k := n-2
  for ; k >= 0; k-- {
    if nums[k] < nums[k+1] {
      break
    }
  }
  if k < 0 {
    for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
      nums[i], nums[j] = nums[j], nums[i]
    }
    return 
  }
  v := nums[k]
  l := n-1
  for ; l > k; l-- {
    if nums[l] > v {
      break
    }
  }
  nums[k], nums[l] = nums[l], v
  for i, j := k+1, n-1; i < j; i, j = i+1, j-1 {
    nums[i], nums[j] = nums[j], nums[i]
  }
}
