func nextGreaterElement(nums1 []int, nums2 []int) []int {
  ans := make([]int, len(nums1))
  for i := range ans {
    ans[i] = -1
  }
  dict := make(map[int]int)
  for i, v := range nums1 {
    dict[v] = i
  }
  stack := make([]int, 0, len(nums2))
  for _, v := range nums2 {
    for l:=len(stack); l > 0 && v > stack[l-1]; l=len(stack) {
      if ind, ok := dict[stack[l-1]]; ok {
        ans[ind] = v
      }
      stack = stack[:l-1]
    }
    stack = append(stack, v)
  }
  return ans
}