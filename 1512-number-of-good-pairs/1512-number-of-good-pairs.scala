object Solution {
  def numIdenticalPairs(nums: Array[Int]): Int = {
    val seen = collection.mutable.Map[Int, Int]()
    var count = 0
    for (n <- nums) {
      if (seen contains n) {
        count = count + seen(n)
        seen(n) = seen(n) + 1
      } else {
        seen(n) = 1
      }
    }
    return count
  }
}