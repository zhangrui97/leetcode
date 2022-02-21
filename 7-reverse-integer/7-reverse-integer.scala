import scala.collection.mutable.Stack

object Solution {
  def reverse(x: Int): Int = {
    if (x <= -Int.MaxValue) return 0
    if (x < 0) return -reverse(-x)  
    val stack = Stack[Int]()
    var rest = x
    while (rest > 0) {
      val d = rest % 10
      if (d != 0 || !stack.isEmpty) {
        stack.push(d)
      }
      rest /= 10
    }
    val rev = stack.reverse
    var result = 0
    while (!rev.isEmpty) {
      val last = rev.pop
      if ((Int.MaxValue - last)/10 < result) return 0
      result = result * 10 + last
    }
    return result
  }
}