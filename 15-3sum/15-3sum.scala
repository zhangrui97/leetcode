object Solution {
  def threeSum(nums: Array[Int]): List[List[Int]] = {
    var result = List[List[Int]]()
    val len = nums.length
    if (len < 3) {
      return result
    }
    val sorted = nums.sorted
    if (sorted(0) > 0 || sorted(len-1) < 0) {
      return result
    }    
    if (sorted(0) == 0 && sorted(len-1) == 0) {
      result ::= List(0, 0, 0)
      return result
    }
    var i = 0
    while (i < len-2) {
      val a = sorted(i)
      if (a > 0) {
        return result
      }
      var l = i + 1
      var r = len - 1
      while (l < r) {
        val b = sorted(l)
        val c = sorted(r)
        if (c < 0) {
          return result
        }
        if (a+b+c < 0) {
          do {
            l += 1
          } while (l < r && b == sorted(l))
        } else if (a+b+c == 0) {
          result ::= List(a, b, c)
          do {
            l += 1
          } while (l < r && b == sorted(l))
          do {
            r -= 1
          } while (l < r && c == sorted(r))
        } else {
          do {
            r -= 1
          } while (l < r && c == sorted(r))
        }
      }
      do {
        i += 1
      } while (i < len-2 && a == sorted(i))
    }
    return result
  }
}