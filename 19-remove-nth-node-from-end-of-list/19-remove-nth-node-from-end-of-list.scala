/**
 * Definition for singly-linked list.
 * class ListNode(_x: Int = 0, _next: ListNode = null) {
 *   var next: ListNode = _next
 *   var x: Int = _x
 * }
 */
object Solution {
  def removeNthFromEnd(head: ListNode, n: Int): ListNode = {
    var f = head
    for (_ <- 0 to n) {
      if (f == null) return head.next
      else f = f.next
    }
    var s = head
    while (f != null) {
      f = f.next
      s = s.next
    }
    s.next = s.next.next
    head
  }
}