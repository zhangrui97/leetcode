/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
  public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
    int carry = 0;
    ListNode p1 = l1;
    ListNode p2 = l2;
    ListNode result = new ListNode();
    ListNode p = result;
    while (carry > 0 || p1 != null || p2 != null) {
      if (p1 != null && p2 != null) {
        int sum = carry + p1.val + p2.val;
        p.next = new ListNode(sum%10);
        carry = sum / 10;
        p1 = p1.next;
        p2 = p2.next;
        p = p.next;
      } else if (p1 == null && p2 == null) {
        p.next = new ListNode(carry);
        break;
      } else if (carry == 0) {
        p.next = (p1 == null) ? p2 : p1;
        break;
      } else {
        if (p1 == null) {
          p1 = p2;
          p2 = null;
        } else {
          int sum = carry + p1.val;
          p.next = new ListNode(sum%10);
          carry = sum / 10;
          p1 = p1.next;
          p = p.next;
        }
      }
    }
    return result.next;
  }
}