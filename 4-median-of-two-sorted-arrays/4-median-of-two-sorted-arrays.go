func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  m, n := len(nums1), len(nums2)
  if m > n { return findMedianSortedArrays(nums2, nums1) }
  l1, r1, l2, r2 := 0, 0, 0, 0
  lo, hi := 0, 2*m
  for lo <= hi {
    c1 := (lo + hi) / 2
    c2 := m + n - c1
    if c1 == 0 {
      l1 = -1000000
    } else {
      l1 = nums1[(c1 - 1) / 2]
    }
    if c1 == 2 * m {
      r1 = 1000000
    } else {
      r1 = nums1[c1 / 2]
    }
    if c2 == 0 {
      l2 = -1000000
    } else {
      l2 = nums2[(c2 - 1) / 2]
    }
    if c2 == 2 * n {
      r2 = 1000000
    } else {
      r2 = nums2[c2 / 2]
    }
    if l1 > r2 {
      hi = c1 - 1
    } else if l2 > r1 {
      lo = c1 + 1
    } else { break }
  }
  return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2.0
}