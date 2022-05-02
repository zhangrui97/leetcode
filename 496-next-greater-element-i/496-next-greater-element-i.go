func nextGreaterElement(nums1 []int, nums2 []int) []int {
  ans := make([]int, len(nums1))
  l := len(nums2)
  for j, v := range nums1 {
    i := 0
    for ; nums2[i] != v; i++ {}
    for ; i < l && nums2[i] <= v; i++ {}
    if i >= l {
      ans[j] = -1
    } else {
      ans[j] = nums2[i]
    }   
  }
  return ans
}