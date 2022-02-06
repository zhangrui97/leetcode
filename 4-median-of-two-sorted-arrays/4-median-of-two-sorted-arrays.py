import sys
class Solution:
  def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
    n = len(nums1)
    m = len(nums2)
    if n > m: return self.findMedianSortedArrays(nums2, nums1)

    l1 = 0
    r1 = 0
    l2 = 0
    r2 = 0
    lo = 0
    hi = 2 * n
    while lo <= hi:
      c1 = (lo + hi) // 2
      c2 = n + m - c1
      l1 = -sys.maxsize if c1 == 0 else nums1[(c1-1)//2]
      r1 = sys.maxsize if c1 == 2*n else nums1[c1//2]
      l2 = -sys.maxsize if c2 == 0 else nums2[(c2-1)//2]
      r2 = sys.maxsize if c2 == 2*m else nums2[c2//2]
      
      if l1 > r2: hi = c1-1
      elif l2 > r1: lo = c1+1
      else: break

    return (max(l1, l2) + min(r1, r2)) / 2.0