object Solution {
  def maxArea(height: Array[Int]): Int = {
    var li = 0
    var ri = height.length - 1
    var l = height(li)
    var r = height(ri)
    var maxH = 0
    var max = 0
    while (li < ri) {
      val minH = math.min(l, r)
      if (minH > maxH) {
        maxH = minH
        max = math.max(max, maxH * (ri - li))
      }
      if (l < r) {
        li += 1
        l = height(li)
      } else if (l > r) {
        ri -= 1
        r = height(ri)
      } else {
        li += 1
        l = height(li)
        ri -= 1
        r = height(ri)
      }
    }
    max
  }
}