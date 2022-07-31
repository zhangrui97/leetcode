#include <algorithm>
using namespace std;

class Solution 
{
public:
  double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    auto m = nums1.size();
    auto n = nums2.size();
    if (m > n) return findMedianSortedArrays(nums2, nums1);
    int l = 0;
    int r = 2 * m;
    int total = m + n;
    int c1{}, c2{};
    int l1{}, r1{}, l2{}, r2{};
    while (l <= r) {
      c1 = l + (r - l) / 2;
      c2 = total - c1;
      l1 = (c1 == 0) ? -1000001 : nums1[(c1-1)/2];
      r1 = (c1 == 2*m) ? 1000001 : nums1[c1/2];
      l2 = (c2 == 0) ? -1000001 : nums2[(c2-1)/2];
      r2 = (c2 == 2*n) ? 1000001 : nums2[c2/2];
      if (l1 > r2) r = c1 - 1;
      else if (l2 > r1) l = c1 + 1;
      else break;
    }
    return ((double)(max(l1, l2) + min(r1, r2)))/2.0;
  }
};