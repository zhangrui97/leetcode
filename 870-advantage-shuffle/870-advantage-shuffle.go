func advantageCount(nums1 []int, nums2 []int) []int {
  sort.Ints(nums1)
  indx := make([][2]int, len(nums2))
  for i := 0; i < len(nums2); i++ {
    indx[i] = [2]int{nums2[i], i}
  }
  sort.Slice(indx, func(i, j int) bool {
    return indx[i][0] > indx[j][0]
  })
  l, r := 0, len(nums2) - 1
  ans := make([]int, len(nums2))
  for _, v := range indx {
    if nums1[r] <= v[0] {
      ans[v[1]] = nums1[l]
      l++
    } else {
      ans[v[1]] = nums1[r]
      r--
    }
  }
  return ans
}