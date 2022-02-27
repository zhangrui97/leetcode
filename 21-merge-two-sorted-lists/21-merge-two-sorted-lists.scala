/**
 * Definition for singly-linked list.
 * class ListNode(_x: Int = 0, _next: ListNode = null) {
 *   var next: ListNode = _next
 *   var x: Int = _x
 * }
 */
object Solution {
  def mergeTwoLists(list1: ListNode, list2: ListNode): ListNode = {
    val dummy = new ListNode()
    var p1 = list1
    var p2 = list2
    var p = dummy
    while (p1 != null && p2 != null) {
      if (p1.x < p2.x) {
        p.next = p1
        p1 = p1.next
      } else {
        p.next = p2
        p2 = p2.next
      }
      p = p.next
    }
    if (p1 == null) {
      p.next = p2
    }
    if (p2 == null) {
      p.next = p1
    }
    return dummy.next
  }
}