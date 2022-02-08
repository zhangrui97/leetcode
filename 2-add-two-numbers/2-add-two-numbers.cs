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
  static private ListNode AddCarry(int carry, ListNode l) {
    if (carry == 0) return l;
    else {
      if (l == null) return new ListNode(carry);
      if (l.val == 9) return new ListNode(0, AddCarry(1, l.next));
      l.val += 1;
      return l;
    }
  }
  public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
    var result = new ListNode();
    var p1 = l1;
    var p2 = l2;
    var carry = 0;
    var p = result;
    while (p1 != null && p2 != null) {
      int sum = carry + p1.val + p2.val;
      p.next = new ListNode(sum % 10);
      p1 = p1.next;
      p2 = p2.next;
      p = p.next;
      carry = sum / 10;
    }
    if (p1 == null) {
      p.next = AddCarry(carry, p2);
    } else {
      p.next = AddCarry(carry, p1);
    }
    return result.next;
  }
}