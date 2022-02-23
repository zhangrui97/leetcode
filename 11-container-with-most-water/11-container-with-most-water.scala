object Solution {
  def maxArea(height: Array[Int]): Int = {
    var (li, ri) = (0, height.length - 1)
    var max = 0
    while (li < ri) {
      max = max max ((height(li) min height(ri)) * (ri - li))
      val (l, r) = (height(li), height(ri))
      if (l <= r) {
        li += 1
      } 
      if (l >= r) {
        ri -= 1
      }
    }
    max
  }
}