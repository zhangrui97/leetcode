import scala.collection.immutable.HashMap

object Solution {
  def twoSum(nums: Array[Int], target: Int): Array[Int] = {
    var map: HashMap[Int, Int] = HashMap[Int, Int]()
    for (i <- 0 until nums.length) {
      val n = nums(i)
      if (map.contains(n)) {
        return Array(map(n), i)
      }
      map += (target - n) -> i
    }
    return Array.empty
  }
}