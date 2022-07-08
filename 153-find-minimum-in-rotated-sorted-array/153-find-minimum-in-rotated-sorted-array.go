func min(a, b int)int {
  if a < b {
    return a
  } else {
    return b
  }
}

func findMin(nums []int) int {
  n := len(nums)
  if n == 1 { return nums[0] }
  if n == 2 { return min(nums[0], nums[1]) }
  l, mid, r := nums[0], nums[(n-1) / 2], nums[n-1]
  if l < r { return nums[0] }
  if mid < r {
    return findMin(nums[:(n+1)/2])
  } else {
    return findMin(nums[(n+1)/2:])
  }
}