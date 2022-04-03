func sortedSquares(nums []int) []int {
  n := len(nums)
  l, r := 0, n - 1
  nums[0] = nums[0] * nums[0]
  if n == 1 {
    return nums
  }
  nums[r] = nums[r] * nums[r]
  ans := make([]int, n)
  for i := n - 1; i >= 0 ; i -- {
    if nums[l] > nums[r] {
      ans[i] = nums[l]
      l ++
      if l != r && l < n {
        nums[l] = nums[l]*nums[l]
      }
    } else {
      ans[i] = nums[r]
      r --
      if l != r && r >= 0 {
        nums[r] = nums[r]*nums[r]
      }
    }
  }  
  return ans
}