object Solution {
  def threeSumClosest(nums: Array[Int], target: Int): Int = {
    if (nums.length <= 3) return nums.sum
    val sorted = nums.sorted
    val avg = target / 3
    var (lo, hi) = (0, sorted.length - 1)
    if (sorted(lo) >= avg + 1) return sorted(0) + sorted(1) + sorted(2)
    if (sorted(hi) <= avg) return sorted(hi) + sorted(hi-1) + sorted(hi - 2)
    var result = 2000
    for (j <- lo to hi-2) {
      val a = sorted(j)
      val rest = target - a
      var (l, r) = (j+1, hi)
      while (l < r) {
        val sum = sorted(l) + sorted(r)
        if (sum < rest) {
          l += 1
          if ((rest - sum) < ((result - target).abs)) {
            result = sum + a
          }
        } else if (sum == rest) {
          return target
        } else { // sum > rest
          r -= 1
          if ((sum - rest) < ((result - target).abs)) {
            result = sum + a
          }
        }
      }
    }
    return result
  }
}