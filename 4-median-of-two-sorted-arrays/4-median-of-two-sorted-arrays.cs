public class Solution {
    public double FindMedianSortedArrays(int[] nums1, int[] nums2) {
      if (nums1.Length > nums2.Length) return FindMedianSortedArrays(nums2, nums1);
      int n = nums1.Length;
      int m = nums2.Length;
      int l1 = 0;
      int r1 = 0;
      int l2 = 0;
      int r2 = 0;
      int lo = 0;
      int hi = n * 2;
      while (lo <= hi) {
        int c1 = (lo + hi) / 2;
        int c2 = m + n - c1;
        l1 = c1 == 0 ? Int32.MinValue : nums1[(c1-1)/2];
        r1 = c1 == 2*n ? Int32.MaxValue : nums1[c1/2];
        l2 = c2 == 0 ? Int32.MinValue : nums2[(c2-1)/2];
        r2 = c2 == 2*m ? Int32.MaxValue : nums2[c2/2];
        if (l1 > r2) {
          hi = c1 - 1;
        } else if (l2 > r1) {
          lo = c1 + 1;
        } else break;          
      }
      return (Math.Max((double)l1, (double)l2) + Math.Min((double)r1, (double)r2)) / 2.0;
    }
}