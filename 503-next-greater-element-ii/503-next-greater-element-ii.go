func nextGreaterElements(nums []int) []int {
  ans := make([]int, len(nums))
  for i := 0; i < len(nums); i++ {
    ans[i] = -1
  }
  stack := make([]int, 0, len(nums))
  for i, v := range nums {
    for l:=len(stack); l > 0 && v > nums[stack[l-1]]; l-- {
      ans[stack[l-1]] = v
      stack = stack[:l-1]
    }
    stack = append(stack, i)
  }
  for _, v := range nums {
    for l:=len(stack); l > 0 && v > nums[stack[l-1]]; l-- {
      ans[stack[l-1]] = v
      stack = stack[:l-1]
    }
  }
  return ans
}