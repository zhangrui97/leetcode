object Solution {
  def longestPalindrome(s: String): String = {
    if (s.length <= 1) return s
    val processed = "^#" + s.mkString("#") + "#$"
    val len = processed.length
    val dp = Array.fill[Int](len)(1)
    var maxIndex = 1
    var c = 2
    var r = 3
    var i = 3
    while (i < len - 2) {
      if (i < r - 1) {
        dp(i) = scala.math.min(dp(c*2-i), r - i)
      }
      while (processed(i-dp(i)) == processed(i+dp(i))) {
        dp(i) += 1
      }
      if (i + dp(i) > r) {
        c = i
        r = i + dp(i)
      }
      if (dp(i) > dp(maxIndex)) {
        maxIndex = i
      }
      i += 1
    }
    val index = (maxIndex - dp(maxIndex)) / 2
    return s.substring(index, index + dp(maxIndex)-1)
  }
}