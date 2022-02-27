import scala.collection.immutable.HashMap
import scala.collection.mutable.Stack
object Solution {
  def isValid(s: String): Boolean = {
    var map = HashMap('(' -> ')', '[' -> ']', '{' -> '}')
    var stack = new Stack[Char]
    for (ch <- s) {
      if (map contains ch) {
        stack.push(map(ch))
      } else {
        if (stack.isEmpty || stack.pop != ch)
          return false
      }
    }
    return stack.isEmpty
  }
}