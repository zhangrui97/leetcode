/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
  if (nums1.length > nums2.length) return findMedianSortedArrays(nums2, nums1)
  const half = i => ~~(i/2)
  const n = nums1.length
  const m = nums2.length
  let l1 = r1 = l2 = r2 = 0
  let lo = 0
  let hi = 2 * n
  while (lo <= hi) {
    const c1 = half(lo + hi)
    const c2 = m + n - c1
    l1 = (c1 === 0) ? Number.MIN_SAFE_INTEGER : nums1[half(c1-1)]
    r1 = (c1 === 2 * n) ? Number.MAX_SAFE_INTEGER : nums1[half(c1)]
    l2 = (c2 === 0) ? Number.MIN_SAFE_INTEGER : nums2[half(c2-1)]
    r2 = (c2 === 2 * m) ? Number.MAX_SAFE_INTEGER : nums2[half(c2)]
    
    if (l1 > r2) hi = c1 - 1
    else if (l2 > r1) lo = c1 + 1
    else break
  }
  return (Math.max(l1, l2) + Math.min(r1, r2)) / 2.0
};