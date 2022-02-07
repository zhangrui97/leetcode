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
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
        var result = new ListNode();
        var p1 = l1;
        var p2 = l2;
        var carry = 0;
        var next = result;
        while (p1 != null || p2 != null) {
            int sum = carry + (p1?.val ?? 0) + (p2?.val ?? 0);
            next.next = new ListNode(sum % 10);
            p1 = p1?.next;
            p2 = p2?.next;
            next = next.next;
            carry = sum / 10;
        }
        if (carry > 0) {
          next.next = new ListNode(carry);
        }
        return result.next;
    }
}