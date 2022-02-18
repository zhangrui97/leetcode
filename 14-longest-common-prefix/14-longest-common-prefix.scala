object Solution {
  def longestCommonPrefix(strs: Array[String]): String = {
    var mins = strs.minBy(_.length)
    while (!strs.forall(_.indexOf(mins) == 0)) {
      mins = mins.substring(0, mins.length - 1)
    }
    return mins
  }
}