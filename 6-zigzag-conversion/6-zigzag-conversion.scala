object Solution {
  def convert(s: String, numRows: Int): String = {
    val len = s.length
    if (len <= 1 || numRows <= 1 || len <= numRows) return s
    val sb = new StringBuilder()
    val step = 2*numRows - 2
    for (i <- 0 until numRows) {
      var index1 = i
      var index2 = if (i%(numRows-1) == 0) len else step - i
      while (index1 < len || index2 < len) {
        if (index1 < len) {
          sb += s.charAt(index1)
          index1 += step
        }
        if (index2 < len) {
          sb += s.charAt(index2)
          index2 += step
        }
      }
    }
    return sb.toString
  }
}