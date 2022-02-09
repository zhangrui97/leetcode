class Solution {
  final int MAX = 1000000 + 1;
  final int MIN = -MAX;
  public double findMedianSortedArrays(int[] nums1, int[] nums2) {
    int n = nums1.length;
    int m = nums2.length;
    if (n > m) return findMedianSortedArrays(nums2, nums1);
    
    int l1 = 0;
    int r1 = 0;
    int l2 = 0;
    int r2 = 0;
    int lo = 0;
    int hi = 2 * n;
    while (lo <= hi) {
      int c1 = (lo + hi) / 2;
      int c2 = m + n - c1;
      l1 = (c1 == 0) ? MIN : nums1[(c1-1)/2];
      r1 = (c1 == 2*n) ? MAX : nums1[c1/2];
      l2 = (c2 == 0) ? MIN : nums2[(c2-1)/2];
      r2 = (c2 == 2*m) ? MAX : nums2[c2/2];
      
      if (l1 > r2) hi = c1 - 1;
      else if (l2 > r1) lo = c1 + 1;
      else break;
    }
    return ((double)(Math.max(l1, l2) + Math.min(r1, r2))) / 2.0;
  }
}