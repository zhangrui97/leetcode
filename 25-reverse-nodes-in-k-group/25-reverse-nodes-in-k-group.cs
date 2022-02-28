/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
  private ListNode Reverse(ListNode head, ListNode tail) {
    ListNode pre = null;
    ListNode current = head;
    ListNode next = head;
    while (current != tail) {
      next = current.next;
      current.next = pre;
      pre = current;
      current = next;
    }
    return pre;
  }
  
  public ListNode ReverseKGroup(ListNode head, int k) {
    ListNode tail = head;
    for (int i = 0; i < k; i++) {
      if (tail == null) return head;
      tail = tail.next;
    }
    var newHead = Reverse(head, tail);
    head.next = ReverseKGroup(tail, k);
    return newHead;
  }
}