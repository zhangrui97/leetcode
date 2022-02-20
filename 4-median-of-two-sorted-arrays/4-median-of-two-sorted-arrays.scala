import scala.util.control.Breaks._
import scala.math.{min, max}

object Solution {
  def findMedianSortedArrays(nums1: Array[Int], nums2: Array[Int]): Double = {
    val m = nums1.length
    val n = nums2.length
    if (m > n) return findMedianSortedArrays(nums2, nums1)
    var l1 = 0
    var r1 = 0
    var l2 = 0
    var r2 = 0
    var lo = 0
    var hi = 2*m
    breakable {
      while (lo <= hi) {
        val c1 = (lo + hi) / 2
        val c2 = m + n - c1
        l1 = if (c1 == 0) Int.MinValue else nums1((c1-1)/2)
        r1 = if (c1 == 2*m) Int.MaxValue else nums1(c1/2)
        l2 = if (c2 == 0) Int.MinValue else nums2((c2-1)/2)
        r2 = if (c2 == 2*n) Int.MaxValue else nums2(c2/2)
        if (l1 > r2) {
          hi = c1-1
        } else if (l2 > r1) {
          lo = c1+1
        } else break
      }
    }
    return (max(l1, l2) + min(r1, r2)).toFloat / 2.0
  }
}