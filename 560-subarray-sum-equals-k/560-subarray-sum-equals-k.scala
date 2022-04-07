object Solution {
  def subarraySum(nums: Array[Int], k: Int): Int = {
    val map = collection.mutable.HashMap[Int, Int](0 -> 1)
    var ans = 0
    var sum = 0
    for (v <- nums) {
      sum += v
      if (map.contains(sum-k)) {
        ans += map(sum-k)
      }
      if (map.contains(sum))
        map.put(sum, map(sum)+1)
      else
        map.put(sum, 1)
    }
    ans
  }
}