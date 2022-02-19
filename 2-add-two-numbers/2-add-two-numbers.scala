/**
 * Definition for singly-linked list.
 * class ListNode(_x: Int = 0, _next: ListNode = null) {
 *   var next: ListNode = _next
 *   var x: Int = _x
 * }
 */
object Solution {
  def addTwoNumbers(l1: ListNode, l2: ListNode): ListNode = {
    val dummy = ListNode(-1)
    var p1 = l1
    var p2 = l2
    var p = dummy
    var c = 0
    while (p1 != null || p2 != null || c > 0) {
      var sum = c
      if (p1 != null) {
        sum += p1.x
        p1 = p1.next
      }
      if (p2 != null) {
        sum += p2.x
        p2 = p2.next
      }
      p.next = ListNode(sum % 10)
      c = sum / 10
      p = p.next
    }
    return dummy.next
  }
}