object Solution {
  def removeElement(nums: Array[Int], `val`: Int): Int = {
    var fast = 0
    while (fast < nums.length && nums(fast) != `val`) fast += 1
    var slow = fast
    while (fast < nums.length) {
      if (nums(fast) == `val`) fast += 1
      else {
        nums(slow) = nums(fast)
        slow += 1
        fast += 1
      }
    }
    slow
  }
}