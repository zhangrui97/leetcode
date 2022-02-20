import scala.collection.mutable.HashMap
object Solution {
  def lengthOfLongestSubstring(s: String): Int = {
    if (s.length < 2) return s.length
    val indexCache = HashMap(s.charAt(0) -> 0)
    var max = 1
    var left = 0
    var right = 1
    while (right < s.length) {
      val ch = s.charAt(right)
      if (indexCache.contains(ch) && indexCache(ch) >= left) {
        left = indexCache(ch) + 1
      } else { // char at right == char at left
        if (max < right - left  + 1) {
          max = right - left + 1
        }
      }
      indexCache.put(ch, right);
      right += 1
    }
    return max
  }
}